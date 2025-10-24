// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateServiceVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateServiceVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateServiceVersionResponseBody) *UpdateServiceVersionResponse
	GetBody() *UpdateServiceVersionResponseBody
}

type UpdateServiceVersionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateServiceVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateServiceVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceVersionResponse) GoString() string {
	return s.String()
}

func (s *UpdateServiceVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateServiceVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateServiceVersionResponse) GetBody() *UpdateServiceVersionResponseBody {
	return s.Body
}

func (s *UpdateServiceVersionResponse) SetHeaders(v map[string]*string) *UpdateServiceVersionResponse {
	s.Headers = v
	return s
}

func (s *UpdateServiceVersionResponse) SetStatusCode(v int32) *UpdateServiceVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateServiceVersionResponse) SetBody(v *UpdateServiceVersionResponseBody) *UpdateServiceVersionResponse {
	s.Body = v
	return s
}

func (s *UpdateServiceVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
