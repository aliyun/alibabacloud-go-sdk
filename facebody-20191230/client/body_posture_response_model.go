// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBodyPostureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BodyPostureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BodyPostureResponse
	GetStatusCode() *int32
	SetBody(v *BodyPostureResponseBody) *BodyPostureResponse
	GetBody() *BodyPostureResponseBody
}

type BodyPostureResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BodyPostureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BodyPostureResponse) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponse) GoString() string {
	return s.String()
}

func (s *BodyPostureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BodyPostureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BodyPostureResponse) GetBody() *BodyPostureResponseBody {
	return s.Body
}

func (s *BodyPostureResponse) SetHeaders(v map[string]*string) *BodyPostureResponse {
	s.Headers = v
	return s
}

func (s *BodyPostureResponse) SetStatusCode(v int32) *BodyPostureResponse {
	s.StatusCode = &v
	return s
}

func (s *BodyPostureResponse) SetBody(v *BodyPostureResponseBody) *BodyPostureResponse {
	s.Body = v
	return s
}

func (s *BodyPostureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
