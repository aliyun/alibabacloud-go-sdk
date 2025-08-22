// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDcdnProjectExistResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckDcdnProjectExistResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckDcdnProjectExistResponse
	GetStatusCode() *int32
	SetBody(v *CheckDcdnProjectExistResponseBody) *CheckDcdnProjectExistResponse
	GetBody() *CheckDcdnProjectExistResponseBody
}

type CheckDcdnProjectExistResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckDcdnProjectExistResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckDcdnProjectExistResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckDcdnProjectExistResponse) GoString() string {
	return s.String()
}

func (s *CheckDcdnProjectExistResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckDcdnProjectExistResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckDcdnProjectExistResponse) GetBody() *CheckDcdnProjectExistResponseBody {
	return s.Body
}

func (s *CheckDcdnProjectExistResponse) SetHeaders(v map[string]*string) *CheckDcdnProjectExistResponse {
	s.Headers = v
	return s
}

func (s *CheckDcdnProjectExistResponse) SetStatusCode(v int32) *CheckDcdnProjectExistResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckDcdnProjectExistResponse) SetBody(v *CheckDcdnProjectExistResponseBody) *CheckDcdnProjectExistResponse {
	s.Body = v
	return s
}

func (s *CheckDcdnProjectExistResponse) Validate() error {
	return dara.Validate(s)
}
