// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleUploadPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSampleUploadPolicyRequest
	GetLang() *string
	SetRegId(v string) *DescribeSampleUploadPolicyRequest
	GetRegId() *string
}

type DescribeSampleUploadPolicyRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeSampleUploadPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleUploadPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeSampleUploadPolicyRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSampleUploadPolicyRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSampleUploadPolicyRequest) SetLang(v string) *DescribeSampleUploadPolicyRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSampleUploadPolicyRequest) SetRegId(v string) *DescribeSampleUploadPolicyRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSampleUploadPolicyRequest) Validate() error {
	return dara.Validate(s)
}
