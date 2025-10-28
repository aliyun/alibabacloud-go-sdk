// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertDeployGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertDeployGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertDeployGroupResponse
	GetStatusCode() *int32
	SetBody(v *InsertDeployGroupResponseBody) *InsertDeployGroupResponse
	GetBody() *InsertDeployGroupResponseBody
}

type InsertDeployGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertDeployGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertDeployGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertDeployGroupResponse) GoString() string {
	return s.String()
}

func (s *InsertDeployGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertDeployGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertDeployGroupResponse) GetBody() *InsertDeployGroupResponseBody {
	return s.Body
}

func (s *InsertDeployGroupResponse) SetHeaders(v map[string]*string) *InsertDeployGroupResponse {
	s.Headers = v
	return s
}

func (s *InsertDeployGroupResponse) SetStatusCode(v int32) *InsertDeployGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertDeployGroupResponse) SetBody(v *InsertDeployGroupResponseBody) *InsertDeployGroupResponse {
	s.Body = v
	return s
}

func (s *InsertDeployGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
