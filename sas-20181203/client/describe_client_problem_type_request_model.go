// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientProblemTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeClientProblemTypeRequest
	GetLang() *string
}

type DescribeClientProblemTypeRequest struct {
	// The language type for the request and response messages. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s DescribeClientProblemTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientProblemTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeClientProblemTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeClientProblemTypeRequest) SetLang(v string) *DescribeClientProblemTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeClientProblemTypeRequest) Validate() error {
	return dara.Validate(s)
}
