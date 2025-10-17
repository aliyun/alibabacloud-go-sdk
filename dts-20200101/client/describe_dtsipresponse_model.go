// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTSIPResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDTSIPResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDTSIPResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDTSIPResponseBody) *DescribeDTSIPResponse
	GetBody() *DescribeDTSIPResponseBody
}

type DescribeDTSIPResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDTSIPResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDTSIPResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTSIPResponse) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDTSIPResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDTSIPResponse) GetBody() *DescribeDTSIPResponseBody {
	return s.Body
}

func (s *DescribeDTSIPResponse) SetHeaders(v map[string]*string) *DescribeDTSIPResponse {
	s.Headers = v
	return s
}

func (s *DescribeDTSIPResponse) SetStatusCode(v int32) *DescribeDTSIPResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDTSIPResponse) SetBody(v *DescribeDTSIPResponseBody) *DescribeDTSIPResponse {
	s.Body = v
	return s
}

func (s *DescribeDTSIPResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
