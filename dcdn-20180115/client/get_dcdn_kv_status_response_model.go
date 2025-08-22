// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDcdnKvStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDcdnKvStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDcdnKvStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetDcdnKvStatusResponseBody) *GetDcdnKvStatusResponse
	GetBody() *GetDcdnKvStatusResponseBody
}

type GetDcdnKvStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDcdnKvStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDcdnKvStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDcdnKvStatusResponse) GoString() string {
	return s.String()
}

func (s *GetDcdnKvStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDcdnKvStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDcdnKvStatusResponse) GetBody() *GetDcdnKvStatusResponseBody {
	return s.Body
}

func (s *GetDcdnKvStatusResponse) SetHeaders(v map[string]*string) *GetDcdnKvStatusResponse {
	s.Headers = v
	return s
}

func (s *GetDcdnKvStatusResponse) SetStatusCode(v int32) *GetDcdnKvStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDcdnKvStatusResponse) SetBody(v *GetDcdnKvStatusResponseBody) *GetDcdnKvStatusResponse {
	s.Body = v
	return s
}

func (s *GetDcdnKvStatusResponse) Validate() error {
	return dara.Validate(s)
}
