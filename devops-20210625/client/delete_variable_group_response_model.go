// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVariableGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVariableGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVariableGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVariableGroupResponseBody) *DeleteVariableGroupResponse
	GetBody() *DeleteVariableGroupResponseBody
}

type DeleteVariableGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVariableGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVariableGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVariableGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteVariableGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVariableGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVariableGroupResponse) GetBody() *DeleteVariableGroupResponseBody {
	return s.Body
}

func (s *DeleteVariableGroupResponse) SetHeaders(v map[string]*string) *DeleteVariableGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteVariableGroupResponse) SetStatusCode(v int32) *DeleteVariableGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVariableGroupResponse) SetBody(v *DeleteVariableGroupResponseBody) *DeleteVariableGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteVariableGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
