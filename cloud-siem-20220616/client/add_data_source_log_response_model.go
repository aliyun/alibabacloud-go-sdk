// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataSourceLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataSourceLogResponse
	GetStatusCode() *int32
	SetBody(v *AddDataSourceLogResponseBody) *AddDataSourceLogResponse
	GetBody() *AddDataSourceLogResponseBody
}

type AddDataSourceLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataSourceLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataSourceLogResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceLogResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataSourceLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataSourceLogResponse) GetBody() *AddDataSourceLogResponseBody {
	return s.Body
}

func (s *AddDataSourceLogResponse) SetHeaders(v map[string]*string) *AddDataSourceLogResponse {
	s.Headers = v
	return s
}

func (s *AddDataSourceLogResponse) SetStatusCode(v int32) *AddDataSourceLogResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataSourceLogResponse) SetBody(v *AddDataSourceLogResponseBody) *AddDataSourceLogResponse {
	s.Body = v
	return s
}

func (s *AddDataSourceLogResponse) Validate() error {
	return dara.Validate(s)
}
