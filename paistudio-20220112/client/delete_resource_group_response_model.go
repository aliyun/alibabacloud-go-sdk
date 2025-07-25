// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse
	GetBody() *DeleteResourceGroupResponseBody
}

type DeleteResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteResourceGroupResponse) GetBody() *DeleteResourceGroupResponseBody {
	return s.Body
}

func (s *DeleteResourceGroupResponse) SetHeaders(v map[string]*string) *DeleteResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteResourceGroupResponse) SetStatusCode(v int32) *DeleteResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteResourceGroupResponse) SetBody(v *DeleteResourceGroupResponseBody) *DeleteResourceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}
