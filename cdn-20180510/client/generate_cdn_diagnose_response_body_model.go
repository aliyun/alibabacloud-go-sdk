// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateCdnDiagnoseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *GenerateCdnDiagnoseResponseBody
	GetContent() *string
	SetRequestId(v string) *GenerateCdnDiagnoseResponseBody
	GetRequestId() *string
}

type GenerateCdnDiagnoseResponseBody struct {
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose/?id=xxxxxxx
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateCdnDiagnoseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateCdnDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateCdnDiagnoseResponseBody) GetContent() *string {
	return s.Content
}

func (s *GenerateCdnDiagnoseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateCdnDiagnoseResponseBody) SetContent(v string) *GenerateCdnDiagnoseResponseBody {
	s.Content = &v
	return s
}

func (s *GenerateCdnDiagnoseResponseBody) SetRequestId(v string) *GenerateCdnDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateCdnDiagnoseResponseBody) Validate() error {
	return dara.Validate(s)
}
