// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeGroupContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterEpisodeGroupContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterEpisodeGroupContentResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterEpisodeGroupContentResponseBody) *AddCasterEpisodeGroupContentResponse
	GetBody() *AddCasterEpisodeGroupContentResponseBody
}

type AddCasterEpisodeGroupContentResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterEpisodeGroupContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterEpisodeGroupContentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupContentResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterEpisodeGroupContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterEpisodeGroupContentResponse) GetBody() *AddCasterEpisodeGroupContentResponseBody {
	return s.Body
}

func (s *AddCasterEpisodeGroupContentResponse) SetHeaders(v map[string]*string) *AddCasterEpisodeGroupContentResponse {
	s.Headers = v
	return s
}

func (s *AddCasterEpisodeGroupContentResponse) SetStatusCode(v int32) *AddCasterEpisodeGroupContentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterEpisodeGroupContentResponse) SetBody(v *AddCasterEpisodeGroupContentResponseBody) *AddCasterEpisodeGroupContentResponse {
	s.Body = v
	return s
}

func (s *AddCasterEpisodeGroupContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
