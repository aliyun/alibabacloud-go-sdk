// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSelectItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSelectItemRequest
	GetLang() *string
	SetRegId(v string) *DescribeSelectItemRequest
	GetRegId() *string
}

type DescribeSelectItemRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSelectItemRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSelectItemRequest) GoString() string {
	return s.String()
}

func (s *DescribeSelectItemRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSelectItemRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSelectItemRequest) SetLang(v string) *DescribeSelectItemRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSelectItemRequest) SetRegId(v string) *DescribeSelectItemRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSelectItemRequest) Validate() error {
	return dara.Validate(s)
}
