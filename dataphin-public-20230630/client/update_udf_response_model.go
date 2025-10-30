// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUdfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUdfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUdfResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUdfResponseBody) *UpdateUdfResponse
	GetBody() *UpdateUdfResponseBody
}

type UpdateUdfResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUdfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUdfResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUdfResponse) GoString() string {
	return s.String()
}

func (s *UpdateUdfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUdfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUdfResponse) GetBody() *UpdateUdfResponseBody {
	return s.Body
}

func (s *UpdateUdfResponse) SetHeaders(v map[string]*string) *UpdateUdfResponse {
	s.Headers = v
	return s
}

func (s *UpdateUdfResponse) SetStatusCode(v int32) *UpdateUdfResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUdfResponse) SetBody(v *UpdateUdfResponseBody) *UpdateUdfResponse {
	s.Body = v
	return s
}

func (s *UpdateUdfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
