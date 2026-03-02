// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetCountResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetCountResponseBody) *ListAssetCountResponse
	GetBody() *ListAssetCountResponseBody
}

type ListAssetCountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetCountResponse) GoString() string {
	return s.String()
}

func (s *ListAssetCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetCountResponse) GetBody() *ListAssetCountResponseBody {
	return s.Body
}

func (s *ListAssetCountResponse) SetHeaders(v map[string]*string) *ListAssetCountResponse {
	s.Headers = v
	return s
}

func (s *ListAssetCountResponse) SetStatusCode(v int32) *ListAssetCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetCountResponse) SetBody(v *ListAssetCountResponseBody) *ListAssetCountResponse {
	s.Body = v
	return s
}

func (s *ListAssetCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
