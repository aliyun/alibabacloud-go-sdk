// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoV4Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitReviewInfoV4Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitReviewInfoV4Response
	GetStatusCode() *int32
	SetBody(v *SubmitReviewInfoV4ResponseBody) *SubmitReviewInfoV4Response
	GetBody() *SubmitReviewInfoV4ResponseBody
}

type SubmitReviewInfoV4Response struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitReviewInfoV4ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitReviewInfoV4Response) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoV4Response) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoV4Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitReviewInfoV4Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitReviewInfoV4Response) GetBody() *SubmitReviewInfoV4ResponseBody {
	return s.Body
}

func (s *SubmitReviewInfoV4Response) SetHeaders(v map[string]*string) *SubmitReviewInfoV4Response {
	s.Headers = v
	return s
}

func (s *SubmitReviewInfoV4Response) SetStatusCode(v int32) *SubmitReviewInfoV4Response {
	s.StatusCode = &v
	return s
}

func (s *SubmitReviewInfoV4Response) SetBody(v *SubmitReviewInfoV4ResponseBody) *SubmitReviewInfoV4Response {
	s.Body = v
	return s
}

func (s *SubmitReviewInfoV4Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
