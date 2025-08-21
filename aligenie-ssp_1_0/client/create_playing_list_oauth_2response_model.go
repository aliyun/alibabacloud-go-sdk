// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListOAuth2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePlayingListOAuth2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePlayingListOAuth2Response
	GetStatusCode() *int32
	SetBody(v *CreatePlayingListOAuth2ResponseBody) *CreatePlayingListOAuth2Response
	GetBody() *CreatePlayingListOAuth2ResponseBody
}

type CreatePlayingListOAuth2Response struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlayingListOAuth2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePlayingListOAuth2Response) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListOAuth2Response) GoString() string {
	return s.String()
}

func (s *CreatePlayingListOAuth2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePlayingListOAuth2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePlayingListOAuth2Response) GetBody() *CreatePlayingListOAuth2ResponseBody {
	return s.Body
}

func (s *CreatePlayingListOAuth2Response) SetHeaders(v map[string]*string) *CreatePlayingListOAuth2Response {
	s.Headers = v
	return s
}

func (s *CreatePlayingListOAuth2Response) SetStatusCode(v int32) *CreatePlayingListOAuth2Response {
	s.StatusCode = &v
	return s
}

func (s *CreatePlayingListOAuth2Response) SetBody(v *CreatePlayingListOAuth2ResponseBody) *CreatePlayingListOAuth2Response {
	s.Body = v
	return s
}

func (s *CreatePlayingListOAuth2Response) Validate() error {
	return dara.Validate(s)
}
