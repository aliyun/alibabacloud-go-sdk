// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveWatchDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLiveWatchDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLiveWatchDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryLiveWatchDetailResponseBody) *QueryLiveWatchDetailResponse
	GetBody() *QueryLiveWatchDetailResponseBody
}

type QueryLiveWatchDetailResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveWatchDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveWatchDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveWatchDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveWatchDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLiveWatchDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLiveWatchDetailResponse) GetBody() *QueryLiveWatchDetailResponseBody {
	return s.Body
}

func (s *QueryLiveWatchDetailResponse) SetHeaders(v map[string]*string) *QueryLiveWatchDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveWatchDetailResponse) SetStatusCode(v int32) *QueryLiveWatchDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveWatchDetailResponse) SetBody(v *QueryLiveWatchDetailResponseBody) *QueryLiveWatchDetailResponse {
	s.Body = v
	return s
}

func (s *QueryLiveWatchDetailResponse) Validate() error {
	return dara.Validate(s)
}
