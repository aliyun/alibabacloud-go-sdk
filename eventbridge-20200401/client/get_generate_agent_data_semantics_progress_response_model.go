// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGenerateAgentDataSemanticsProgressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetGenerateAgentDataSemanticsProgressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetGenerateAgentDataSemanticsProgressResponse
	GetStatusCode() *int32
	SetBody(v *GetGenerateAgentDataSemanticsProgressResponseBody) *GetGenerateAgentDataSemanticsProgressResponse
	GetBody() *GetGenerateAgentDataSemanticsProgressResponseBody
}

type GetGenerateAgentDataSemanticsProgressResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetGenerateAgentDataSemanticsProgressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGenerateAgentDataSemanticsProgressResponse) String() string {
	return dara.Prettify(s)
}

func (s GetGenerateAgentDataSemanticsProgressResponse) GoString() string {
	return s.String()
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) GetBody() *GetGenerateAgentDataSemanticsProgressResponseBody {
	return s.Body
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) SetHeaders(v map[string]*string) *GetGenerateAgentDataSemanticsProgressResponse {
	s.Headers = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) SetStatusCode(v int32) *GetGenerateAgentDataSemanticsProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) SetBody(v *GetGenerateAgentDataSemanticsProgressResponseBody) *GetGenerateAgentDataSemanticsProgressResponse {
	s.Body = v
	return s
}

func (s *GetGenerateAgentDataSemanticsProgressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
