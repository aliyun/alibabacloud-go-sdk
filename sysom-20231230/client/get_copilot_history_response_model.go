// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCopilotHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCopilotHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCopilotHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetCopilotHistoryResponseBody) *GetCopilotHistoryResponse
	GetBody() *GetCopilotHistoryResponseBody
}

type GetCopilotHistoryResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCopilotHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCopilotHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCopilotHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetCopilotHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCopilotHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCopilotHistoryResponse) GetBody() *GetCopilotHistoryResponseBody {
	return s.Body
}

func (s *GetCopilotHistoryResponse) SetHeaders(v map[string]*string) *GetCopilotHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetCopilotHistoryResponse) SetStatusCode(v int32) *GetCopilotHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCopilotHistoryResponse) SetBody(v *GetCopilotHistoryResponseBody) *GetCopilotHistoryResponse {
	s.Body = v
	return s
}

func (s *GetCopilotHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
