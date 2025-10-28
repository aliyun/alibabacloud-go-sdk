// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListResourceGroupResponseBody) *ListResourceGroupResponse
	GetBody() *ListResourceGroupResponseBody
}

type ListResourceGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ListResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListResourceGroupResponse) GetBody() *ListResourceGroupResponseBody {
	return s.Body
}

func (s *ListResourceGroupResponse) SetHeaders(v map[string]*string) *ListResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ListResourceGroupResponse) SetStatusCode(v int32) *ListResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListResourceGroupResponse) SetBody(v *ListResourceGroupResponseBody) *ListResourceGroupResponse {
	s.Body = v
	return s
}

func (s *ListResourceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
