// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListMcpServerToolsResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetListMcpServerToolsResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetListMcpServerToolsResultResponse
	GetStatusCode() *int32
	SetBody(v *GetListMcpServerToolsResultResponseBody) *GetListMcpServerToolsResultResponse
	GetBody() *GetListMcpServerToolsResultResponseBody
}

type GetListMcpServerToolsResultResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetListMcpServerToolsResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetListMcpServerToolsResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetListMcpServerToolsResultResponse) GoString() string {
	return s.String()
}

func (s *GetListMcpServerToolsResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetListMcpServerToolsResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetListMcpServerToolsResultResponse) GetBody() *GetListMcpServerToolsResultResponseBody {
	return s.Body
}

func (s *GetListMcpServerToolsResultResponse) SetHeaders(v map[string]*string) *GetListMcpServerToolsResultResponse {
	s.Headers = v
	return s
}

func (s *GetListMcpServerToolsResultResponse) SetStatusCode(v int32) *GetListMcpServerToolsResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetListMcpServerToolsResultResponse) SetBody(v *GetListMcpServerToolsResultResponseBody) *GetListMcpServerToolsResultResponse {
	s.Body = v
	return s
}

func (s *GetListMcpServerToolsResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
