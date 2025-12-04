// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGeneralConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGeneralConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGeneralConfigsResponse
	GetStatusCode() *int32
	SetBody(v *ListGeneralConfigsResponseBody) *ListGeneralConfigsResponse
	GetBody() *ListGeneralConfigsResponseBody
}

type ListGeneralConfigsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGeneralConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGeneralConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGeneralConfigsResponse) GoString() string {
	return s.String()
}

func (s *ListGeneralConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGeneralConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGeneralConfigsResponse) GetBody() *ListGeneralConfigsResponseBody {
	return s.Body
}

func (s *ListGeneralConfigsResponse) SetHeaders(v map[string]*string) *ListGeneralConfigsResponse {
	s.Headers = v
	return s
}

func (s *ListGeneralConfigsResponse) SetStatusCode(v int32) *ListGeneralConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGeneralConfigsResponse) SetBody(v *ListGeneralConfigsResponseBody) *ListGeneralConfigsResponse {
	s.Body = v
	return s
}

func (s *ListGeneralConfigsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
