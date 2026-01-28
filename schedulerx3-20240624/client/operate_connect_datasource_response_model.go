// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateConnectDatasourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateConnectDatasourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateConnectDatasourceResponse
	GetStatusCode() *int32
	SetBody(v *OperateConnectDatasourceResponseBody) *OperateConnectDatasourceResponse
	GetBody() *OperateConnectDatasourceResponseBody
}

type OperateConnectDatasourceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateConnectDatasourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateConnectDatasourceResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateConnectDatasourceResponse) GoString() string {
	return s.String()
}

func (s *OperateConnectDatasourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateConnectDatasourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateConnectDatasourceResponse) GetBody() *OperateConnectDatasourceResponseBody {
	return s.Body
}

func (s *OperateConnectDatasourceResponse) SetHeaders(v map[string]*string) *OperateConnectDatasourceResponse {
	s.Headers = v
	return s
}

func (s *OperateConnectDatasourceResponse) SetStatusCode(v int32) *OperateConnectDatasourceResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateConnectDatasourceResponse) SetBody(v *OperateConnectDatasourceResponseBody) *OperateConnectDatasourceResponse {
	s.Body = v
	return s
}

func (s *OperateConnectDatasourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
