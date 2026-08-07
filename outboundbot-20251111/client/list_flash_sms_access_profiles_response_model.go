// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsAccessProfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlashSmsAccessProfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlashSmsAccessProfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListFlashSmsAccessProfilesResponseBody) *ListFlashSmsAccessProfilesResponse
	GetBody() *ListFlashSmsAccessProfilesResponseBody
}

type ListFlashSmsAccessProfilesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlashSmsAccessProfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlashSmsAccessProfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsAccessProfilesResponse) GoString() string {
	return s.String()
}

func (s *ListFlashSmsAccessProfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlashSmsAccessProfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlashSmsAccessProfilesResponse) GetBody() *ListFlashSmsAccessProfilesResponseBody {
	return s.Body
}

func (s *ListFlashSmsAccessProfilesResponse) SetHeaders(v map[string]*string) *ListFlashSmsAccessProfilesResponse {
	s.Headers = v
	return s
}

func (s *ListFlashSmsAccessProfilesResponse) SetStatusCode(v int32) *ListFlashSmsAccessProfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlashSmsAccessProfilesResponse) SetBody(v *ListFlashSmsAccessProfilesResponseBody) *ListFlashSmsAccessProfilesResponse {
	s.Body = v
	return s
}

func (s *ListFlashSmsAccessProfilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
