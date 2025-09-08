// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse
	GetBody() *ModifyDataSourceResponseBody
}

type ModifyDataSourceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDataSourceResponse) GetBody() *ModifyDataSourceResponseBody {
	return s.Body
}

func (s *ModifyDataSourceResponse) SetHeaders(v map[string]*string) *ModifyDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyDataSourceResponse) SetStatusCode(v int32) *ModifyDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDataSourceResponse) SetBody(v *ModifyDataSourceResponseBody) *ModifyDataSourceResponse {
	s.Body = v
	return s
}

func (s *ModifyDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
