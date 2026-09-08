// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAgentDataSemanticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAgentDataSemanticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAgentDataSemanticsResponse
	GetStatusCode() *int32
	SetBody(v *SaveAgentDataSemanticsResponseBody) *SaveAgentDataSemanticsResponse
	GetBody() *SaveAgentDataSemanticsResponseBody
}

type SaveAgentDataSemanticsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAgentDataSemanticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAgentDataSemanticsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAgentDataSemanticsResponse) GoString() string {
	return s.String()
}

func (s *SaveAgentDataSemanticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAgentDataSemanticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAgentDataSemanticsResponse) GetBody() *SaveAgentDataSemanticsResponseBody {
	return s.Body
}

func (s *SaveAgentDataSemanticsResponse) SetHeaders(v map[string]*string) *SaveAgentDataSemanticsResponse {
	s.Headers = v
	return s
}

func (s *SaveAgentDataSemanticsResponse) SetStatusCode(v int32) *SaveAgentDataSemanticsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAgentDataSemanticsResponse) SetBody(v *SaveAgentDataSemanticsResponseBody) *SaveAgentDataSemanticsResponse {
	s.Body = v
	return s
}

func (s *SaveAgentDataSemanticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
