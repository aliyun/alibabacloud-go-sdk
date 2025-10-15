// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForGroupResponseBody) *ListApplicationsForGroupResponse
	GetBody() *ListApplicationsForGroupResponseBody
}

type ListApplicationsForGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForGroupResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForGroupResponse) GetBody() *ListApplicationsForGroupResponseBody {
	return s.Body
}

func (s *ListApplicationsForGroupResponse) SetHeaders(v map[string]*string) *ListApplicationsForGroupResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForGroupResponse) SetStatusCode(v int32) *ListApplicationsForGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForGroupResponse) SetBody(v *ListApplicationsForGroupResponseBody) *ListApplicationsForGroupResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
