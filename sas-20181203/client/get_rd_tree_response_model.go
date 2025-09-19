// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRdTreeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRdTreeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRdTreeResponse
	GetStatusCode() *int32
	SetBody(v *GetRdTreeResponseBody) *GetRdTreeResponse
	GetBody() *GetRdTreeResponseBody
}

type GetRdTreeResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRdTreeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRdTreeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRdTreeResponse) GoString() string {
	return s.String()
}

func (s *GetRdTreeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRdTreeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRdTreeResponse) GetBody() *GetRdTreeResponseBody {
	return s.Body
}

func (s *GetRdTreeResponse) SetHeaders(v map[string]*string) *GetRdTreeResponse {
	s.Headers = v
	return s
}

func (s *GetRdTreeResponse) SetStatusCode(v int32) *GetRdTreeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRdTreeResponse) SetBody(v *GetRdTreeResponseBody) *GetRdTreeResponse {
	s.Body = v
	return s
}

func (s *GetRdTreeResponse) Validate() error {
	return dara.Validate(s)
}
