// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKeystoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateKeystoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateKeystoresResponse
	GetStatusCode() *int32
	SetBody(v *UpdateKeystoresResponseBody) *UpdateKeystoresResponse
	GetBody() *UpdateKeystoresResponseBody
}

type UpdateKeystoresResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateKeystoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateKeystoresResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateKeystoresResponse) GoString() string {
	return s.String()
}

func (s *UpdateKeystoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateKeystoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateKeystoresResponse) GetBody() *UpdateKeystoresResponseBody {
	return s.Body
}

func (s *UpdateKeystoresResponse) SetHeaders(v map[string]*string) *UpdateKeystoresResponse {
	s.Headers = v
	return s
}

func (s *UpdateKeystoresResponse) SetStatusCode(v int32) *UpdateKeystoresResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateKeystoresResponse) SetBody(v *UpdateKeystoresResponseBody) *UpdateKeystoresResponse {
	s.Body = v
	return s
}

func (s *UpdateKeystoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
