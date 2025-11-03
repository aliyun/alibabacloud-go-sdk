// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseVpnConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DiagnoseVpnConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DiagnoseVpnConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *DiagnoseVpnConnectionsResponseBody) *DiagnoseVpnConnectionsResponse
	GetBody() *DiagnoseVpnConnectionsResponseBody
}

type DiagnoseVpnConnectionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DiagnoseVpnConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DiagnoseVpnConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseVpnConnectionsResponse) GoString() string {
	return s.String()
}

func (s *DiagnoseVpnConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DiagnoseVpnConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DiagnoseVpnConnectionsResponse) GetBody() *DiagnoseVpnConnectionsResponseBody {
	return s.Body
}

func (s *DiagnoseVpnConnectionsResponse) SetHeaders(v map[string]*string) *DiagnoseVpnConnectionsResponse {
	s.Headers = v
	return s
}

func (s *DiagnoseVpnConnectionsResponse) SetStatusCode(v int32) *DiagnoseVpnConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DiagnoseVpnConnectionsResponse) SetBody(v *DiagnoseVpnConnectionsResponseBody) *DiagnoseVpnConnectionsResponse {
	s.Body = v
	return s
}

func (s *DiagnoseVpnConnectionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
