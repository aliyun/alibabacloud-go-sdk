// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiMcpServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetApiMcpServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetApiMcpServerResponse
	GetStatusCode() *int32
	SetBody(v *GetApiMcpServerResponseBody) *GetApiMcpServerResponse
	GetBody() *GetApiMcpServerResponseBody
}

type GetApiMcpServerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetApiMcpServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetApiMcpServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetApiMcpServerResponse) GoString() string {
	return s.String()
}

func (s *GetApiMcpServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetApiMcpServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetApiMcpServerResponse) GetBody() *GetApiMcpServerResponseBody {
	return s.Body
}

func (s *GetApiMcpServerResponse) SetHeaders(v map[string]*string) *GetApiMcpServerResponse {
	s.Headers = v
	return s
}

func (s *GetApiMcpServerResponse) SetStatusCode(v int32) *GetApiMcpServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetApiMcpServerResponse) SetBody(v *GetApiMcpServerResponseBody) *GetApiMcpServerResponse {
	s.Body = v
	return s
}

func (s *GetApiMcpServerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
