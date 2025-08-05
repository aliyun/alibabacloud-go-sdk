// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConnectionResponse
	GetStatusCode() *int32
	SetBody(v *GetConnectionResponseBody) *GetConnectionResponse
	GetBody() *GetConnectionResponseBody
}

type GetConnectionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConnectionResponse) GoString() string {
	return s.String()
}

func (s *GetConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConnectionResponse) GetBody() *GetConnectionResponseBody {
	return s.Body
}

func (s *GetConnectionResponse) SetHeaders(v map[string]*string) *GetConnectionResponse {
	s.Headers = v
	return s
}

func (s *GetConnectionResponse) SetStatusCode(v int32) *GetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectionResponse) SetBody(v *GetConnectionResponseBody) *GetConnectionResponse {
	s.Body = v
	return s
}

func (s *GetConnectionResponse) Validate() error {
	return dara.Validate(s)
}
