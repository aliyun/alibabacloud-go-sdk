// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLlmCubeWithThemeListByUserIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLlmCubeWithThemeListByUserIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLlmCubeWithThemeListByUserIdResponse
	GetStatusCode() *int32
	SetBody(v *QueryLlmCubeWithThemeListByUserIdResponseBody) *QueryLlmCubeWithThemeListByUserIdResponse
	GetBody() *QueryLlmCubeWithThemeListByUserIdResponseBody
}

type QueryLlmCubeWithThemeListByUserIdResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLlmCubeWithThemeListByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLlmCubeWithThemeListByUserIdResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLlmCubeWithThemeListByUserIdResponse) GoString() string {
	return s.String()
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) GetBody() *QueryLlmCubeWithThemeListByUserIdResponseBody {
	return s.Body
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) SetHeaders(v map[string]*string) *QueryLlmCubeWithThemeListByUserIdResponse {
	s.Headers = v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) SetStatusCode(v int32) *QueryLlmCubeWithThemeListByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) SetBody(v *QueryLlmCubeWithThemeListByUserIdResponseBody) *QueryLlmCubeWithThemeListByUserIdResponse {
	s.Body = v
	return s
}

func (s *QueryLlmCubeWithThemeListByUserIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
