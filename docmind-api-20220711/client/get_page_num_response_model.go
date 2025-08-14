// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageNumResponse
	GetStatusCode() *int32
	SetBody(v *GetPageNumResponseBody) *GetPageNumResponse
	GetBody() *GetPageNumResponseBody
}

type GetPageNumResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageNumResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageNumResponse) GoString() string {
	return s.String()
}

func (s *GetPageNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageNumResponse) GetBody() *GetPageNumResponseBody {
	return s.Body
}

func (s *GetPageNumResponse) SetHeaders(v map[string]*string) *GetPageNumResponse {
	s.Headers = v
	return s
}

func (s *GetPageNumResponse) SetStatusCode(v int32) *GetPageNumResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageNumResponse) SetBody(v *GetPageNumResponseBody) *GetPageNumResponse {
	s.Body = v
	return s
}

func (s *GetPageNumResponse) Validate() error {
	return dara.Validate(s)
}
