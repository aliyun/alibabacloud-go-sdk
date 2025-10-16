// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclCheckQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAclCheckQuotaRequest
	GetLang() *string
}

type DescribeAclCheckQuotaRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAclCheckQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclCheckQuotaRequest) GoString() string {
	return s.String()
}

func (s *DescribeAclCheckQuotaRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAclCheckQuotaRequest) SetLang(v string) *DescribeAclCheckQuotaRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAclCheckQuotaRequest) Validate() error {
	return dara.Validate(s)
}
