// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RollbackApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RollbackApplicationResponse
	GetStatusCode() *int32
	SetBody(v *RollbackApplicationResponseBody) *RollbackApplicationResponse
	GetBody() *RollbackApplicationResponseBody
}

type RollbackApplicationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RollbackApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RollbackApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationResponse) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RollbackApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RollbackApplicationResponse) GetBody() *RollbackApplicationResponseBody {
	return s.Body
}

func (s *RollbackApplicationResponse) SetHeaders(v map[string]*string) *RollbackApplicationResponse {
	s.Headers = v
	return s
}

func (s *RollbackApplicationResponse) SetStatusCode(v int32) *RollbackApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *RollbackApplicationResponse) SetBody(v *RollbackApplicationResponseBody) *RollbackApplicationResponse {
	s.Body = v
	return s
}

func (s *RollbackApplicationResponse) Validate() error {
	return dara.Validate(s)
}
