// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorServicesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListModelOperatorServicesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelOperatorServicesRequest
	GetPageSize() *int32
}

type ListModelOperatorServicesRequest struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListModelOperatorServicesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorServicesRequest) GoString() string {
	return s.String()
}

func (s *ListModelOperatorServicesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelOperatorServicesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelOperatorServicesRequest) SetPageNumber(v int32) *ListModelOperatorServicesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListModelOperatorServicesRequest) SetPageSize(v int32) *ListModelOperatorServicesRequest {
	s.PageSize = &v
	return s
}

func (s *ListModelOperatorServicesRequest) Validate() error {
	return dara.Validate(s)
}
