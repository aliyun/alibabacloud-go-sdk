// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDocInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDocInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDocInfoResponseBody) *GetDocInfoResponse
	GetBody() *GetDocInfoResponseBody
}

type GetDocInfoResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDocInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDocInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDocInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDocInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDocInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDocInfoResponse) GetBody() *GetDocInfoResponseBody {
	return s.Body
}

func (s *GetDocInfoResponse) SetHeaders(v map[string]*string) *GetDocInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDocInfoResponse) SetStatusCode(v int32) *GetDocInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDocInfoResponse) SetBody(v *GetDocInfoResponseBody) *GetDocInfoResponse {
	s.Body = v
	return s
}

func (s *GetDocInfoResponse) Validate() error {
	return dara.Validate(s)
}
