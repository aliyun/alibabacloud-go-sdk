// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneParameterGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneParameterGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneParameterGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloneParameterGroupResponseBody) *CloneParameterGroupResponse
	GetBody() *CloneParameterGroupResponseBody
}

type CloneParameterGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneParameterGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneParameterGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneParameterGroupResponse) GoString() string {
	return s.String()
}

func (s *CloneParameterGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneParameterGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneParameterGroupResponse) GetBody() *CloneParameterGroupResponseBody {
	return s.Body
}

func (s *CloneParameterGroupResponse) SetHeaders(v map[string]*string) *CloneParameterGroupResponse {
	s.Headers = v
	return s
}

func (s *CloneParameterGroupResponse) SetStatusCode(v int32) *CloneParameterGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneParameterGroupResponse) SetBody(v *CloneParameterGroupResponseBody) *CloneParameterGroupResponse {
	s.Body = v
	return s
}

func (s *CloneParameterGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
