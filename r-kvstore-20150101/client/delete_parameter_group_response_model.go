// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteParameterGroupResponseBody) *DeleteParameterGroupResponse
	GetBody() *DeleteParameterGroupResponseBody
}

type DeleteParameterGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteParameterGroupResponse) GetBody() *DeleteParameterGroupResponseBody {
	return s.Body
}

func (s *DeleteParameterGroupResponse) SetHeaders(v map[string]*string) *DeleteParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteParameterGroupResponse) SetStatusCode(v int32) *DeleteParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteParameterGroupResponse) SetBody(v *DeleteParameterGroupResponseBody) *DeleteParameterGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteParameterGroupResponse) Validate() error {
	return dara.Validate(s)
}
