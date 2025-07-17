// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataAssetsResponse
	GetStatusCode() *int32
	SetBody(v *ListDataAssetsResponseBody) *ListDataAssetsResponse
	GetBody() *ListDataAssetsResponseBody
}

type ListDataAssetsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataAssetsResponse) GoString() string {
	return s.String()
}

func (s *ListDataAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataAssetsResponse) GetBody() *ListDataAssetsResponseBody {
	return s.Body
}

func (s *ListDataAssetsResponse) SetHeaders(v map[string]*string) *ListDataAssetsResponse {
	s.Headers = v
	return s
}

func (s *ListDataAssetsResponse) SetStatusCode(v int32) *ListDataAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataAssetsResponse) SetBody(v *ListDataAssetsResponseBody) *ListDataAssetsResponse {
	s.Body = v
	return s
}

func (s *ListDataAssetsResponse) Validate() error {
	return dara.Validate(s)
}
