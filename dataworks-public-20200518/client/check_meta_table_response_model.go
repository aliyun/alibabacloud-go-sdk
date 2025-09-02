// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckMetaTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckMetaTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckMetaTableResponse
	GetStatusCode() *int32
	SetBody(v *CheckMetaTableResponseBody) *CheckMetaTableResponse
	GetBody() *CheckMetaTableResponseBody
}

type CheckMetaTableResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckMetaTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckMetaTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckMetaTableResponse) GoString() string {
	return s.String()
}

func (s *CheckMetaTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckMetaTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckMetaTableResponse) GetBody() *CheckMetaTableResponseBody {
	return s.Body
}

func (s *CheckMetaTableResponse) SetHeaders(v map[string]*string) *CheckMetaTableResponse {
	s.Headers = v
	return s
}

func (s *CheckMetaTableResponse) SetStatusCode(v int32) *CheckMetaTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckMetaTableResponse) SetBody(v *CheckMetaTableResponseBody) *CheckMetaTableResponse {
	s.Body = v
	return s
}

func (s *CheckMetaTableResponse) Validate() error {
	return dara.Validate(s)
}
