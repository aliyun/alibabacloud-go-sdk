// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiagnoseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiagnoseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DiagnoseInstanceResponseBody) *DiagnoseInstanceResponse
	GetBody() *DiagnoseInstanceResponseBody
}

type DiagnoseInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiagnoseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiagnoseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseInstanceResponse) GoString() string {
	return s.String()
}

func (s *DiagnoseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiagnoseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiagnoseInstanceResponse) GetBody() *DiagnoseInstanceResponseBody {
	return s.Body
}

func (s *DiagnoseInstanceResponse) SetHeaders(v map[string]*string) *DiagnoseInstanceResponse {
	s.Headers = v
	return s
}

func (s *DiagnoseInstanceResponse) SetStatusCode(v int32) *DiagnoseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DiagnoseInstanceResponse) SetBody(v *DiagnoseInstanceResponseBody) *DiagnoseInstanceResponse {
	s.Body = v
	return s
}

func (s *DiagnoseInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
