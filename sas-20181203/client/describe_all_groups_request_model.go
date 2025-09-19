// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAllGroupsRequest
	GetLang() *string
}

type DescribeAllGroupsRequest struct {
	// The natural language of the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAllGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAllGroupsRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAllGroupsRequest) SetLang(v string) *DescribeAllGroupsRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAllGroupsRequest) Validate() error {
	return dara.Validate(s)
}
