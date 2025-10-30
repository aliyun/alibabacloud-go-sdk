// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckHadoopDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckHadoopDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckHadoopDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *CheckHadoopDataSourceResponseBody) *CheckHadoopDataSourceResponse
	GetBody() *CheckHadoopDataSourceResponseBody
}

type CheckHadoopDataSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckHadoopDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckHadoopDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckHadoopDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CheckHadoopDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckHadoopDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckHadoopDataSourceResponse) GetBody() *CheckHadoopDataSourceResponseBody {
	return s.Body
}

func (s *CheckHadoopDataSourceResponse) SetHeaders(v map[string]*string) *CheckHadoopDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CheckHadoopDataSourceResponse) SetStatusCode(v int32) *CheckHadoopDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckHadoopDataSourceResponse) SetBody(v *CheckHadoopDataSourceResponseBody) *CheckHadoopDataSourceResponse {
	s.Body = v
	return s
}

func (s *CheckHadoopDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
