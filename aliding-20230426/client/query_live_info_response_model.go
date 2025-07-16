// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLiveInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryLiveInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryLiveInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryLiveInfoResponseBody) *QueryLiveInfoResponse
	GetBody() *QueryLiveInfoResponseBody
}

type QueryLiveInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryLiveInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryLiveInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryLiveInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryLiveInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryLiveInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryLiveInfoResponse) GetBody() *QueryLiveInfoResponseBody {
	return s.Body
}

func (s *QueryLiveInfoResponse) SetHeaders(v map[string]*string) *QueryLiveInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryLiveInfoResponse) SetStatusCode(v int32) *QueryLiveInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryLiveInfoResponse) SetBody(v *QueryLiveInfoResponseBody) *QueryLiveInfoResponse {
	s.Body = v
	return s
}

func (s *QueryLiveInfoResponse) Validate() error {
	return dara.Validate(s)
}
