// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserTagValueListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserTagValueListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserTagValueListResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserTagValueListResponseBody) *QueryUserTagValueListResponse
	GetBody() *QueryUserTagValueListResponseBody
}

type QueryUserTagValueListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserTagValueListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserTagValueListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagValueListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserTagValueListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserTagValueListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserTagValueListResponse) GetBody() *QueryUserTagValueListResponseBody {
	return s.Body
}

func (s *QueryUserTagValueListResponse) SetHeaders(v map[string]*string) *QueryUserTagValueListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserTagValueListResponse) SetStatusCode(v int32) *QueryUserTagValueListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserTagValueListResponse) SetBody(v *QueryUserTagValueListResponseBody) *QueryUserTagValueListResponse {
	s.Body = v
	return s
}

func (s *QueryUserTagValueListResponse) Validate() error {
	return dara.Validate(s)
}
