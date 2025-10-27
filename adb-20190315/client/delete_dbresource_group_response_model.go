// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDBResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDBResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDBResourceGroupResponseBody) *DeleteDBResourceGroupResponse
	GetBody() *DeleteDBResourceGroupResponseBody
}

type DeleteDBResourceGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDBResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDBResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDBResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDBResourceGroupResponse) GetBody() *DeleteDBResourceGroupResponseBody {
	return s.Body
}

func (s *DeleteDBResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteDBResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetStatusCode(v int32) *DeleteDBResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDBResourceGroupResponse) SetBody(v *DeleteDBResourceGroupResponseBody) *DeleteDBResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteDBResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
