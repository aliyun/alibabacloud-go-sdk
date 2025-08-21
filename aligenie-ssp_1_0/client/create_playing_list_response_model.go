// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePlayingListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePlayingListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePlayingListResponse
	GetStatusCode() *int32
	SetBody(v *CreatePlayingListResponseBody) *CreatePlayingListResponse
	GetBody() *CreatePlayingListResponseBody
}

type CreatePlayingListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePlayingListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePlayingListResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePlayingListResponse) GoString() string {
	return s.String()
}

func (s *CreatePlayingListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePlayingListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePlayingListResponse) GetBody() *CreatePlayingListResponseBody {
	return s.Body
}

func (s *CreatePlayingListResponse) SetHeaders(v map[string]*string) *CreatePlayingListResponse {
	s.Headers = v
	return s
}

func (s *CreatePlayingListResponse) SetStatusCode(v int32) *CreatePlayingListResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePlayingListResponse) SetBody(v *CreatePlayingListResponseBody) *CreatePlayingListResponse {
	s.Body = v
	return s
}

func (s *CreatePlayingListResponse) Validate() error {
	return dara.Validate(s)
}
