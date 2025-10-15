// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenDocQaResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenDocQaResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenDocQaResultResponse
	GetStatusCode() *int32
	SetBody(v *GenDocQaResultResponseBody) *GenDocQaResultResponse
	GetBody() *GenDocQaResultResponseBody
}

type GenDocQaResultResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenDocQaResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenDocQaResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GenDocQaResultResponse) GoString() string {
	return s.String()
}

func (s *GenDocQaResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenDocQaResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenDocQaResultResponse) GetBody() *GenDocQaResultResponseBody {
	return s.Body
}

func (s *GenDocQaResultResponse) SetHeaders(v map[string]*string) *GenDocQaResultResponse {
	s.Headers = v
	return s
}

func (s *GenDocQaResultResponse) SetStatusCode(v int32) *GenDocQaResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GenDocQaResultResponse) SetBody(v *GenDocQaResultResponseBody) *GenDocQaResultResponse {
	s.Body = v
	return s
}

func (s *GenDocQaResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
