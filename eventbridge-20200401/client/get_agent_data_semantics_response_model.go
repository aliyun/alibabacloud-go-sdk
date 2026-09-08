// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentDataSemanticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAgentDataSemanticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAgentDataSemanticsResponse
	GetStatusCode() *int32
	SetBody(v *GetAgentDataSemanticsResponseBody) *GetAgentDataSemanticsResponse
	GetBody() *GetAgentDataSemanticsResponseBody
}

type GetAgentDataSemanticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAgentDataSemanticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAgentDataSemanticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAgentDataSemanticsResponse) GoString() string {
	return s.String()
}

func (s *GetAgentDataSemanticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentDataSemanticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAgentDataSemanticsResponse) GetBody() *GetAgentDataSemanticsResponseBody {
	return s.Body
}

func (s *GetAgentDataSemanticsResponse) SetHeaders(v map[string]*string) *GetAgentDataSemanticsResponse {
	s.Headers = v
	return s
}

func (s *GetAgentDataSemanticsResponse) SetStatusCode(v int32) *GetAgentDataSemanticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAgentDataSemanticsResponse) SetBody(v *GetAgentDataSemanticsResponseBody) *GetAgentDataSemanticsResponse {
	s.Body = v
	return s
}

func (s *GetAgentDataSemanticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
