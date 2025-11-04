// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetInstanceHealthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetInstanceHealthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetInstanceHealthResponse
	GetStatusCode() *int32
	SetBody(v *SetInstanceHealthResponseBody) *SetInstanceHealthResponse
	GetBody() *SetInstanceHealthResponseBody
}

type SetInstanceHealthResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetInstanceHealthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetInstanceHealthResponse) String() string {
	return dara.Prettify(s)
}

func (s SetInstanceHealthResponse) GoString() string {
	return s.String()
}

func (s *SetInstanceHealthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetInstanceHealthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetInstanceHealthResponse) GetBody() *SetInstanceHealthResponseBody {
	return s.Body
}

func (s *SetInstanceHealthResponse) SetHeaders(v map[string]*string) *SetInstanceHealthResponse {
	s.Headers = v
	return s
}

func (s *SetInstanceHealthResponse) SetStatusCode(v int32) *SetInstanceHealthResponse {
	s.StatusCode = &v
	return s
}

func (s *SetInstanceHealthResponse) SetBody(v *SetInstanceHealthResponseBody) *SetInstanceHealthResponse {
	s.Body = v
	return s
}

func (s *SetInstanceHealthResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
