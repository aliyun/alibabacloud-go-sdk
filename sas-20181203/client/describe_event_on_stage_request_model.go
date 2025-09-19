// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventOnStageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventOnStageRequest
	GetLang() *string
}

type DescribeEventOnStageRequest struct {
	// The language of the content within the request and response. Default value: **zh**. Valid values:
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

func (s DescribeEventOnStageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventOnStageRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventOnStageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventOnStageRequest) SetLang(v string) *DescribeEventOnStageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventOnStageRequest) Validate() error {
	return dara.Validate(s)
}
