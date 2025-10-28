// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatasetJobConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetJobConfigs(v []*DatasetJobConfig) *ListDatasetJobConfigsResponseBody
	GetDatasetJobConfigs() []*DatasetJobConfig
	SetRequestId(v string) *ListDatasetJobConfigsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDatasetJobConfigsResponseBody
	GetTotalCount() *int32
}

type ListDatasetJobConfigsResponseBody struct {
	// The dataset job configurations.
	DatasetJobConfigs []*DatasetJobConfig `json:"DatasetJobConfigs,omitempty" xml:"DatasetJobConfigs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDatasetJobConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatasetJobConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatasetJobConfigsResponseBody) GetDatasetJobConfigs() []*DatasetJobConfig {
	return s.DatasetJobConfigs
}

func (s *ListDatasetJobConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatasetJobConfigsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDatasetJobConfigsResponseBody) SetDatasetJobConfigs(v []*DatasetJobConfig) *ListDatasetJobConfigsResponseBody {
	s.DatasetJobConfigs = v
	return s
}

func (s *ListDatasetJobConfigsResponseBody) SetRequestId(v string) *ListDatasetJobConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatasetJobConfigsResponseBody) SetTotalCount(v int32) *ListDatasetJobConfigsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDatasetJobConfigsResponseBody) Validate() error {
	if s.DatasetJobConfigs != nil {
		for _, item := range s.DatasetJobConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
