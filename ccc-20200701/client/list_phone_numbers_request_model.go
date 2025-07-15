// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ListPhoneNumbersRequest
	GetActive() *bool
	SetInstanceId(v string) *ListPhoneNumbersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListPhoneNumbersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPhoneNumbersRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListPhoneNumbersRequest
	GetSearchPattern() *string
	SetUsage(v string) *ListPhoneNumbersRequest
	GetUsage() *string
}

type ListPhoneNumbersRequest struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0833
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// example:
	//
	// Bidirection
	Usage *string `json:"Usage,omitempty" xml:"Usage,omitempty"`
}

func (s ListPhoneNumbersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersRequest) GetActive() *bool {
	return s.Active
}

func (s *ListPhoneNumbersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPhoneNumbersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPhoneNumbersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPhoneNumbersRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListPhoneNumbersRequest) GetUsage() *string {
	return s.Usage
}

func (s *ListPhoneNumbersRequest) SetActive(v bool) *ListPhoneNumbersRequest {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetInstanceId(v string) *ListPhoneNumbersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetPageNumber(v int32) *ListPhoneNumbersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetPageSize(v int32) *ListPhoneNumbersRequest {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetSearchPattern(v string) *ListPhoneNumbersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListPhoneNumbersRequest) SetUsage(v string) *ListPhoneNumbersRequest {
	s.Usage = &v
	return s
}

func (s *ListPhoneNumbersRequest) Validate() error {
	return dara.Validate(s)
}
