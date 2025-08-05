// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMmsDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMmsDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMmsDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CreateMmsDataSourceResponseBody) *CreateMmsDataSourceResponse
	GetBody() *CreateMmsDataSourceResponseBody
}

type CreateMmsDataSourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMmsDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMmsDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMmsDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateMmsDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMmsDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMmsDataSourceResponse) GetBody() *CreateMmsDataSourceResponseBody {
	return s.Body
}

func (s *CreateMmsDataSourceResponse) SetHeaders(v map[string]*string) *CreateMmsDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateMmsDataSourceResponse) SetStatusCode(v int32) *CreateMmsDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMmsDataSourceResponse) SetBody(v *CreateMmsDataSourceResponseBody) *CreateMmsDataSourceResponse {
	s.Body = v
	return s
}

func (s *CreateMmsDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
