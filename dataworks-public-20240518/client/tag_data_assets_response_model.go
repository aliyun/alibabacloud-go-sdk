// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTagDataAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagDataAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagDataAssetsResponse
	GetStatusCode() *int32
	SetBody(v *TagDataAssetsResponseBody) *TagDataAssetsResponse
	GetBody() *TagDataAssetsResponseBody
}

type TagDataAssetsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagDataAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s TagDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *TagDataAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagDataAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagDataAssetsResponse) GetBody() *TagDataAssetsResponseBody {
	return s.Body
}

func (s *TagDataAssetsResponse) SetHeaders(v map[string]*string) *TagDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *TagDataAssetsResponse) SetStatusCode(v int32) *TagDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *TagDataAssetsResponse) SetBody(v *TagDataAssetsResponseBody) *TagDataAssetsResponse {
	s.Body = v
	return s
}

func (s *TagDataAssetsResponse) Validate() error {
	return dara.Validate(s)
}
