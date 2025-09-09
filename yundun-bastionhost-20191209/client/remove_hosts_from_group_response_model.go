// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHostsFromGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveHostsFromGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveHostsFromGroupResponse
	GetStatusCode() *int32
	SetBody(v *RemoveHostsFromGroupResponseBody) *RemoveHostsFromGroupResponse
	GetBody() *RemoveHostsFromGroupResponseBody
}

type RemoveHostsFromGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveHostsFromGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveHostsFromGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveHostsFromGroupResponse) GoString() string {
	return s.String()
}

func (s *RemoveHostsFromGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveHostsFromGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveHostsFromGroupResponse) GetBody() *RemoveHostsFromGroupResponseBody {
	return s.Body
}

func (s *RemoveHostsFromGroupResponse) SetHeaders(v map[string]*string) *RemoveHostsFromGroupResponse {
	s.Headers = v
	return s
}

func (s *RemoveHostsFromGroupResponse) SetStatusCode(v int32) *RemoveHostsFromGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveHostsFromGroupResponse) SetBody(v *RemoveHostsFromGroupResponseBody) *RemoveHostsFromGroupResponse {
	s.Body = v
	return s
}

func (s *RemoveHostsFromGroupResponse) Validate() error {
	return dara.Validate(s)
}
