// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateHumanAnimeStyleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateHumanAnimeStyleResponse
	GetStatusCode() *int32
	SetBody(v *GenerateHumanAnimeStyleResponseBody) *GenerateHumanAnimeStyleResponse
	GetBody() *GenerateHumanAnimeStyleResponseBody
}

type GenerateHumanAnimeStyleResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateHumanAnimeStyleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateHumanAnimeStyleResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateHumanAnimeStyleResponse) GoString() string {
	return s.String()
}

func (s *GenerateHumanAnimeStyleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateHumanAnimeStyleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateHumanAnimeStyleResponse) GetBody() *GenerateHumanAnimeStyleResponseBody {
	return s.Body
}

func (s *GenerateHumanAnimeStyleResponse) SetHeaders(v map[string]*string) *GenerateHumanAnimeStyleResponse {
	s.Headers = v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) SetStatusCode(v int32) *GenerateHumanAnimeStyleResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) SetBody(v *GenerateHumanAnimeStyleResponseBody) *GenerateHumanAnimeStyleResponse {
	s.Body = v
	return s
}

func (s *GenerateHumanAnimeStyleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
