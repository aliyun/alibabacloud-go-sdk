// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeWorkModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeWorkModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeWorkModeResponse
	GetStatusCode() *int32
	SetBody(v *ChangeWorkModeResponseBody) *ChangeWorkModeResponse
	GetBody() *ChangeWorkModeResponseBody
}

type ChangeWorkModeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeWorkModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeWorkModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeWorkModeResponse) GoString() string {
	return s.String()
}

func (s *ChangeWorkModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeWorkModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeWorkModeResponse) GetBody() *ChangeWorkModeResponseBody {
	return s.Body
}

func (s *ChangeWorkModeResponse) SetHeaders(v map[string]*string) *ChangeWorkModeResponse {
	s.Headers = v
	return s
}

func (s *ChangeWorkModeResponse) SetStatusCode(v int32) *ChangeWorkModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeWorkModeResponse) SetBody(v *ChangeWorkModeResponseBody) *ChangeWorkModeResponse {
	s.Body = v
	return s
}

func (s *ChangeWorkModeResponse) Validate() error {
	return dara.Validate(s)
}
