// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTrackListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTrackListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTrackListResponse
	GetStatusCode() *int32
	SetBody(v *GetTrackListResponseBody) *GetTrackListResponse
	GetBody() *GetTrackListResponseBody
}

type GetTrackListResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTrackListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTrackListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTrackListResponse) GoString() string {
	return s.String()
}

func (s *GetTrackListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTrackListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTrackListResponse) GetBody() *GetTrackListResponseBody {
	return s.Body
}

func (s *GetTrackListResponse) SetHeaders(v map[string]*string) *GetTrackListResponse {
	s.Headers = v
	return s
}

func (s *GetTrackListResponse) SetStatusCode(v int32) *GetTrackListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTrackListResponse) SetBody(v *GetTrackListResponseBody) *GetTrackListResponse {
	s.Body = v
	return s
}

func (s *GetTrackListResponse) Validate() error {
	return dara.Validate(s)
}
