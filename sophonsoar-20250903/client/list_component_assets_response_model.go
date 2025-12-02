// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComponentAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListComponentAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListComponentAssetsResponse
	GetStatusCode() *int32
	SetBody(v *ListComponentAssetsResponseBody) *ListComponentAssetsResponse
	GetBody() *ListComponentAssetsResponseBody
}

type ListComponentAssetsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListComponentAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListComponentAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListComponentAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListComponentAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListComponentAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListComponentAssetsResponse) GetBody() *ListComponentAssetsResponseBody {
	return s.Body
}

func (s *ListComponentAssetsResponse) SetHeaders(v map[string]*string) *ListComponentAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListComponentAssetsResponse) SetStatusCode(v int32) *ListComponentAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListComponentAssetsResponse) SetBody(v *ListComponentAssetsResponseBody) *ListComponentAssetsResponse {
	s.Body = v
	return s
}

func (s *ListComponentAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
