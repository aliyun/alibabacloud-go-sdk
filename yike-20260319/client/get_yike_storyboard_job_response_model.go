// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeStoryboardJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetYikeStoryboardJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetYikeStoryboardJobResponse
	GetStatusCode() *int32
	SetBody(v *GetYikeStoryboardJobResponseBody) *GetYikeStoryboardJobResponse
	GetBody() *GetYikeStoryboardJobResponseBody
}

type GetYikeStoryboardJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetYikeStoryboardJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetYikeStoryboardJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetYikeStoryboardJobResponse) GoString() string {
	return s.String()
}

func (s *GetYikeStoryboardJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetYikeStoryboardJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetYikeStoryboardJobResponse) GetBody() *GetYikeStoryboardJobResponseBody {
	return s.Body
}

func (s *GetYikeStoryboardJobResponse) SetHeaders(v map[string]*string) *GetYikeStoryboardJobResponse {
	s.Headers = v
	return s
}

func (s *GetYikeStoryboardJobResponse) SetStatusCode(v int32) *GetYikeStoryboardJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetYikeStoryboardJobResponse) SetBody(v *GetYikeStoryboardJobResponseBody) *GetYikeStoryboardJobResponse {
	s.Body = v
	return s
}

func (s *GetYikeStoryboardJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
