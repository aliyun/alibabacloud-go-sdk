// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAvatarVideoJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAvatarVideoJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAvatarVideoJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAvatarVideoJobResponseBody) *SubmitAvatarVideoJobResponse
	GetBody() *SubmitAvatarVideoJobResponseBody
}

type SubmitAvatarVideoJobResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAvatarVideoJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAvatarVideoJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAvatarVideoJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAvatarVideoJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAvatarVideoJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAvatarVideoJobResponse) GetBody() *SubmitAvatarVideoJobResponseBody {
	return s.Body
}

func (s *SubmitAvatarVideoJobResponse) SetHeaders(v map[string]*string) *SubmitAvatarVideoJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAvatarVideoJobResponse) SetStatusCode(v int32) *SubmitAvatarVideoJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAvatarVideoJobResponse) SetBody(v *SubmitAvatarVideoJobResponseBody) *SubmitAvatarVideoJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAvatarVideoJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
