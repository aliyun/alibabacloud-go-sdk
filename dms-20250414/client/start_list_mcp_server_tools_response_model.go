// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartListMcpServerToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartListMcpServerToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartListMcpServerToolsResponse
	GetStatusCode() *int32
	SetBody(v *StartListMcpServerToolsResponseBody) *StartListMcpServerToolsResponse
	GetBody() *StartListMcpServerToolsResponseBody
}

type StartListMcpServerToolsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartListMcpServerToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartListMcpServerToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s StartListMcpServerToolsResponse) GoString() string {
	return s.String()
}

func (s *StartListMcpServerToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartListMcpServerToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartListMcpServerToolsResponse) GetBody() *StartListMcpServerToolsResponseBody {
	return s.Body
}

func (s *StartListMcpServerToolsResponse) SetHeaders(v map[string]*string) *StartListMcpServerToolsResponse {
	s.Headers = v
	return s
}

func (s *StartListMcpServerToolsResponse) SetStatusCode(v int32) *StartListMcpServerToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *StartListMcpServerToolsResponse) SetBody(v *StartListMcpServerToolsResponseBody) *StartListMcpServerToolsResponse {
	s.Body = v
	return s
}

func (s *StartListMcpServerToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
