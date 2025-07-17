// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNetworkResponse
	GetStatusCode() *int32
	SetBody(v *GetNetworkResponseBody) *GetNetworkResponse
	GetBody() *GetNetworkResponseBody
}

type GetNetworkResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkResponse) GoString() string {
	return s.String()
}

func (s *GetNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNetworkResponse) GetBody() *GetNetworkResponseBody {
	return s.Body
}

func (s *GetNetworkResponse) SetHeaders(v map[string]*string) *GetNetworkResponse {
	s.Headers = v
	return s
}

func (s *GetNetworkResponse) SetStatusCode(v int32) *GetNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNetworkResponse) SetBody(v *GetNetworkResponseBody) *GetNetworkResponse {
	s.Body = v
	return s
}

func (s *GetNetworkResponse) Validate() error {
	return dara.Validate(s)
}
