// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMediaCensorJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryMediaCensorJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryMediaCensorJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryMediaCensorJobDetailResponseBody) *QueryMediaCensorJobDetailResponse
	GetBody() *QueryMediaCensorJobDetailResponseBody
}

type QueryMediaCensorJobDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryMediaCensorJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryMediaCensorJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryMediaCensorJobDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryMediaCensorJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryMediaCensorJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryMediaCensorJobDetailResponse) GetBody() *QueryMediaCensorJobDetailResponseBody {
	return s.Body
}

func (s *QueryMediaCensorJobDetailResponse) SetHeaders(v map[string]*string) *QueryMediaCensorJobDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryMediaCensorJobDetailResponse) SetStatusCode(v int32) *QueryMediaCensorJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryMediaCensorJobDetailResponse) SetBody(v *QueryMediaCensorJobDetailResponseBody) *QueryMediaCensorJobDetailResponse {
	s.Body = v
	return s
}

func (s *QueryMediaCensorJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
