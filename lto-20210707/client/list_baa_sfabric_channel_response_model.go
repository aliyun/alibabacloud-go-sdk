// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBaaSFabricChannelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBaaSFabricChannelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBaaSFabricChannelResponse
	GetStatusCode() *int32
	SetBody(v *ListBaaSFabricChannelResponseBody) *ListBaaSFabricChannelResponse
	GetBody() *ListBaaSFabricChannelResponseBody
}

type ListBaaSFabricChannelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBaaSFabricChannelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBaaSFabricChannelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBaaSFabricChannelResponse) GoString() string {
	return s.String()
}

func (s *ListBaaSFabricChannelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBaaSFabricChannelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBaaSFabricChannelResponse) GetBody() *ListBaaSFabricChannelResponseBody {
	return s.Body
}

func (s *ListBaaSFabricChannelResponse) SetHeaders(v map[string]*string) *ListBaaSFabricChannelResponse {
	s.Headers = v
	return s
}

func (s *ListBaaSFabricChannelResponse) SetStatusCode(v int32) *ListBaaSFabricChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBaaSFabricChannelResponse) SetBody(v *ListBaaSFabricChannelResponseBody) *ListBaaSFabricChannelResponse {
	s.Body = v
	return s
}

func (s *ListBaaSFabricChannelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
