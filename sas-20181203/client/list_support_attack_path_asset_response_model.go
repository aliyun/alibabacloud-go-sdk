// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportAttackPathAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSupportAttackPathAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSupportAttackPathAssetResponse
	GetStatusCode() *int32
	SetBody(v *ListSupportAttackPathAssetResponseBody) *ListSupportAttackPathAssetResponse
	GetBody() *ListSupportAttackPathAssetResponseBody
}

type ListSupportAttackPathAssetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSupportAttackPathAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSupportAttackPathAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSupportAttackPathAssetResponse) GoString() string {
	return s.String()
}

func (s *ListSupportAttackPathAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSupportAttackPathAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSupportAttackPathAssetResponse) GetBody() *ListSupportAttackPathAssetResponseBody {
	return s.Body
}

func (s *ListSupportAttackPathAssetResponse) SetHeaders(v map[string]*string) *ListSupportAttackPathAssetResponse {
	s.Headers = v
	return s
}

func (s *ListSupportAttackPathAssetResponse) SetStatusCode(v int32) *ListSupportAttackPathAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSupportAttackPathAssetResponse) SetBody(v *ListSupportAttackPathAssetResponseBody) *ListSupportAttackPathAssetResponse {
	s.Body = v
	return s
}

func (s *ListSupportAttackPathAssetResponse) Validate() error {
	return dara.Validate(s)
}
