// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDcdnKvDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDcdnKvDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetDcdnKvDetailResponseBody) *GetDcdnKvDetailResponse
	GetBody() *GetDcdnKvDetailResponseBody
}

type GetDcdnKvDetailResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDcdnKvDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDcdnKvDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvDetailResponse) GoString() string {
	return s.String()
}

func (s *GetDcdnKvDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDcdnKvDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDcdnKvDetailResponse) GetBody() *GetDcdnKvDetailResponseBody {
	return s.Body
}

func (s *GetDcdnKvDetailResponse) SetHeaders(v map[string]*string) *GetDcdnKvDetailResponse {
	s.Headers = v
	return s
}

func (s *GetDcdnKvDetailResponse) SetStatusCode(v int32) *GetDcdnKvDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDcdnKvDetailResponse) SetBody(v *GetDcdnKvDetailResponseBody) *GetDcdnKvDetailResponse {
	s.Body = v
	return s
}

func (s *GetDcdnKvDetailResponse) Validate() error {
	return dara.Validate(s)
}
