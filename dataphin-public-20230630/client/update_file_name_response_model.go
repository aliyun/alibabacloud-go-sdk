// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileNameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileNameResponseBody) *UpdateFileNameResponse
	GetBody() *UpdateFileNameResponseBody
}

type UpdateFileNameResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileNameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileNameResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileNameResponse) GetBody() *UpdateFileNameResponseBody {
	return s.Body
}

func (s *UpdateFileNameResponse) SetHeaders(v map[string]*string) *UpdateFileNameResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileNameResponse) SetStatusCode(v int32) *UpdateFileNameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileNameResponse) SetBody(v *UpdateFileNameResponseBody) *UpdateFileNameResponse {
	s.Body = v
	return s
}

func (s *UpdateFileNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
