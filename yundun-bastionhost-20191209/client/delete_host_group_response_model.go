// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHostGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHostGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHostGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse
	GetBody() *DeleteHostGroupResponseBody
}

type DeleteHostGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHostGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHostGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHostGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteHostGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHostGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHostGroupResponse) GetBody() *DeleteHostGroupResponseBody {
	return s.Body
}

func (s *DeleteHostGroupResponse) SetHeaders(v map[string]*string) *DeleteHostGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteHostGroupResponse) SetStatusCode(v int32) *DeleteHostGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHostGroupResponse) SetBody(v *DeleteHostGroupResponseBody) *DeleteHostGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteHostGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
