// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAvatarResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAvatarResourceResponse
	GetStatusCode() *int32
	SetBody(v *QueryAvatarResourceResponseBody) *QueryAvatarResourceResponse
	GetBody() *QueryAvatarResourceResponseBody
}

type QueryAvatarResourceResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAvatarResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAvatarResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarResourceResponse) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAvatarResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAvatarResourceResponse) GetBody() *QueryAvatarResourceResponseBody {
	return s.Body
}

func (s *QueryAvatarResourceResponse) SetHeaders(v map[string]*string) *QueryAvatarResourceResponse {
	s.Headers = v
	return s
}

func (s *QueryAvatarResourceResponse) SetStatusCode(v int32) *QueryAvatarResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAvatarResourceResponse) SetBody(v *QueryAvatarResourceResponseBody) *QueryAvatarResourceResponse {
	s.Body = v
	return s
}

func (s *QueryAvatarResourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
