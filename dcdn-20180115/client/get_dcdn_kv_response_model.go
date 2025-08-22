// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDcdnKvResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDcdnKvResponse
	GetStatusCode() *int32
	SetBody(v *GetDcdnKvResponseBody) *GetDcdnKvResponse
	GetBody() *GetDcdnKvResponseBody
}

type GetDcdnKvResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDcdnKvResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDcdnKvResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvResponse) GoString() string {
	return s.String()
}

func (s *GetDcdnKvResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDcdnKvResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDcdnKvResponse) GetBody() *GetDcdnKvResponseBody {
	return s.Body
}

func (s *GetDcdnKvResponse) SetHeaders(v map[string]*string) *GetDcdnKvResponse {
	s.Headers = v
	return s
}

func (s *GetDcdnKvResponse) SetStatusCode(v int32) *GetDcdnKvResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDcdnKvResponse) SetBody(v *GetDcdnKvResponseBody) *GetDcdnKvResponse {
	s.Body = v
	return s
}

func (s *GetDcdnKvResponse) Validate() error {
	return dara.Validate(s)
}
