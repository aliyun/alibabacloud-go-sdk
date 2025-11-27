// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCNodePoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCNodePoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCNodePoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCNodePoolResponseBody) *DeleteRCNodePoolResponse
	GetBody() *DeleteRCNodePoolResponseBody
}

type DeleteRCNodePoolResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCNodePoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCNodePoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCNodePoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCNodePoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCNodePoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCNodePoolResponse) GetBody() *DeleteRCNodePoolResponseBody {
	return s.Body
}

func (s *DeleteRCNodePoolResponse) SetHeaders(v map[string]*string) *DeleteRCNodePoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCNodePoolResponse) SetStatusCode(v int32) *DeleteRCNodePoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCNodePoolResponse) SetBody(v *DeleteRCNodePoolResponseBody) *DeleteRCNodePoolResponse {
	s.Body = v
	return s
}

func (s *DeleteRCNodePoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
