// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModelConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModelConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateModelConnectionResponseBody) *CreateModelConnectionResponse
	GetBody() *CreateModelConnectionResponseBody
}

type CreateModelConnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModelConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModelConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModelConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateModelConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModelConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModelConnectionResponse) GetBody() *CreateModelConnectionResponseBody {
	return s.Body
}

func (s *CreateModelConnectionResponse) SetHeaders(v map[string]*string) *CreateModelConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateModelConnectionResponse) SetStatusCode(v int32) *CreateModelConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModelConnectionResponse) SetBody(v *CreateModelConnectionResponseBody) *CreateModelConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateModelConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
