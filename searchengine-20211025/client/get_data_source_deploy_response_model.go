// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataSourceDeployResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDataSourceDeployResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDataSourceDeployResponse
	GetStatusCode() *int32
	SetBody(v *GetDataSourceDeployResponseBody) *GetDataSourceDeployResponse
	GetBody() *GetDataSourceDeployResponseBody
}

type GetDataSourceDeployResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDataSourceDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDataSourceDeployResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDataSourceDeployResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceDeployResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDataSourceDeployResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDataSourceDeployResponse) GetBody() *GetDataSourceDeployResponseBody {
	return s.Body
}

func (s *GetDataSourceDeployResponse) SetHeaders(v map[string]*string) *GetDataSourceDeployResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceDeployResponse) SetStatusCode(v int32) *GetDataSourceDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceDeployResponse) SetBody(v *GetDataSourceDeployResponseBody) *GetDataSourceDeployResponse {
	s.Body = v
	return s
}

func (s *GetDataSourceDeployResponse) Validate() error {
	return dara.Validate(s)
}
