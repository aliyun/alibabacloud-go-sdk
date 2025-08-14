// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafStartConfigRequest
	GetLang() *string
	SetRegId(v string) *DescribeSafStartConfigRequest
	GetRegId() *string
}

type DescribeSafStartConfigRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
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

func (s DescribeSafStartConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafStartConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafStartConfigRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSafStartConfigRequest) SetLang(v string) *DescribeSafStartConfigRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafStartConfigRequest) SetRegId(v string) *DescribeSafStartConfigRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSafStartConfigRequest) Validate() error {
	return dara.Validate(s)
}
