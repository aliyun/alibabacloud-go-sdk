// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCheckRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCheckRunResponse
	GetStatusCode() *int32
	SetBody(v *CreateCheckRunResponseBody) *CreateCheckRunResponse
	GetBody() *CreateCheckRunResponseBody
}

type CreateCheckRunResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCheckRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCheckRunResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckRunResponse) GoString() string {
	return s.String()
}

func (s *CreateCheckRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCheckRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCheckRunResponse) GetBody() *CreateCheckRunResponseBody {
	return s.Body
}

func (s *CreateCheckRunResponse) SetHeaders(v map[string]*string) *CreateCheckRunResponse {
	s.Headers = v
	return s
}

func (s *CreateCheckRunResponse) SetStatusCode(v int32) *CreateCheckRunResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCheckRunResponse) SetBody(v *CreateCheckRunResponseBody) *CreateCheckRunResponse {
	s.Body = v
	return s
}

func (s *CreateCheckRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
