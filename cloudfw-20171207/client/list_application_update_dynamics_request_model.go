// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationUpdateDynamicsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ListApplicationUpdateDynamicsRequest
	GetLang() *string
	SetPageNo(v int64) *ListApplicationUpdateDynamicsRequest
	GetPageNo() *int64
	SetPageSize(v int64) *ListApplicationUpdateDynamicsRequest
	GetPageSize() *int64
}

type ListApplicationUpdateDynamicsRequest struct {
	// The language of the returned information. Valid values:
	//
	// - **zh**: Chinese. This is the default value.
	//
	// - **en**: English.
	//
	// This parameter determines whether the update description and application name are returned in Chinese or English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. The value must be greater than or equal to 1. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListApplicationUpdateDynamicsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationUpdateDynamicsRequest) GoString() string {
	return s.String()
}

func (s *ListApplicationUpdateDynamicsRequest) GetLang() *string {
	return s.Lang
}

func (s *ListApplicationUpdateDynamicsRequest) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListApplicationUpdateDynamicsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListApplicationUpdateDynamicsRequest) SetLang(v string) *ListApplicationUpdateDynamicsRequest {
	s.Lang = &v
	return s
}

func (s *ListApplicationUpdateDynamicsRequest) SetPageNo(v int64) *ListApplicationUpdateDynamicsRequest {
	s.PageNo = &v
	return s
}

func (s *ListApplicationUpdateDynamicsRequest) SetPageSize(v int64) *ListApplicationUpdateDynamicsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApplicationUpdateDynamicsRequest) Validate() error {
	return dara.Validate(s)
}
