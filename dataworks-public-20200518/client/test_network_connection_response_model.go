// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestNetworkConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestNetworkConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestNetworkConnectionResponse
	GetStatusCode() *int32
	SetBody(v *TestNetworkConnectionResponseBody) *TestNetworkConnectionResponse
	GetBody() *TestNetworkConnectionResponseBody
}

type TestNetworkConnectionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestNetworkConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestNetworkConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s TestNetworkConnectionResponse) GoString() string {
	return s.String()
}

func (s *TestNetworkConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestNetworkConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestNetworkConnectionResponse) GetBody() *TestNetworkConnectionResponseBody {
	return s.Body
}

func (s *TestNetworkConnectionResponse) SetHeaders(v map[string]*string) *TestNetworkConnectionResponse {
	s.Headers = v
	return s
}

func (s *TestNetworkConnectionResponse) SetStatusCode(v int32) *TestNetworkConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *TestNetworkConnectionResponse) SetBody(v *TestNetworkConnectionResponseBody) *TestNetworkConnectionResponse {
	s.Body = v
	return s
}

func (s *TestNetworkConnectionResponse) Validate() error {
	return dara.Validate(s)
}
