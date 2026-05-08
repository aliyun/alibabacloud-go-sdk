// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIllustrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIllustrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIllustrationResponse
	GetStatusCode() *int32
	SetBody(v *IllustrationResult) *GetIllustrationResponse
	GetBody() *IllustrationResult
}

type GetIllustrationResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IllustrationResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIllustrationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIllustrationResponse) GoString() string {
	return s.String()
}

func (s *GetIllustrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIllustrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIllustrationResponse) GetBody() *IllustrationResult {
	return s.Body
}

func (s *GetIllustrationResponse) SetHeaders(v map[string]*string) *GetIllustrationResponse {
	s.Headers = v
	return s
}

func (s *GetIllustrationResponse) SetStatusCode(v int32) *GetIllustrationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIllustrationResponse) SetBody(v *IllustrationResult) *GetIllustrationResponse {
	s.Body = v
	return s
}

func (s *GetIllustrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
