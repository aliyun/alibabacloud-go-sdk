// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScanTaskStatisticsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLevels(v string) *DescribeScanTaskStatisticsRequest
	GetLevels() *string
}

type DescribeScanTaskStatisticsRequest struct {
	// The severities of the alert events handled by the virus detection task. Separate multiple severities with commas (,). The severities decrease in descending order. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious,suspicious,remind
	Levels *string `json:"Levels,omitempty" xml:"Levels,omitempty"`
}

func (s DescribeScanTaskStatisticsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeScanTaskStatisticsRequest) GoString() string {
	return s.String()
}

func (s *DescribeScanTaskStatisticsRequest) GetLevels() *string {
	return s.Levels
}

func (s *DescribeScanTaskStatisticsRequest) SetLevels(v string) *DescribeScanTaskStatisticsRequest {
	s.Levels = &v
	return s
}

func (s *DescribeScanTaskStatisticsRequest) Validate() error {
	return dara.Validate(s)
}
