// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryZnodeDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryZnodeDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryZnodeDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryZnodeDetailResponseBody) *QueryZnodeDetailResponse
	GetBody() *QueryZnodeDetailResponseBody
}

type QueryZnodeDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryZnodeDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryZnodeDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryZnodeDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryZnodeDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryZnodeDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryZnodeDetailResponse) GetBody() *QueryZnodeDetailResponseBody {
	return s.Body
}

func (s *QueryZnodeDetailResponse) SetHeaders(v map[string]*string) *QueryZnodeDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryZnodeDetailResponse) SetStatusCode(v int32) *QueryZnodeDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryZnodeDetailResponse) SetBody(v *QueryZnodeDetailResponseBody) *QueryZnodeDetailResponse {
	s.Body = v
	return s
}

func (s *QueryZnodeDetailResponse) Validate() error {
	return dara.Validate(s)
}
