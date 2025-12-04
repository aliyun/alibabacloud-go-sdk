// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateErResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateErResponse
	GetStatusCode() *int32
	SetBody(v *UpdateErResponseBody) *UpdateErResponse
	GetBody() *UpdateErResponseBody
}

type UpdateErResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateErResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateErResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateErResponse) GoString() string {
	return s.String()
}

func (s *UpdateErResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateErResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateErResponse) GetBody() *UpdateErResponseBody {
	return s.Body
}

func (s *UpdateErResponse) SetHeaders(v map[string]*string) *UpdateErResponse {
	s.Headers = v
	return s
}

func (s *UpdateErResponse) SetStatusCode(v int32) *UpdateErResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateErResponse) SetBody(v *UpdateErResponseBody) *UpdateErResponse {
	s.Body = v
	return s
}

func (s *UpdateErResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
