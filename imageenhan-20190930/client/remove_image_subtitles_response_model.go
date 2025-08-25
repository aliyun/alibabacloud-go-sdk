// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageSubtitlesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveImageSubtitlesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveImageSubtitlesResponse
	GetStatusCode() *int32
	SetBody(v *RemoveImageSubtitlesResponseBody) *RemoveImageSubtitlesResponse
	GetBody() *RemoveImageSubtitlesResponseBody
}

type RemoveImageSubtitlesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveImageSubtitlesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveImageSubtitlesResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageSubtitlesResponse) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveImageSubtitlesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveImageSubtitlesResponse) GetBody() *RemoveImageSubtitlesResponseBody {
	return s.Body
}

func (s *RemoveImageSubtitlesResponse) SetHeaders(v map[string]*string) *RemoveImageSubtitlesResponse {
	s.Headers = v
	return s
}

func (s *RemoveImageSubtitlesResponse) SetStatusCode(v int32) *RemoveImageSubtitlesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveImageSubtitlesResponse) SetBody(v *RemoveImageSubtitlesResponseBody) *RemoveImageSubtitlesResponse {
	s.Body = v
	return s
}

func (s *RemoveImageSubtitlesResponse) Validate() error {
	return dara.Validate(s)
}
