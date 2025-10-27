// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateLogHubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateLogHubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateLogHubResponse
	GetStatusCode() *int32
	SetBody(v *OperateLogHubResponseBody) *OperateLogHubResponse
	GetBody() *OperateLogHubResponseBody
}

type OperateLogHubResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateLogHubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateLogHubResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateLogHubResponse) GoString() string {
	return s.String()
}

func (s *OperateLogHubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateLogHubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateLogHubResponse) GetBody() *OperateLogHubResponseBody {
	return s.Body
}

func (s *OperateLogHubResponse) SetHeaders(v map[string]*string) *OperateLogHubResponse {
	s.Headers = v
	return s
}

func (s *OperateLogHubResponse) SetStatusCode(v int32) *OperateLogHubResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateLogHubResponse) SetBody(v *OperateLogHubResponseBody) *OperateLogHubResponse {
	s.Body = v
	return s
}

func (s *OperateLogHubResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
