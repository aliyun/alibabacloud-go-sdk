// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCurrentPlayingListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCurrentPlayingListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCurrentPlayingListResponse
	GetStatusCode() *int32
	SetBody(v *GetCurrentPlayingListResponseBody) *GetCurrentPlayingListResponse
	GetBody() *GetCurrentPlayingListResponseBody
}

type GetCurrentPlayingListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCurrentPlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCurrentPlayingListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCurrentPlayingListResponse) GoString() string {
	return s.String()
}

func (s *GetCurrentPlayingListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCurrentPlayingListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCurrentPlayingListResponse) GetBody() *GetCurrentPlayingListResponseBody {
	return s.Body
}

func (s *GetCurrentPlayingListResponse) SetHeaders(v map[string]*string) *GetCurrentPlayingListResponse {
	s.Headers = v
	return s
}

func (s *GetCurrentPlayingListResponse) SetStatusCode(v int32) *GetCurrentPlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCurrentPlayingListResponse) SetBody(v *GetCurrentPlayingListResponseBody) *GetCurrentPlayingListResponse {
	s.Body = v
	return s
}

func (s *GetCurrentPlayingListResponse) Validate() error {
	return dara.Validate(s)
}
