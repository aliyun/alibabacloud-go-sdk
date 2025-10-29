// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterEpisodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterEpisodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterEpisodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterEpisodeGroupResponseBody) *DeleteCasterEpisodeGroupResponse
	GetBody() *DeleteCasterEpisodeGroupResponseBody
}

type DeleteCasterEpisodeGroupResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterEpisodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterEpisodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterEpisodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterEpisodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterEpisodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterEpisodeGroupResponse) GetBody() *DeleteCasterEpisodeGroupResponseBody {
	return s.Body
}

func (s *DeleteCasterEpisodeGroupResponse) SetHeaders(v map[string]*string) *DeleteCasterEpisodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterEpisodeGroupResponse) SetStatusCode(v int32) *DeleteCasterEpisodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterEpisodeGroupResponse) SetBody(v *DeleteCasterEpisodeGroupResponseBody) *DeleteCasterEpisodeGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterEpisodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
