// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayoutDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLayoutDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLayoutDataResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLayoutDataResponseBody) *UpdateLayoutDataResponse
	GetBody() *UpdateLayoutDataResponseBody
}

type UpdateLayoutDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLayoutDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLayoutDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayoutDataResponse) GoString() string {
	return s.String()
}

func (s *UpdateLayoutDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLayoutDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLayoutDataResponse) GetBody() *UpdateLayoutDataResponseBody {
	return s.Body
}

func (s *UpdateLayoutDataResponse) SetHeaders(v map[string]*string) *UpdateLayoutDataResponse {
	s.Headers = v
	return s
}

func (s *UpdateLayoutDataResponse) SetStatusCode(v int32) *UpdateLayoutDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLayoutDataResponse) SetBody(v *UpdateLayoutDataResponseBody) *UpdateLayoutDataResponse {
	s.Body = v
	return s
}

func (s *UpdateLayoutDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
