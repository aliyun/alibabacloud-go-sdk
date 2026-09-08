// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateAgentDataSemanticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateAgentDataSemanticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateAgentDataSemanticsResponse
	GetStatusCode() *int32
	SetBody(v *GenerateAgentDataSemanticsResponseBody) *GenerateAgentDataSemanticsResponse
	GetBody() *GenerateAgentDataSemanticsResponseBody
}

type GenerateAgentDataSemanticsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateAgentDataSemanticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateAgentDataSemanticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateAgentDataSemanticsResponse) GoString() string {
	return s.String()
}

func (s *GenerateAgentDataSemanticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateAgentDataSemanticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateAgentDataSemanticsResponse) GetBody() *GenerateAgentDataSemanticsResponseBody {
	return s.Body
}

func (s *GenerateAgentDataSemanticsResponse) SetHeaders(v map[string]*string) *GenerateAgentDataSemanticsResponse {
	s.Headers = v
	return s
}

func (s *GenerateAgentDataSemanticsResponse) SetStatusCode(v int32) *GenerateAgentDataSemanticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateAgentDataSemanticsResponse) SetBody(v *GenerateAgentDataSemanticsResponseBody) *GenerateAgentDataSemanticsResponse {
	s.Body = v
	return s
}

func (s *GenerateAgentDataSemanticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
