// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocOssTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribePocOssTokenRequest
	GetLang() *string
	SetRegId(v string) *DescribePocOssTokenRequest
	GetRegId() *string
}

type DescribePocOssTokenRequest struct {
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

func (s DescribePocOssTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePocOssTokenRequest) GoString() string {
	return s.String()
}

func (s *DescribePocOssTokenRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribePocOssTokenRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribePocOssTokenRequest) SetLang(v string) *DescribePocOssTokenRequest {
	s.Lang = &v
	return s
}

func (s *DescribePocOssTokenRequest) SetRegId(v string) *DescribePocOssTokenRequest {
	s.RegId = &v
	return s
}

func (s *DescribePocOssTokenRequest) Validate() error {
	return dara.Validate(s)
}
