// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceDeployResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataSourceDeployResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataSourceDeployResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataSourceDeployResponseBody) *ModifyDataSourceDeployResponse
	GetBody() *ModifyDataSourceDeployResponseBody
}

type ModifyDataSourceDeployResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceDeployResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceDeployResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceDeployResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceDeployResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataSourceDeployResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataSourceDeployResponse) GetBody() *ModifyDataSourceDeployResponseBody {
	return s.Body
}

func (s *ModifyDataSourceDeployResponse) SetHeaders(v map[string]*string) *ModifyDataSourceDeployResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceDeployResponse) SetStatusCode(v int32) *ModifyDataSourceDeployResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceDeployResponse) SetBody(v *ModifyDataSourceDeployResponseBody) *ModifyDataSourceDeployResponse {
	s.Body = v
	return s
}

func (s *ModifyDataSourceDeployResponse) Validate() error {
	return dara.Validate(s)
}
