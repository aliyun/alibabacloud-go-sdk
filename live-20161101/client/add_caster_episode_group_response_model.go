// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterEpisodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterEpisodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterEpisodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterEpisodeGroupResponseBody) *AddCasterEpisodeGroupResponse
	GetBody() *AddCasterEpisodeGroupResponseBody
}

type AddCasterEpisodeGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterEpisodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterEpisodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterEpisodeGroupResponse) GoString() string {
	return s.String()
}

func (s *AddCasterEpisodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterEpisodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterEpisodeGroupResponse) GetBody() *AddCasterEpisodeGroupResponseBody {
	return s.Body
}

func (s *AddCasterEpisodeGroupResponse) SetHeaders(v map[string]*string) *AddCasterEpisodeGroupResponse {
	s.Headers = v
	return s
}

func (s *AddCasterEpisodeGroupResponse) SetStatusCode(v int32) *AddCasterEpisodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterEpisodeGroupResponse) SetBody(v *AddCasterEpisodeGroupResponseBody) *AddCasterEpisodeGroupResponse {
	s.Body = v
	return s
}

func (s *AddCasterEpisodeGroupResponse) Validate() error {
	return dara.Validate(s)
}
