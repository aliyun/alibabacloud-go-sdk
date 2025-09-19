// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAssetInfoPublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAssetInfoPublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAssetInfoPublishResponse
	GetStatusCode() *int32
	SetBody(v *ListAssetInfoPublishResponseBody) *ListAssetInfoPublishResponse
	GetBody() *ListAssetInfoPublishResponseBody
}

type ListAssetInfoPublishResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssetInfoPublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssetInfoPublishResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAssetInfoPublishResponse) GoString() string {
	return s.String()
}

func (s *ListAssetInfoPublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAssetInfoPublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAssetInfoPublishResponse) GetBody() *ListAssetInfoPublishResponseBody {
	return s.Body
}

func (s *ListAssetInfoPublishResponse) SetHeaders(v map[string]*string) *ListAssetInfoPublishResponse {
	s.Headers = v
	return s
}

func (s *ListAssetInfoPublishResponse) SetStatusCode(v int32) *ListAssetInfoPublishResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssetInfoPublishResponse) SetBody(v *ListAssetInfoPublishResponseBody) *ListAssetInfoPublishResponse {
	s.Body = v
	return s
}

func (s *ListAssetInfoPublishResponse) Validate() error {
	return dara.Validate(s)
}
