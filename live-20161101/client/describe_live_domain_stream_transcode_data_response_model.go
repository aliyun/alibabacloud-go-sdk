// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainStreamTranscodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainStreamTranscodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainStreamTranscodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainStreamTranscodeDataResponseBody) *DescribeLiveDomainStreamTranscodeDataResponse
	GetBody() *DescribeLiveDomainStreamTranscodeDataResponseBody
}

type DescribeLiveDomainStreamTranscodeDataResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainStreamTranscodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainStreamTranscodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainStreamTranscodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) GetBody() *DescribeLiveDomainStreamTranscodeDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainStreamTranscodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) SetStatusCode(v int32) *DescribeLiveDomainStreamTranscodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) SetBody(v *DescribeLiveDomainStreamTranscodeDataResponseBody) *DescribeLiveDomainStreamTranscodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainStreamTranscodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
