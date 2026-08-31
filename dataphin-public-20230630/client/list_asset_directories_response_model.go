// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetDirectoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetDirectoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetDirectoriesResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetDirectoriesResponseBody) *ListAssetDirectoriesResponse
	GetBody() *ListAssetDirectoriesResponseBody
}

type ListAssetDirectoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetDirectoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetDirectoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetDirectoriesResponse) GoString() string {
	return s.String()
}

func (s *ListAssetDirectoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetDirectoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetDirectoriesResponse) GetBody() *ListAssetDirectoriesResponseBody {
	return s.Body
}

func (s *ListAssetDirectoriesResponse) SetHeaders(v map[string]*string) *ListAssetDirectoriesResponse {
	s.Headers = v
	return s
}

func (s *ListAssetDirectoriesResponse) SetStatusCode(v int32) *ListAssetDirectoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetDirectoriesResponse) SetBody(v *ListAssetDirectoriesResponseBody) *ListAssetDirectoriesResponse {
	s.Body = v
	return s
}

func (s *ListAssetDirectoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
