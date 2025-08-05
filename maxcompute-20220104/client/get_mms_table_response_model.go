// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMmsTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMmsTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMmsTableResponse
	GetStatusCode() *int32
	SetBody(v *GetMmsTableResponseBody) *GetMmsTableResponse
	GetBody() *GetMmsTableResponseBody
}

type GetMmsTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMmsTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMmsTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMmsTableResponse) GoString() string {
	return s.String()
}

func (s *GetMmsTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMmsTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMmsTableResponse) GetBody() *GetMmsTableResponseBody {
	return s.Body
}

func (s *GetMmsTableResponse) SetHeaders(v map[string]*string) *GetMmsTableResponse {
	s.Headers = v
	return s
}

func (s *GetMmsTableResponse) SetStatusCode(v int32) *GetMmsTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMmsTableResponse) SetBody(v *GetMmsTableResponseBody) *GetMmsTableResponse {
	s.Body = v
	return s
}

func (s *GetMmsTableResponse) Validate() error {
	return dara.Validate(s)
}
