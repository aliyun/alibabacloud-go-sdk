// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectorNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCollectorNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCollectorNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCollectorNameResponseBody) *UpdateCollectorNameResponse
	GetBody() *UpdateCollectorNameResponseBody
}

type UpdateCollectorNameResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCollectorNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCollectorNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCollectorNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCollectorNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCollectorNameResponse) GetBody() *UpdateCollectorNameResponseBody {
	return s.Body
}

func (s *UpdateCollectorNameResponse) SetHeaders(v map[string]*string) *UpdateCollectorNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCollectorNameResponse) SetStatusCode(v int32) *UpdateCollectorNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCollectorNameResponse) SetBody(v *UpdateCollectorNameResponseBody) *UpdateCollectorNameResponse {
	s.Body = v
	return s
}

func (s *UpdateCollectorNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
