// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVirtualResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVirtualResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVirtualResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVirtualResourceResponseBody) *DeleteVirtualResourceResponse
	GetBody() *DeleteVirtualResourceResponseBody
}

type DeleteVirtualResourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVirtualResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVirtualResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVirtualResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteVirtualResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVirtualResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVirtualResourceResponse) GetBody() *DeleteVirtualResourceResponseBody {
	return s.Body
}

func (s *DeleteVirtualResourceResponse) SetHeaders(v map[string]*string) *DeleteVirtualResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteVirtualResourceResponse) SetStatusCode(v int32) *DeleteVirtualResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVirtualResourceResponse) SetBody(v *DeleteVirtualResourceResponseBody) *DeleteVirtualResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteVirtualResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
