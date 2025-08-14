// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFavoritePublicMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelFavoritePublicMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelFavoritePublicMediaResponse
	GetStatusCode() *int32
	SetBody(v *CancelFavoritePublicMediaResponseBody) *CancelFavoritePublicMediaResponse
	GetBody() *CancelFavoritePublicMediaResponseBody
}

type CancelFavoritePublicMediaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelFavoritePublicMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelFavoritePublicMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelFavoritePublicMediaResponse) GoString() string {
	return s.String()
}

func (s *CancelFavoritePublicMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelFavoritePublicMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelFavoritePublicMediaResponse) GetBody() *CancelFavoritePublicMediaResponseBody {
	return s.Body
}

func (s *CancelFavoritePublicMediaResponse) SetHeaders(v map[string]*string) *CancelFavoritePublicMediaResponse {
	s.Headers = v
	return s
}

func (s *CancelFavoritePublicMediaResponse) SetStatusCode(v int32) *CancelFavoritePublicMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelFavoritePublicMediaResponse) SetBody(v *CancelFavoritePublicMediaResponseBody) *CancelFavoritePublicMediaResponse {
	s.Body = v
	return s
}

func (s *CancelFavoritePublicMediaResponse) Validate() error {
	return dara.Validate(s)
}
