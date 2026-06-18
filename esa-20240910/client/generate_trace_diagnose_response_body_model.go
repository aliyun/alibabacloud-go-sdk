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
	// The request ID.
	//
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// A diagnostic message.
	//
	// example:
	//
	// Origin IP x.x.x.x: Test URL returned HTTP 404 (non-2xx status).
	//
	// Origin IP x.x.x.x: DNS A record does not point to an ESA service IP.
	Tip *string `json:"Tip,omitempty" xml:"Tip,omitempty"`
	// The generated diagnostic link.
	//
	// example:
	//
	// http://cdn.dns-detect.alicdn.com/diagnose_v2?id=5d97ac9b&token=WFji65gy2mGNM11bD929%2BCMoyI6mbk2deRR9hOC6INH%2FoYbccQZQcvEn4wc%2FDPHlTshxRSAa5HokX%2BabItBpJ0FdnteROssomXqgIdjHpM46L%2BbaIeweZfsWG6QnbXT5n7O5APMyc%2Fe8d1o9PwwB429Ccks1FU1AfjNZfvBcLeo%3D
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
