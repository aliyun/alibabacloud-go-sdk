// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *AddDataSourceResponseBody) *AddDataSourceResponse
	GetBody() *AddDataSourceResponseBody
}

type AddDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceResponse) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddDataSourceResponse) GetBody() *AddDataSourceResponseBody {
	return s.Body
}

func (s *AddDataSourceResponse) SetHeaders(v map[string]*string) *AddDataSourceResponse {
	s.Headers = v
	return s
}

func (s *AddDataSourceResponse) SetStatusCode(v int32) *AddDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *AddDataSourceResponse) SetBody(v *AddDataSourceResponseBody) *AddDataSourceResponse {
	s.Body = v
	return s
}

func (s *AddDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
