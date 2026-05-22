// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateTraceDiagnoseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GenerateTraceDiagnoseResponseBody
	GetRequestId() *string
	SetTip(v string) *GenerateTraceDiagnoseResponseBody
	GetTip() *string
	SetUrl(v string) *GenerateTraceDiagnoseResponseBody
	GetUrl() *string
}

type GenerateTraceDiagnoseResponseBody struct {
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose/?id=xxxxxxx
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GenerateTraceDiagnoseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateTraceDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateTraceDiagnoseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateTraceDiagnoseResponseBody) GetTip() *string {
	return s.Tip
}

func (s *GenerateTraceDiagnoseResponseBody) GetUrl() *string {
	return s.Url
}

func (s *GenerateTraceDiagnoseResponseBody) SetRequestId(v string) *GenerateTraceDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateTraceDiagnoseResponseBody) SetTip(v string) *GenerateTraceDiagnoseResponseBody {
	s.Tip = &v
	return s
}

func (s *GenerateTraceDiagnoseResponseBody) SetUrl(v string) *GenerateTraceDiagnoseResponseBody {
	s.Url = &v
	return s
}

func (s *GenerateTraceDiagnoseResponseBody) Validate() error {
	return dara.Validate(s)
}
