// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalFineTuneDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalFineTuneDatasetRequest
	GetDBClusterId() *string
	SetInputField(v string) *ListMultimodalFineTuneDatasetRequest
	GetInputField() *string
	SetPageNumber(v int32) *ListMultimodalFineTuneDatasetRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalFineTuneDatasetRequest
	GetPageSize() *int32
}

type ListMultimodalFineTuneDatasetRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// dataset1
	InputField *string `json:"InputField,omitempty" xml:"InputField,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalFineTuneDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalFineTuneDatasetRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalFineTuneDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalFineTuneDatasetRequest) GetInputField() *string {
	return s.InputField
}

func (s *ListMultimodalFineTuneDatasetRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalFineTuneDatasetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalFineTuneDatasetRequest) SetDBClusterId(v string) *ListMultimodalFineTuneDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetRequest) SetInputField(v string) *ListMultimodalFineTuneDatasetRequest {
	s.InputField = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetRequest) SetPageNumber(v int32) *ListMultimodalFineTuneDatasetRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetRequest) SetPageSize(v int32) *ListMultimodalFineTuneDatasetRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalFineTuneDatasetRequest) Validate() error {
	return dara.Validate(s)
}
