// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListMultimodalDatasetRequest
	GetDBClusterId() *string
	SetInputField(v string) *ListMultimodalDatasetRequest
	GetInputField() *string
	SetPageNumber(v int32) *ListMultimodalDatasetRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListMultimodalDatasetRequest
	GetPageSize() *int32
}

type ListMultimodalDatasetRequest struct {
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
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListMultimodalDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalDatasetRequest) GoString() string {
	return s.String()
}

func (s *ListMultimodalDatasetRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListMultimodalDatasetRequest) GetInputField() *string {
	return s.InputField
}

func (s *ListMultimodalDatasetRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMultimodalDatasetRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMultimodalDatasetRequest) SetDBClusterId(v string) *ListMultimodalDatasetRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListMultimodalDatasetRequest) SetInputField(v string) *ListMultimodalDatasetRequest {
	s.InputField = &v
	return s
}

func (s *ListMultimodalDatasetRequest) SetPageNumber(v int32) *ListMultimodalDatasetRequest {
	s.PageNumber = &v
	return s
}

func (s *ListMultimodalDatasetRequest) SetPageSize(v int32) *ListMultimodalDatasetRequest {
	s.PageSize = &v
	return s
}

func (s *ListMultimodalDatasetRequest) Validate() error {
	return dara.Validate(s)
}
