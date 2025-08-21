// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexControlPlayingListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *IndexControlPlayingListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *IndexControlPlayingListResponse
	GetStatusCode() *int32
	SetBody(v *IndexControlPlayingListResponseBody) *IndexControlPlayingListResponse
	GetBody() *IndexControlPlayingListResponseBody
}

type IndexControlPlayingListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IndexControlPlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s IndexControlPlayingListResponse) String() string {
	return dara.Prettify(s)
}

func (s IndexControlPlayingListResponse) GoString() string {
	return s.String()
}

func (s *IndexControlPlayingListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *IndexControlPlayingListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *IndexControlPlayingListResponse) GetBody() *IndexControlPlayingListResponseBody {
	return s.Body
}

func (s *IndexControlPlayingListResponse) SetHeaders(v map[string]*string) *IndexControlPlayingListResponse {
	s.Headers = v
	return s
}

func (s *IndexControlPlayingListResponse) SetStatusCode(v int32) *IndexControlPlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *IndexControlPlayingListResponse) SetBody(v *IndexControlPlayingListResponseBody) *IndexControlPlayingListResponse {
	s.Body = v
	return s
}

func (s *IndexControlPlayingListResponse) Validate() error {
	return dara.Validate(s)
}
