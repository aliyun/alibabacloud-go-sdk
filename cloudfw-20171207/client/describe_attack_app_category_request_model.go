// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAttackAppCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAttackAppCategoryRequest
	GetLang() *string
}

type DescribeAttackAppCategoryRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeAttackAppCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAttackAppCategoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAttackAppCategoryRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAttackAppCategoryRequest) SetLang(v string) *DescribeAttackAppCategoryRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAttackAppCategoryRequest) Validate() error {
	return dara.Validate(s)
}
