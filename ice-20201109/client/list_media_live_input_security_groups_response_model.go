// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaLiveInputSecurityGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaLiveInputSecurityGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaLiveInputSecurityGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaLiveInputSecurityGroupsResponseBody) *ListMediaLiveInputSecurityGroupsResponse
	GetBody() *ListMediaLiveInputSecurityGroupsResponseBody
}

type ListMediaLiveInputSecurityGroupsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaLiveInputSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaLiveInputSecurityGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaLiveInputSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListMediaLiveInputSecurityGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaLiveInputSecurityGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaLiveInputSecurityGroupsResponse) GetBody() *ListMediaLiveInputSecurityGroupsResponseBody {
	return s.Body
}

func (s *ListMediaLiveInputSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListMediaLiveInputSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponse) SetStatusCode(v int32) *ListMediaLiveInputSecurityGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponse) SetBody(v *ListMediaLiveInputSecurityGroupsResponseBody) *ListMediaLiveInputSecurityGroupsResponse {
	s.Body = v
	return s
}

func (s *ListMediaLiveInputSecurityGroupsResponse) Validate() error {
	return dara.Validate(s)
}
