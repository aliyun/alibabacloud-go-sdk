// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerBySessionGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignReviewerBySessionGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignReviewerBySessionGroupResponse
	GetStatusCode() *int32
	SetBody(v *AssignReviewerBySessionGroupResponseBody) *AssignReviewerBySessionGroupResponse
	GetBody() *AssignReviewerBySessionGroupResponseBody
}

type AssignReviewerBySessionGroupResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignReviewerBySessionGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignReviewerBySessionGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerBySessionGroupResponse) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignReviewerBySessionGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignReviewerBySessionGroupResponse) GetBody() *AssignReviewerBySessionGroupResponseBody {
	return s.Body
}

func (s *AssignReviewerBySessionGroupResponse) SetHeaders(v map[string]*string) *AssignReviewerBySessionGroupResponse {
	s.Headers = v
	return s
}

func (s *AssignReviewerBySessionGroupResponse) SetStatusCode(v int32) *AssignReviewerBySessionGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignReviewerBySessionGroupResponse) SetBody(v *AssignReviewerBySessionGroupResponseBody) *AssignReviewerBySessionGroupResponse {
	s.Body = v
	return s
}

func (s *AssignReviewerBySessionGroupResponse) Validate() error {
	return dara.Validate(s)
}
