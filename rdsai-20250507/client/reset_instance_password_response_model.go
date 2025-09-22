// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetInstancePasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetInstancePasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetInstancePasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetInstancePasswordResponseBody) *ResetInstancePasswordResponse
	GetBody() *ResetInstancePasswordResponseBody
}

type ResetInstancePasswordResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetInstancePasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetInstancePasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetInstancePasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetInstancePasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetInstancePasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetInstancePasswordResponse) GetBody() *ResetInstancePasswordResponseBody {
	return s.Body
}

func (s *ResetInstancePasswordResponse) SetHeaders(v map[string]*string) *ResetInstancePasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetInstancePasswordResponse) SetStatusCode(v int32) *ResetInstancePasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetInstancePasswordResponse) SetBody(v *ResetInstancePasswordResponseBody) *ResetInstancePasswordResponse {
	s.Body = v
	return s
}

func (s *ResetInstancePasswordResponse) Validate() error {
	return dara.Validate(s)
}
