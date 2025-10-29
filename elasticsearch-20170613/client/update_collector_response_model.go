// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCollectorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCollectorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCollectorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCollectorResponseBody) *UpdateCollectorResponse
	GetBody() *UpdateCollectorResponseBody
}

type UpdateCollectorResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCollectorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCollectorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCollectorResponse) GoString() string {
	return s.String()
}

func (s *UpdateCollectorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCollectorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCollectorResponse) GetBody() *UpdateCollectorResponseBody {
	return s.Body
}

func (s *UpdateCollectorResponse) SetHeaders(v map[string]*string) *UpdateCollectorResponse {
	s.Headers = v
	return s
}

func (s *UpdateCollectorResponse) SetStatusCode(v int32) *UpdateCollectorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCollectorResponse) SetBody(v *UpdateCollectorResponseBody) *UpdateCollectorResponse {
	s.Body = v
	return s
}

func (s *UpdateCollectorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
