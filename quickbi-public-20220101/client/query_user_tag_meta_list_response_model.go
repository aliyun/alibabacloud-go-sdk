// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserTagMetaListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryUserTagMetaListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryUserTagMetaListResponse
	GetStatusCode() *int32
	SetBody(v *QueryUserTagMetaListResponseBody) *QueryUserTagMetaListResponse
	GetBody() *QueryUserTagMetaListResponseBody
}

type QueryUserTagMetaListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryUserTagMetaListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryUserTagMetaListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryUserTagMetaListResponse) GoString() string {
	return s.String()
}

func (s *QueryUserTagMetaListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryUserTagMetaListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryUserTagMetaListResponse) GetBody() *QueryUserTagMetaListResponseBody {
	return s.Body
}

func (s *QueryUserTagMetaListResponse) SetHeaders(v map[string]*string) *QueryUserTagMetaListResponse {
	s.Headers = v
	return s
}

func (s *QueryUserTagMetaListResponse) SetStatusCode(v int32) *QueryUserTagMetaListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryUserTagMetaListResponse) SetBody(v *QueryUserTagMetaListResponseBody) *QueryUserTagMetaListResponse {
	s.Body = v
	return s
}

func (s *QueryUserTagMetaListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
