// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutionHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecutionHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecutionHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetExecutionHistoryResponseBody) *GetExecutionHistoryResponse
	GetBody() *GetExecutionHistoryResponseBody
}

type GetExecutionHistoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecutionHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecutionHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecutionHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetExecutionHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecutionHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecutionHistoryResponse) GetBody() *GetExecutionHistoryResponseBody {
	return s.Body
}

func (s *GetExecutionHistoryResponse) SetHeaders(v map[string]*string) *GetExecutionHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetExecutionHistoryResponse) SetStatusCode(v int32) *GetExecutionHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecutionHistoryResponse) SetBody(v *GetExecutionHistoryResponseBody) *GetExecutionHistoryResponse {
	s.Body = v
	return s
}

func (s *GetExecutionHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
