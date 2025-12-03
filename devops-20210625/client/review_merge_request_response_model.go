// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReviewMergeRequestResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReviewMergeRequestResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReviewMergeRequestResponse
	GetStatusCode() *int32
	SetBody(v *ReviewMergeRequestResponseBody) *ReviewMergeRequestResponse
	GetBody() *ReviewMergeRequestResponseBody
}

type ReviewMergeRequestResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReviewMergeRequestResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReviewMergeRequestResponse) String() string {
	return dara.Prettify(s)
}

func (s ReviewMergeRequestResponse) GoString() string {
	return s.String()
}

func (s *ReviewMergeRequestResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReviewMergeRequestResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReviewMergeRequestResponse) GetBody() *ReviewMergeRequestResponseBody {
	return s.Body
}

func (s *ReviewMergeRequestResponse) SetHeaders(v map[string]*string) *ReviewMergeRequestResponse {
	s.Headers = v
	return s
}

func (s *ReviewMergeRequestResponse) SetStatusCode(v int32) *ReviewMergeRequestResponse {
	s.StatusCode = &v
	return s
}

func (s *ReviewMergeRequestResponse) SetBody(v *ReviewMergeRequestResponseBody) *ReviewMergeRequestResponse {
	s.Body = v
	return s
}

func (s *ReviewMergeRequestResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
