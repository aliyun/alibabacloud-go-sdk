// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCdnDiagnoseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveCdnDiagnoseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveCdnDiagnoseInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveCdnDiagnoseInfoResponseBody) *DescribeLiveCdnDiagnoseInfoResponse
	GetBody() *DescribeLiveCdnDiagnoseInfoResponseBody
}

type DescribeLiveCdnDiagnoseInfoResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveCdnDiagnoseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveCdnDiagnoseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCdnDiagnoseInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) GetBody() *DescribeLiveCdnDiagnoseInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveCdnDiagnoseInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) SetStatusCode(v int32) *DescribeLiveCdnDiagnoseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) SetBody(v *DescribeLiveCdnDiagnoseInfoResponseBody) *DescribeLiveCdnDiagnoseInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveCdnDiagnoseInfoResponse) Validate() error {
	return dara.Validate(s)
}
