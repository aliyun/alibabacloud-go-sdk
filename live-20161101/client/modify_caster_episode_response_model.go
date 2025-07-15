// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCasterEpisodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCasterEpisodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCasterEpisodeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCasterEpisodeResponseBody) *ModifyCasterEpisodeResponse
	GetBody() *ModifyCasterEpisodeResponseBody
}

type ModifyCasterEpisodeResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCasterEpisodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCasterEpisodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCasterEpisodeResponse) GoString() string {
	return s.String()
}

func (s *ModifyCasterEpisodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCasterEpisodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCasterEpisodeResponse) GetBody() *ModifyCasterEpisodeResponseBody {
	return s.Body
}

func (s *ModifyCasterEpisodeResponse) SetHeaders(v map[string]*string) *ModifyCasterEpisodeResponse {
	s.Headers = v
	return s
}

func (s *ModifyCasterEpisodeResponse) SetStatusCode(v int32) *ModifyCasterEpisodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCasterEpisodeResponse) SetBody(v *ModifyCasterEpisodeResponseBody) *ModifyCasterEpisodeResponse {
	s.Body = v
	return s
}

func (s *ModifyCasterEpisodeResponse) Validate() error {
	return dara.Validate(s)
}
