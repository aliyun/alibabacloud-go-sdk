// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSingleCallByTtsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SingleCallByTtsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SingleCallByTtsResponse
	GetStatusCode() *int32
	SetBody(v *SingleCallByTtsResponseBody) *SingleCallByTtsResponse
	GetBody() *SingleCallByTtsResponseBody
}

type SingleCallByTtsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SingleCallByTtsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SingleCallByTtsResponse) String() string {
	return dara.Prettify(s)
}

func (s SingleCallByTtsResponse) GoString() string {
	return s.String()
}

func (s *SingleCallByTtsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SingleCallByTtsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SingleCallByTtsResponse) GetBody() *SingleCallByTtsResponseBody {
	return s.Body
}

func (s *SingleCallByTtsResponse) SetHeaders(v map[string]*string) *SingleCallByTtsResponse {
	s.Headers = v
	return s
}

func (s *SingleCallByTtsResponse) SetStatusCode(v int32) *SingleCallByTtsResponse {
	s.StatusCode = &v
	return s
}

func (s *SingleCallByTtsResponse) SetBody(v *SingleCallByTtsResponseBody) *SingleCallByTtsResponse {
	s.Body = v
	return s
}

func (s *SingleCallByTtsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
