// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDepartmentSaveResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DepartmentSaveResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DepartmentSaveResponse
	GetStatusCode() *int32
	SetBody(v *DepartmentSaveResponseBody) *DepartmentSaveResponse
	GetBody() *DepartmentSaveResponseBody
}

type DepartmentSaveResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DepartmentSaveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DepartmentSaveResponse) String() string {
	return dara.Prettify(s)
}

func (s DepartmentSaveResponse) GoString() string {
	return s.String()
}

func (s *DepartmentSaveResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DepartmentSaveResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DepartmentSaveResponse) GetBody() *DepartmentSaveResponseBody {
	return s.Body
}

func (s *DepartmentSaveResponse) SetHeaders(v map[string]*string) *DepartmentSaveResponse {
	s.Headers = v
	return s
}

func (s *DepartmentSaveResponse) SetStatusCode(v int32) *DepartmentSaveResponse {
	s.StatusCode = &v
	return s
}

func (s *DepartmentSaveResponse) SetBody(v *DepartmentSaveResponseBody) *DepartmentSaveResponse {
	s.Body = v
	return s
}

func (s *DepartmentSaveResponse) Validate() error {
	return dara.Validate(s)
}
