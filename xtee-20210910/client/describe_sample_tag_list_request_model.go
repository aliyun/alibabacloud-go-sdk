// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleTagListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleTagListRequest
	GetLang() *string
	SetRegId(v string) *DescribeSampleTagListRequest
	GetRegId() *string
}

type DescribeSampleTagListRequest struct {
	// Sets the language type for requests and responses, with a default value of **zh**. Values:
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

func (s DescribeSampleTagListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleTagListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleTagListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleTagListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleTagListRequest) SetLang(v string) *DescribeSampleTagListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleTagListRequest) SetRegId(v string) *DescribeSampleTagListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleTagListRequest) Validate() error {
	return dara.Validate(s)
}
