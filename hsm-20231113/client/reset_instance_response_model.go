// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetInstanceResponse
	GetStatusCode() *int32
	SetBody(v *ResetInstanceResponseBody) *ResetInstanceResponse
	GetBody() *ResetInstanceResponseBody
}

type ResetInstanceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetInstanceResponse) GoString() string {
	return s.String()
}

func (s *ResetInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetInstanceResponse) GetBody() *ResetInstanceResponseBody {
	return s.Body
}

func (s *ResetInstanceResponse) SetHeaders(v map[string]*string) *ResetInstanceResponse {
	s.Headers = v
	return s
}

func (s *ResetInstanceResponse) SetStatusCode(v int32) *ResetInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstanceResponse) SetBody(v *ResetInstanceResponseBody) *ResetInstanceResponse {
	s.Body = v
	return s
}

func (s *ResetInstanceResponse) Validate() error {
	return dara.Validate(s)
}
