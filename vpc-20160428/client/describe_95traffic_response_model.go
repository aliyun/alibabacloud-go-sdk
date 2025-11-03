// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribe95TrafficResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *Describe95TrafficResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *Describe95TrafficResponse
	GetStatusCode() *int32
	SetBody(v *Describe95TrafficResponseBody) *Describe95TrafficResponse
	GetBody() *Describe95TrafficResponseBody
}

type Describe95TrafficResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Describe95TrafficResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s Describe95TrafficResponse) String() string {
	return dara.Prettify(s)
}

func (s Describe95TrafficResponse) GoString() string {
	return s.String()
}

func (s *Describe95TrafficResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *Describe95TrafficResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *Describe95TrafficResponse) GetBody() *Describe95TrafficResponseBody {
	return s.Body
}

func (s *Describe95TrafficResponse) SetHeaders(v map[string]*string) *Describe95TrafficResponse {
	s.Headers = v
	return s
}

func (s *Describe95TrafficResponse) SetStatusCode(v int32) *Describe95TrafficResponse {
	s.StatusCode = &v
	return s
}

func (s *Describe95TrafficResponse) SetBody(v *Describe95TrafficResponseBody) *Describe95TrafficResponse {
	s.Body = v
	return s
}

func (s *Describe95TrafficResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
