// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForPrivateAccessTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForPrivateAccessTagResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForPrivateAccessTagResponseBody) *ListApplicationsForPrivateAccessTagResponse
	GetBody() *ListApplicationsForPrivateAccessTagResponseBody
}

type ListApplicationsForPrivateAccessTagResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForPrivateAccessTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForPrivateAccessTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForPrivateAccessTagResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForPrivateAccessTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForPrivateAccessTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForPrivateAccessTagResponse) GetBody() *ListApplicationsForPrivateAccessTagResponseBody {
	return s.Body
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetHeaders(v map[string]*string) *ListApplicationsForPrivateAccessTagResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetStatusCode(v int32) *ListApplicationsForPrivateAccessTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponse) SetBody(v *ListApplicationsForPrivateAccessTagResponseBody) *ListApplicationsForPrivateAccessTagResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForPrivateAccessTagResponse) Validate() error {
	return dara.Validate(s)
}
