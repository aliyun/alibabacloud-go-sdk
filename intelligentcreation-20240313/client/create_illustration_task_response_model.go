// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIllustrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIllustrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIllustrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *IllustrationTaskResult) *CreateIllustrationTaskResponse
	GetBody() *IllustrationTaskResult
}

type CreateIllustrationTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationTaskResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIllustrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIllustrationTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateIllustrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIllustrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIllustrationTaskResponse) GetBody() *IllustrationTaskResult {
	return s.Body
}

func (s *CreateIllustrationTaskResponse) SetHeaders(v map[string]*string) *CreateIllustrationTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateIllustrationTaskResponse) SetStatusCode(v int32) *CreateIllustrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIllustrationTaskResponse) SetBody(v *IllustrationTaskResult) *CreateIllustrationTaskResponse {
	s.Body = v
	return s
}

func (s *CreateIllustrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
