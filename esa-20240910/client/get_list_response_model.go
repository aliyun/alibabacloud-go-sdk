// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetListResponse
	GetStatusCode() *int32
	SetBody(v *GetListResponseBody) *GetListResponse
	GetBody() *GetListResponseBody
}

type GetListResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetListResponse) GoString() string {
	return s.String()
}

func (s *GetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetListResponse) GetBody() *GetListResponseBody {
	return s.Body
}

func (s *GetListResponse) SetHeaders(v map[string]*string) *GetListResponse {
	s.Headers = v
	return s
}

func (s *GetListResponse) SetStatusCode(v int32) *GetListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListResponse) SetBody(v *GetListResponseBody) *GetListResponse {
	s.Body = v
	return s
}

func (s *GetListResponse) Validate() error {
	return dara.Validate(s)
}
