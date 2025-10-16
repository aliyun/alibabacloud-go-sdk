// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceTypeAutoEnableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeResourceTypeAutoEnableRequest
	GetLang() *string
}

type DescribeResourceTypeAutoEnableRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeResourceTypeAutoEnableRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceTypeAutoEnableRequest) GoString() string {
	return s.String()
}

func (s *DescribeResourceTypeAutoEnableRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeResourceTypeAutoEnableRequest) SetLang(v string) *DescribeResourceTypeAutoEnableRequest {
	s.Lang = &v
	return s
}

func (s *DescribeResourceTypeAutoEnableRequest) Validate() error {
	return dara.Validate(s)
}
