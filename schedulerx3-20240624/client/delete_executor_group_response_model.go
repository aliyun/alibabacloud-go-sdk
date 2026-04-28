// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExecutorGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExecutorGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExecutorGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExecutorGroupResponseBody) *DeleteExecutorGroupResponse
	GetBody() *DeleteExecutorGroupResponseBody
}

type DeleteExecutorGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExecutorGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExecutorGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExecutorGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteExecutorGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExecutorGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExecutorGroupResponse) GetBody() *DeleteExecutorGroupResponseBody {
	return s.Body
}

func (s *DeleteExecutorGroupResponse) SetHeaders(v map[string]*string) *DeleteExecutorGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteExecutorGroupResponse) SetStatusCode(v int32) *DeleteExecutorGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExecutorGroupResponse) SetBody(v *DeleteExecutorGroupResponseBody) *DeleteExecutorGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteExecutorGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
