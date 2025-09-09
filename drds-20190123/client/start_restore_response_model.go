// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartRestoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartRestoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartRestoreResponse
	GetStatusCode() *int32
	SetBody(v *StartRestoreResponseBody) *StartRestoreResponse
	GetBody() *StartRestoreResponseBody
}

type StartRestoreResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartRestoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartRestoreResponse) String() string {
	return dara.Prettify(s)
}

func (s StartRestoreResponse) GoString() string {
	return s.String()
}

func (s *StartRestoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartRestoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartRestoreResponse) GetBody() *StartRestoreResponseBody {
	return s.Body
}

func (s *StartRestoreResponse) SetHeaders(v map[string]*string) *StartRestoreResponse {
	s.Headers = v
	return s
}

func (s *StartRestoreResponse) SetStatusCode(v int32) *StartRestoreResponse {
	s.StatusCode = &v
	return s
}

func (s *StartRestoreResponse) SetBody(v *StartRestoreResponseBody) *StartRestoreResponse {
	s.Body = v
	return s
}

func (s *StartRestoreResponse) Validate() error {
	return dara.Validate(s)
}
