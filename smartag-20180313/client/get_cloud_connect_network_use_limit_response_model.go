// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudConnectNetworkUseLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCloudConnectNetworkUseLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCloudConnectNetworkUseLimitResponse
	GetStatusCode() *int32
	SetBody(v *GetCloudConnectNetworkUseLimitResponseBody) *GetCloudConnectNetworkUseLimitResponse
	GetBody() *GetCloudConnectNetworkUseLimitResponseBody
}

type GetCloudConnectNetworkUseLimitResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCloudConnectNetworkUseLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCloudConnectNetworkUseLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCloudConnectNetworkUseLimitResponse) GoString() string {
	return s.String()
}

func (s *GetCloudConnectNetworkUseLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCloudConnectNetworkUseLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCloudConnectNetworkUseLimitResponse) GetBody() *GetCloudConnectNetworkUseLimitResponseBody {
	return s.Body
}

func (s *GetCloudConnectNetworkUseLimitResponse) SetHeaders(v map[string]*string) *GetCloudConnectNetworkUseLimitResponse {
	s.Headers = v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponse) SetStatusCode(v int32) *GetCloudConnectNetworkUseLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponse) SetBody(v *GetCloudConnectNetworkUseLimitResponseBody) *GetCloudConnectNetworkUseLimitResponse {
	s.Body = v
	return s
}

func (s *GetCloudConnectNetworkUseLimitResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
