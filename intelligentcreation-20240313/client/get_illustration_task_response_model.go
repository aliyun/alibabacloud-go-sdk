// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIllustrationTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIllustrationTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIllustrationTaskResponse
	GetStatusCode() *int32
	SetBody(v *IllustrationTaskResult) *GetIllustrationTaskResponse
	GetBody() *IllustrationTaskResult
}

type GetIllustrationTaskResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationTaskResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIllustrationTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIllustrationTaskResponse) GoString() string {
	return s.String()
}

func (s *GetIllustrationTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIllustrationTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIllustrationTaskResponse) GetBody() *IllustrationTaskResult {
	return s.Body
}

func (s *GetIllustrationTaskResponse) SetHeaders(v map[string]*string) *GetIllustrationTaskResponse {
	s.Headers = v
	return s
}

func (s *GetIllustrationTaskResponse) SetStatusCode(v int32) *GetIllustrationTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIllustrationTaskResponse) SetBody(v *IllustrationTaskResult) *GetIllustrationTaskResponse {
	s.Body = v
	return s
}

func (s *GetIllustrationTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
