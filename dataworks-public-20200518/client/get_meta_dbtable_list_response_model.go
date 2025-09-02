// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaDBTableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMetaDBTableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMetaDBTableListResponse
	GetStatusCode() *int32
	SetBody(v *GetMetaDBTableListResponseBody) *GetMetaDBTableListResponse
	GetBody() *GetMetaDBTableListResponseBody
}

type GetMetaDBTableListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMetaDBTableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMetaDBTableListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMetaDBTableListResponse) GoString() string {
	return s.String()
}

func (s *GetMetaDBTableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMetaDBTableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMetaDBTableListResponse) GetBody() *GetMetaDBTableListResponseBody {
	return s.Body
}

func (s *GetMetaDBTableListResponse) SetHeaders(v map[string]*string) *GetMetaDBTableListResponse {
	s.Headers = v
	return s
}

func (s *GetMetaDBTableListResponse) SetStatusCode(v int32) *GetMetaDBTableListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMetaDBTableListResponse) SetBody(v *GetMetaDBTableListResponseBody) *GetMetaDBTableListResponse {
	s.Body = v
	return s
}

func (s *GetMetaDBTableListResponse) Validate() error {
	return dara.Validate(s)
}
