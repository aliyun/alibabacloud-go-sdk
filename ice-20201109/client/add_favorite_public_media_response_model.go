// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFavoritePublicMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFavoritePublicMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFavoritePublicMediaResponse
	GetStatusCode() *int32
	SetBody(v *AddFavoritePublicMediaResponseBody) *AddFavoritePublicMediaResponse
	GetBody() *AddFavoritePublicMediaResponseBody
}

type AddFavoritePublicMediaResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFavoritePublicMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFavoritePublicMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFavoritePublicMediaResponse) GoString() string {
	return s.String()
}

func (s *AddFavoritePublicMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFavoritePublicMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFavoritePublicMediaResponse) GetBody() *AddFavoritePublicMediaResponseBody {
	return s.Body
}

func (s *AddFavoritePublicMediaResponse) SetHeaders(v map[string]*string) *AddFavoritePublicMediaResponse {
	s.Headers = v
	return s
}

func (s *AddFavoritePublicMediaResponse) SetStatusCode(v int32) *AddFavoritePublicMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFavoritePublicMediaResponse) SetBody(v *AddFavoritePublicMediaResponseBody) *AddFavoritePublicMediaResponse {
	s.Body = v
	return s
}

func (s *AddFavoritePublicMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
