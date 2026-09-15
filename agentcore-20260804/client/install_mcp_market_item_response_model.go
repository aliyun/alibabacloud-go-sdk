// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallMcpMarketItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InstallMcpMarketItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InstallMcpMarketItemResponse
	GetStatusCode() *int32
	SetBody(v *InstallMcpMarketItemResponseBody) *InstallMcpMarketItemResponse
	GetBody() *InstallMcpMarketItemResponseBody
}

type InstallMcpMarketItemResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InstallMcpMarketItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InstallMcpMarketItemResponse) String() string {
	return dara.Prettify(s)
}

func (s InstallMcpMarketItemResponse) GoString() string {
	return s.String()
}

func (s *InstallMcpMarketItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InstallMcpMarketItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InstallMcpMarketItemResponse) GetBody() *InstallMcpMarketItemResponseBody {
	return s.Body
}

func (s *InstallMcpMarketItemResponse) SetHeaders(v map[string]*string) *InstallMcpMarketItemResponse {
	s.Headers = v
	return s
}

func (s *InstallMcpMarketItemResponse) SetStatusCode(v int32) *InstallMcpMarketItemResponse {
	s.StatusCode = &v
	return s
}

func (s *InstallMcpMarketItemResponse) SetBody(v *InstallMcpMarketItemResponseBody) *InstallMcpMarketItemResponse {
	s.Body = v
	return s
}

func (s *InstallMcpMarketItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
