// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFieldRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeFieldRequest
	GetLang() *string
	SetQueryKey(v string) *DescribeFieldRequest
	GetQueryKey() *string
}

type DescribeFieldRequest struct {
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The key of the global configuration. Valid values:
	//
	// 	- **soar_filed_tags**: queries the input template of the playbook.
	//
	// This parameter is required.
	//
	// example:
	//
	// soar_filed_tags
	QueryKey *string `json:"QueryKey,omitempty" xml:"QueryKey,omitempty"`
}

func (s DescribeFieldRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFieldRequest) GoString() string {
	return s.String()
}

func (s *DescribeFieldRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFieldRequest) GetQueryKey() *string {
	return s.QueryKey
}

func (s *DescribeFieldRequest) SetLang(v string) *DescribeFieldRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFieldRequest) SetQueryKey(v string) *DescribeFieldRequest {
	s.QueryKey = &v
	return s
}

func (s *DescribeFieldRequest) Validate() error {
	return dara.Validate(s)
}
