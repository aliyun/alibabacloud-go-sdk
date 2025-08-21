// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCurrentPlayingItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCurrentPlayingItemResponse
	GetStatusCode() *int32
	SetBody(v *GetCurrentPlayingItemResponseBody) *GetCurrentPlayingItemResponse
	GetBody() *GetCurrentPlayingItemResponseBody
}

type GetCurrentPlayingItemResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCurrentPlayingItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCurrentPlayingItemResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingItemResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCurrentPlayingItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCurrentPlayingItemResponse) GetBody() *GetCurrentPlayingItemResponseBody {
	return s.Body
}

func (s *GetCurrentPlayingItemResponse) SetHeaders(v map[string]*string) *GetCurrentPlayingItemResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentPlayingItemResponse) SetStatusCode(v int32) *GetCurrentPlayingItemResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentPlayingItemResponse) SetBody(v *GetCurrentPlayingItemResponseBody) *GetCurrentPlayingItemResponse {
	s.Body = v
	return s
}

func (s *GetCurrentPlayingItemResponse) Validate() error {
	return dara.Validate(s)
}
