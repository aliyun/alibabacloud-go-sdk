// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTextThemesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTextThemesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTextThemesResponse
	GetStatusCode() *int32
	SetBody(v *TextThemeListResult) *ListTextThemesResponse
	GetBody() *TextThemeListResult
}

type ListTextThemesResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextThemeListResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTextThemesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTextThemesResponse) GoString() string {
	return s.String()
}

func (s *ListTextThemesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTextThemesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTextThemesResponse) GetBody() *TextThemeListResult {
	return s.Body
}

func (s *ListTextThemesResponse) SetHeaders(v map[string]*string) *ListTextThemesResponse {
	s.Headers = v
	return s
}

func (s *ListTextThemesResponse) SetStatusCode(v int32) *ListTextThemesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTextThemesResponse) SetBody(v *TextThemeListResult) *ListTextThemesResponse {
	s.Body = v
	return s
}

func (s *ListTextThemesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
