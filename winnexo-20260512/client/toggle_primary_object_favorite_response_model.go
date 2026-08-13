// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTogglePrimaryObjectFavoriteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TogglePrimaryObjectFavoriteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TogglePrimaryObjectFavoriteResponse
	GetStatusCode() *int32
	SetBody(v *TogglePrimaryObjectFavoriteResponseBody) *TogglePrimaryObjectFavoriteResponse
	GetBody() *TogglePrimaryObjectFavoriteResponseBody
}

type TogglePrimaryObjectFavoriteResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TogglePrimaryObjectFavoriteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TogglePrimaryObjectFavoriteResponse) String() string {
	return dara.Prettify(s)
}

func (s TogglePrimaryObjectFavoriteResponse) GoString() string {
	return s.String()
}

func (s *TogglePrimaryObjectFavoriteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TogglePrimaryObjectFavoriteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TogglePrimaryObjectFavoriteResponse) GetBody() *TogglePrimaryObjectFavoriteResponseBody {
	return s.Body
}

func (s *TogglePrimaryObjectFavoriteResponse) SetHeaders(v map[string]*string) *TogglePrimaryObjectFavoriteResponse {
	s.Headers = v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponse) SetStatusCode(v int32) *TogglePrimaryObjectFavoriteResponse {
	s.StatusCode = &v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponse) SetBody(v *TogglePrimaryObjectFavoriteResponseBody) *TogglePrimaryObjectFavoriteResponse {
	s.Body = v
	return s
}

func (s *TogglePrimaryObjectFavoriteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
