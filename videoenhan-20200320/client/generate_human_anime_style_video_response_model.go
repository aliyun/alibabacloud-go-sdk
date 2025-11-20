// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleVideoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleVideoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateHumanAnimeStyleVideoResponse
	GetStatusCode() *int32
	SetBody(v *GenerateHumanAnimeStyleVideoResponseBody) *GenerateHumanAnimeStyleVideoResponse
	GetBody() *GenerateHumanAnimeStyleVideoResponseBody
}

type GenerateHumanAnimeStyleVideoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateHumanAnimeStyleVideoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateHumanAnimeStyleVideoResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleVideoResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleVideoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateHumanAnimeStyleVideoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateHumanAnimeStyleVideoResponse) GetBody() *GenerateHumanAnimeStyleVideoResponseBody {
	return s.Body
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleVideoResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetStatusCode(v int32) *GenerateHumanAnimeStyleVideoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponse) SetBody(v *GenerateHumanAnimeStyleVideoResponseBody) *GenerateHumanAnimeStyleVideoResponse {
	s.Body = v
	return s
}

func (s *GenerateHumanAnimeStyleVideoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
