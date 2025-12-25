// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeThreatEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeThreatEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeThreatEventResponse
	GetStatusCode() *int32
	SetBody(v *DescribeThreatEventResponseBody) *DescribeThreatEventResponse
	GetBody() *DescribeThreatEventResponseBody
}

type DescribeThreatEventResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeThreatEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeThreatEventResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeThreatEventResponse) GoString() string {
	return s.String()
}

func (s *DescribeThreatEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeThreatEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeThreatEventResponse) GetBody() *DescribeThreatEventResponseBody {
	return s.Body
}

func (s *DescribeThreatEventResponse) SetHeaders(v map[string]*string) *DescribeThreatEventResponse {
	s.Headers = v
	return s
}

func (s *DescribeThreatEventResponse) SetStatusCode(v int32) *DescribeThreatEventResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeThreatEventResponse) SetBody(v *DescribeThreatEventResponseBody) *DescribeThreatEventResponse {
	s.Body = v
	return s
}

func (s *DescribeThreatEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
