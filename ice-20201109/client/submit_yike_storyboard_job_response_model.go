// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeStoryboardJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeStoryboardJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeStoryboardJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeStoryboardJobResponseBody) *SubmitYikeStoryboardJobResponse
	GetBody() *SubmitYikeStoryboardJobResponseBody
}

type SubmitYikeStoryboardJobResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeStoryboardJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeStoryboardJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeStoryboardJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeStoryboardJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeStoryboardJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeStoryboardJobResponse) GetBody() *SubmitYikeStoryboardJobResponseBody {
	return s.Body
}

func (s *SubmitYikeStoryboardJobResponse) SetHeaders(v map[string]*string) *SubmitYikeStoryboardJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeStoryboardJobResponse) SetStatusCode(v int32) *SubmitYikeStoryboardJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeStoryboardJobResponse) SetBody(v *SubmitYikeStoryboardJobResponseBody) *SubmitYikeStoryboardJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeStoryboardJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
