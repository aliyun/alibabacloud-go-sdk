// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAdminsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportAdminsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportAdminsResponse
	GetStatusCode() *int32
	SetBody(v *ImportAdminsResponseBody) *ImportAdminsResponse
	GetBody() *ImportAdminsResponseBody
}

type ImportAdminsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportAdminsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportAdminsResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportAdminsResponse) GoString() string {
	return s.String()
}

func (s *ImportAdminsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportAdminsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportAdminsResponse) GetBody() *ImportAdminsResponseBody {
	return s.Body
}

func (s *ImportAdminsResponse) SetHeaders(v map[string]*string) *ImportAdminsResponse {
	s.Headers = v
	return s
}

func (s *ImportAdminsResponse) SetStatusCode(v int32) *ImportAdminsResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAdminsResponse) SetBody(v *ImportAdminsResponseBody) *ImportAdminsResponse {
	s.Body = v
	return s
}

func (s *ImportAdminsResponse) Validate() error {
	return dara.Validate(s)
}
