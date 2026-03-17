// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudConnectNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCloudConnectNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCloudConnectNetworkResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCloudConnectNetworkResponseBody) *ModifyCloudConnectNetworkResponse
	GetBody() *ModifyCloudConnectNetworkResponseBody
}

type ModifyCloudConnectNetworkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCloudConnectNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCloudConnectNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudConnectNetworkResponse) GoString() string {
	return s.String()
}

func (s *ModifyCloudConnectNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCloudConnectNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCloudConnectNetworkResponse) GetBody() *ModifyCloudConnectNetworkResponseBody {
	return s.Body
}

func (s *ModifyCloudConnectNetworkResponse) SetHeaders(v map[string]*string) *ModifyCloudConnectNetworkResponse {
	s.Headers = v
	return s
}

func (s *ModifyCloudConnectNetworkResponse) SetStatusCode(v int32) *ModifyCloudConnectNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCloudConnectNetworkResponse) SetBody(v *ModifyCloudConnectNetworkResponseBody) *ModifyCloudConnectNetworkResponse {
	s.Body = v
	return s
}

func (s *ModifyCloudConnectNetworkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
