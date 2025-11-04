// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachServerGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachServerGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachServerGroupsResponse
	GetStatusCode() *int32
	SetBody(v *AttachServerGroupsResponseBody) *AttachServerGroupsResponse
	GetBody() *AttachServerGroupsResponseBody
}

type AttachServerGroupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachServerGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachServerGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachServerGroupsResponse) GoString() string {
	return s.String()
}

func (s *AttachServerGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachServerGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachServerGroupsResponse) GetBody() *AttachServerGroupsResponseBody {
	return s.Body
}

func (s *AttachServerGroupsResponse) SetHeaders(v map[string]*string) *AttachServerGroupsResponse {
	s.Headers = v
	return s
}

func (s *AttachServerGroupsResponse) SetStatusCode(v int32) *AttachServerGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachServerGroupsResponse) SetBody(v *AttachServerGroupsResponseBody) *AttachServerGroupsResponse {
	s.Body = v
	return s
}

func (s *AttachServerGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
