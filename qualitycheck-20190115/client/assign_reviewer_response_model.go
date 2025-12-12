// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssignReviewerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssignReviewerResponse
	GetStatusCode() *int32
	SetBody(v *AssignReviewerResponseBody) *AssignReviewerResponse
	GetBody() *AssignReviewerResponseBody
}

type AssignReviewerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssignReviewerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssignReviewerResponse) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerResponse) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssignReviewerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssignReviewerResponse) GetBody() *AssignReviewerResponseBody {
	return s.Body
}

func (s *AssignReviewerResponse) SetHeaders(v map[string]*string) *AssignReviewerResponse {
	s.Headers = v
	return s
}

func (s *AssignReviewerResponse) SetStatusCode(v int32) *AssignReviewerResponse {
	s.StatusCode = &v
	return s
}

func (s *AssignReviewerResponse) SetBody(v *AssignReviewerResponseBody) *AssignReviewerResponse {
	s.Body = v
	return s
}

func (s *AssignReviewerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
