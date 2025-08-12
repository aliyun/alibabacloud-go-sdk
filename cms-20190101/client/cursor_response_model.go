// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCursorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CursorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CursorResponse
	GetStatusCode() *int32
	SetBody(v *CursorResponseBody) *CursorResponse
	GetBody() *CursorResponseBody
}

type CursorResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CursorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CursorResponse) String() string {
	return dara.Prettify(s)
}

func (s CursorResponse) GoString() string {
	return s.String()
}

func (s *CursorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CursorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CursorResponse) GetBody() *CursorResponseBody {
	return s.Body
}

func (s *CursorResponse) SetHeaders(v map[string]*string) *CursorResponse {
	s.Headers = v
	return s
}

func (s *CursorResponse) SetStatusCode(v int32) *CursorResponse {
	s.StatusCode = &v
	return s
}

func (s *CursorResponse) SetBody(v *CursorResponseBody) *CursorResponse {
	s.Body = v
	return s
}

func (s *CursorResponse) Validate() error {
	return dara.Validate(s)
}
