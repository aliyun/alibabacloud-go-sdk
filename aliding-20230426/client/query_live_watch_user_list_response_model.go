// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLiveWatchUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLiveWatchUserListResponse
	GetStatusCode() *int32
	SetBody(v *QueryLiveWatchUserListResponseBody) *QueryLiveWatchUserListResponse
	GetBody() *QueryLiveWatchUserListResponseBody
}

type QueryLiveWatchUserListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveWatchUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveWatchUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchUserListResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLiveWatchUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLiveWatchUserListResponse) GetBody() *QueryLiveWatchUserListResponseBody {
	return s.Body
}

func (s *QueryLiveWatchUserListResponse) SetHeaders(v map[string]*string) *QueryLiveWatchUserListResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveWatchUserListResponse) SetStatusCode(v int32) *QueryLiveWatchUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveWatchUserListResponse) SetBody(v *QueryLiveWatchUserListResponseBody) *QueryLiveWatchUserListResponse {
	s.Body = v
	return s
}

func (s *QueryLiveWatchUserListResponse) Validate() error {
	return dara.Validate(s)
}
