// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemoveOperatingObjectFavoritesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRemoveOperatingObjectFavoritesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRemoveOperatingObjectFavoritesResponse
	GetStatusCode() *int32
	SetBody(v *BatchRemoveOperatingObjectFavoritesResponseBody) *BatchRemoveOperatingObjectFavoritesResponse
	GetBody() *BatchRemoveOperatingObjectFavoritesResponseBody
}

type BatchRemoveOperatingObjectFavoritesResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRemoveOperatingObjectFavoritesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRemoveOperatingObjectFavoritesResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRemoveOperatingObjectFavoritesResponse) GoString() string {
	return s.String()
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) GetBody() *BatchRemoveOperatingObjectFavoritesResponseBody {
	return s.Body
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) SetHeaders(v map[string]*string) *BatchRemoveOperatingObjectFavoritesResponse {
	s.Headers = v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) SetStatusCode(v int32) *BatchRemoveOperatingObjectFavoritesResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) SetBody(v *BatchRemoveOperatingObjectFavoritesResponseBody) *BatchRemoveOperatingObjectFavoritesResponse {
	s.Body = v
	return s
}

func (s *BatchRemoveOperatingObjectFavoritesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
