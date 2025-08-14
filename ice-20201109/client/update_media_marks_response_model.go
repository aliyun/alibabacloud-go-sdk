// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMediaMarksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMediaMarksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMediaMarksResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMediaMarksResponseBody) *UpdateMediaMarksResponse
	GetBody() *UpdateMediaMarksResponseBody
}

type UpdateMediaMarksResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMediaMarksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMediaMarksResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMediaMarksResponse) GoString() string {
	return s.String()
}

func (s *UpdateMediaMarksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMediaMarksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMediaMarksResponse) GetBody() *UpdateMediaMarksResponseBody {
	return s.Body
}

func (s *UpdateMediaMarksResponse) SetHeaders(v map[string]*string) *UpdateMediaMarksResponse {
	s.Headers = v
	return s
}

func (s *UpdateMediaMarksResponse) SetStatusCode(v int32) *UpdateMediaMarksResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMediaMarksResponse) SetBody(v *UpdateMediaMarksResponseBody) *UpdateMediaMarksResponse {
	s.Body = v
	return s
}

func (s *UpdateMediaMarksResponse) Validate() error {
	return dara.Validate(s)
}
