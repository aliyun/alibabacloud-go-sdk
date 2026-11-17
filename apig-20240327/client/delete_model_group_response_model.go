// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteModelGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteModelGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteModelGroupResponseBody) *DeleteModelGroupResponse
	GetBody() *DeleteModelGroupResponseBody
}

type DeleteModelGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteModelGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteModelGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteModelGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteModelGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteModelGroupResponse) GetBody() *DeleteModelGroupResponseBody {
	return s.Body
}

func (s *DeleteModelGroupResponse) SetHeaders(v map[string]*string) *DeleteModelGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteModelGroupResponse) SetStatusCode(v int32) *DeleteModelGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteModelGroupResponse) SetBody(v *DeleteModelGroupResponseBody) *DeleteModelGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteModelGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
