// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGlobalDataNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGlobalDataNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGlobalDataNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGlobalDataNetworkResponseBody) *DeleteGlobalDataNetworkResponse
	GetBody() *DeleteGlobalDataNetworkResponseBody
}

type DeleteGlobalDataNetworkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGlobalDataNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGlobalDataNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGlobalDataNetworkResponse) GoString() string {
	return s.String()
}

func (s *DeleteGlobalDataNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGlobalDataNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGlobalDataNetworkResponse) GetBody() *DeleteGlobalDataNetworkResponseBody {
	return s.Body
}

func (s *DeleteGlobalDataNetworkResponse) SetHeaders(v map[string]*string) *DeleteGlobalDataNetworkResponse {
	s.Headers = v
	return s
}

func (s *DeleteGlobalDataNetworkResponse) SetStatusCode(v int32) *DeleteGlobalDataNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGlobalDataNetworkResponse) SetBody(v *DeleteGlobalDataNetworkResponseBody) *DeleteGlobalDataNetworkResponse {
	s.Body = v
	return s
}

func (s *DeleteGlobalDataNetworkResponse) Validate() error {
	return dara.Validate(s)
}
