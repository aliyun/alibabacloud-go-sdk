// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAndRemoveFavoriteContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAndRemoveFavoriteContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAndRemoveFavoriteContentResponse
	GetStatusCode() *int32
	SetBody(v *AddAndRemoveFavoriteContentResponseBody) *AddAndRemoveFavoriteContentResponse
	GetBody() *AddAndRemoveFavoriteContentResponseBody
}

type AddAndRemoveFavoriteContentResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAndRemoveFavoriteContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAndRemoveFavoriteContentResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAndRemoveFavoriteContentResponse) GoString() string {
	return s.String()
}

func (s *AddAndRemoveFavoriteContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAndRemoveFavoriteContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAndRemoveFavoriteContentResponse) GetBody() *AddAndRemoveFavoriteContentResponseBody {
	return s.Body
}

func (s *AddAndRemoveFavoriteContentResponse) SetHeaders(v map[string]*string) *AddAndRemoveFavoriteContentResponse {
	s.Headers = v
	return s
}

func (s *AddAndRemoveFavoriteContentResponse) SetStatusCode(v int32) *AddAndRemoveFavoriteContentResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAndRemoveFavoriteContentResponse) SetBody(v *AddAndRemoveFavoriteContentResponseBody) *AddAndRemoveFavoriteContentResponse {
	s.Body = v
	return s
}

func (s *AddAndRemoveFavoriteContentResponse) Validate() error {
	return dara.Validate(s)
}
