// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualHostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualHostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualHostResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualHostResponseBody) *DeleteVirtualHostResponse
	GetBody() *DeleteVirtualHostResponseBody
}

type DeleteVirtualHostResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualHostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualHostResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualHostResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualHostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualHostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualHostResponse) GetBody() *DeleteVirtualHostResponseBody {
	return s.Body
}

func (s *DeleteVirtualHostResponse) SetHeaders(v map[string]*string) *DeleteVirtualHostResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualHostResponse) SetStatusCode(v int32) *DeleteVirtualHostResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualHostResponse) SetBody(v *DeleteVirtualHostResponseBody) *DeleteVirtualHostResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualHostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
