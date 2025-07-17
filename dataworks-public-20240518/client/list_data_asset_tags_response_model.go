// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAssetTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAssetTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAssetTagsResponseBody) *ListDataAssetTagsResponse
	GetBody() *ListDataAssetTagsResponseBody
}

type ListDataAssetTagsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAssetTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAssetTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetTagsResponse) GoString() string {
	return s.String()
}

func (s *ListDataAssetTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAssetTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAssetTagsResponse) GetBody() *ListDataAssetTagsResponseBody {
	return s.Body
}

func (s *ListDataAssetTagsResponse) SetHeaders(v map[string]*string) *ListDataAssetTagsResponse {
	s.Headers = v
	return s
}

func (s *ListDataAssetTagsResponse) SetStatusCode(v int32) *ListDataAssetTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAssetTagsResponse) SetBody(v *ListDataAssetTagsResponseBody) *ListDataAssetTagsResponse {
	s.Body = v
	return s
}

func (s *ListDataAssetTagsResponse) Validate() error {
	return dara.Validate(s)
}
