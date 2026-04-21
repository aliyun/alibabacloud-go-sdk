// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeAssetFoldersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListYikeAssetFoldersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListYikeAssetFoldersResponse
	GetStatusCode() *int32
	SetBody(v *ListYikeAssetFoldersResponseBody) *ListYikeAssetFoldersResponse
	GetBody() *ListYikeAssetFoldersResponseBody
}

type ListYikeAssetFoldersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYikeAssetFoldersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYikeAssetFoldersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListYikeAssetFoldersResponse) GoString() string {
	return s.String()
}

func (s *ListYikeAssetFoldersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListYikeAssetFoldersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListYikeAssetFoldersResponse) GetBody() *ListYikeAssetFoldersResponseBody {
	return s.Body
}

func (s *ListYikeAssetFoldersResponse) SetHeaders(v map[string]*string) *ListYikeAssetFoldersResponse {
	s.Headers = v
	return s
}

func (s *ListYikeAssetFoldersResponse) SetStatusCode(v int32) *ListYikeAssetFoldersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYikeAssetFoldersResponse) SetBody(v *ListYikeAssetFoldersResponseBody) *ListYikeAssetFoldersResponse {
	s.Body = v
	return s
}

func (s *ListYikeAssetFoldersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
