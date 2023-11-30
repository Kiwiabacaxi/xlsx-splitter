package myapp

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

func ProcessExcelFile(excelFileName string) ([]string, error) {
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		return nil, fmt.Errorf("open failed: %w", err)
	}

	var data [][]string
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for i, cell := range row.Cells {
				if len(data) <= i {
					data = append(data, []string{})
				}
				data[i] = append(data[i], fmt.Sprintf("\"%s\"", cell.String()))
			}
		}
	}

	var result []string
	for _, col := range data {
		for i, val := range col {
			if i != 0 {
				col[i] = strings.ToUpper(val)
			}
		}
		result = append(result, strings.Join(col, ","))
	}

	return result, nil
}
