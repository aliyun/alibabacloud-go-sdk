// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClonePolicyGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClonePolicyGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClonePolicyGroupResponse
	GetStatusCode() *int32
	SetBody(v *ClonePolicyGroupResponseBody) *ClonePolicyGroupResponse
	GetBody() *ClonePolicyGroupResponseBody
}

type ClonePolicyGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClonePolicyGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClonePolicyGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ClonePolicyGroupResponse) GoString() string {
	return s.String()
}

func (s *ClonePolicyGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClonePolicyGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClonePolicyGroupResponse) GetBody() *ClonePolicyGroupResponseBody {
	return s.Body
}

func (s *ClonePolicyGroupResponse) SetHeaders(v map[string]*string) *ClonePolicyGroupResponse {
	s.Headers = v
	return s
}

func (s *ClonePolicyGroupResponse) SetStatusCode(v int32) *ClonePolicyGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ClonePolicyGroupResponse) SetBody(v *ClonePolicyGroupResponseBody) *ClonePolicyGroupResponse {
	s.Body = v
	return s
}

func (s *ClonePolicyGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
