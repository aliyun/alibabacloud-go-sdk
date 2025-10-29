// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMessageAppResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMessageAppResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMessageAppResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMessageAppResponseBody) *UpdateMessageAppResponse
	GetBody() *UpdateMessageAppResponseBody
}

type UpdateMessageAppResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMessageAppResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMessageAppResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMessageAppResponse) GoString() string {
	return s.String()
}

func (s *UpdateMessageAppResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMessageAppResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMessageAppResponse) GetBody() *UpdateMessageAppResponseBody {
	return s.Body
}

func (s *UpdateMessageAppResponse) SetHeaders(v map[string]*string) *UpdateMessageAppResponse {
	s.Headers = v
	return s
}

func (s *UpdateMessageAppResponse) SetStatusCode(v int32) *UpdateMessageAppResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMessageAppResponse) SetBody(v *UpdateMessageAppResponseBody) *UpdateMessageAppResponse {
	s.Body = v
	return s
}

func (s *UpdateMessageAppResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
