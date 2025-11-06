// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNacosMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNacosMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNacosMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *GetNacosMcpServerResponseBody) *GetNacosMcpServerResponse
	GetBody() *GetNacosMcpServerResponseBody
}

type GetNacosMcpServerResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNacosMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNacosMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNacosMcpServerResponse) GoString() string {
	return s.String()
}

func (s *GetNacosMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNacosMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNacosMcpServerResponse) GetBody() *GetNacosMcpServerResponseBody {
	return s.Body
}

func (s *GetNacosMcpServerResponse) SetHeaders(v map[string]*string) *GetNacosMcpServerResponse {
	s.Headers = v
	return s
}

func (s *GetNacosMcpServerResponse) SetStatusCode(v int32) *GetNacosMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNacosMcpServerResponse) SetBody(v *GetNacosMcpServerResponseBody) *GetNacosMcpServerResponse {
	s.Body = v
	return s
}

func (s *GetNacosMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
