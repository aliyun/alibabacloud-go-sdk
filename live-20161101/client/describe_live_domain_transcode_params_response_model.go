// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTranscodeParamsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainTranscodeParamsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainTranscodeParamsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainTranscodeParamsResponseBody) *DescribeLiveDomainTranscodeParamsResponse
	GetBody() *DescribeLiveDomainTranscodeParamsResponseBody
}

type DescribeLiveDomainTranscodeParamsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainTranscodeParamsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainTranscodeParamsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTranscodeParamsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTranscodeParamsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainTranscodeParamsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainTranscodeParamsResponse) GetBody() *DescribeLiveDomainTranscodeParamsResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainTranscodeParamsResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainTranscodeParamsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsResponse) SetStatusCode(v int32) *DescribeLiveDomainTranscodeParamsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsResponse) SetBody(v *DescribeLiveDomainTranscodeParamsResponseBody) *DescribeLiveDomainTranscodeParamsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainTranscodeParamsResponse) Validate() error {
	return dara.Validate(s)
}
