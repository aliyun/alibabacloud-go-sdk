// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsvUserSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IsvUserSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IsvUserSaveResponse
	GetStatusCode() *int32
	SetBody(v *IsvUserSaveResponseBody) *IsvUserSaveResponse
	GetBody() *IsvUserSaveResponseBody
}

type IsvUserSaveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IsvUserSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IsvUserSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s IsvUserSaveResponse) GoString() string {
	return s.String()
}

func (s *IsvUserSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IsvUserSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IsvUserSaveResponse) GetBody() *IsvUserSaveResponseBody {
	return s.Body
}

func (s *IsvUserSaveResponse) SetHeaders(v map[string]*string) *IsvUserSaveResponse {
	s.Headers = v
	return s
}

func (s *IsvUserSaveResponse) SetStatusCode(v int32) *IsvUserSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *IsvUserSaveResponse) SetBody(v *IsvUserSaveResponseBody) *IsvUserSaveResponse {
	s.Body = v
	return s
}

func (s *IsvUserSaveResponse) Validate() error {
	return dara.Validate(s)
}
