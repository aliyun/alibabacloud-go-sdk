// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExecutorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExecutorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExecutorsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExecutorsResponseBody) *UpdateExecutorsResponse
	GetBody() *UpdateExecutorsResponseBody
}

type UpdateExecutorsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExecutorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExecutorsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExecutorsResponse) GoString() string {
	return s.String()
}

func (s *UpdateExecutorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExecutorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExecutorsResponse) GetBody() *UpdateExecutorsResponseBody {
	return s.Body
}

func (s *UpdateExecutorsResponse) SetHeaders(v map[string]*string) *UpdateExecutorsResponse {
	s.Headers = v
	return s
}

func (s *UpdateExecutorsResponse) SetStatusCode(v int32) *UpdateExecutorsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExecutorsResponse) SetBody(v *UpdateExecutorsResponseBody) *UpdateExecutorsResponse {
	s.Body = v
	return s
}

func (s *UpdateExecutorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
