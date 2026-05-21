// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeAvatarNarratorJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitYikeAvatarNarratorJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitYikeAvatarNarratorJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitYikeAvatarNarratorJobResponseBody) *SubmitYikeAvatarNarratorJobResponse
	GetBody() *SubmitYikeAvatarNarratorJobResponseBody
}

type SubmitYikeAvatarNarratorJobResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitYikeAvatarNarratorJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitYikeAvatarNarratorJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeAvatarNarratorJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitYikeAvatarNarratorJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitYikeAvatarNarratorJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitYikeAvatarNarratorJobResponse) GetBody() *SubmitYikeAvatarNarratorJobResponseBody {
	return s.Body
}

func (s *SubmitYikeAvatarNarratorJobResponse) SetHeaders(v map[string]*string) *SubmitYikeAvatarNarratorJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitYikeAvatarNarratorJobResponse) SetStatusCode(v int32) *SubmitYikeAvatarNarratorJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitYikeAvatarNarratorJobResponse) SetBody(v *SubmitYikeAvatarNarratorJobResponseBody) *SubmitYikeAvatarNarratorJobResponse {
	s.Body = v
	return s
}

func (s *SubmitYikeAvatarNarratorJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
