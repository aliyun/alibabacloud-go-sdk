// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHadoopDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHadoopDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHadoopDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHadoopDataSourceResponseBody) *ModifyHadoopDataSourceResponse
	GetBody() *ModifyHadoopDataSourceResponseBody
}

type ModifyHadoopDataSourceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHadoopDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHadoopDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHadoopDataSourceResponse) GoString() string {
	return s.String()
}

func (s *ModifyHadoopDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHadoopDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHadoopDataSourceResponse) GetBody() *ModifyHadoopDataSourceResponseBody {
	return s.Body
}

func (s *ModifyHadoopDataSourceResponse) SetHeaders(v map[string]*string) *ModifyHadoopDataSourceResponse {
	s.Headers = v
	return s
}

func (s *ModifyHadoopDataSourceResponse) SetStatusCode(v int32) *ModifyHadoopDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHadoopDataSourceResponse) SetBody(v *ModifyHadoopDataSourceResponseBody) *ModifyHadoopDataSourceResponse {
	s.Body = v
	return s
}

func (s *ModifyHadoopDataSourceResponse) Validate() error {
	return dara.Validate(s)
}
