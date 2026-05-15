// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteStatementEnabledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecuteStatementEnabledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecuteStatementEnabledResponse
	GetStatusCode() *int32
	SetBody(v *GetExecuteStatementEnabledResponseBody) *GetExecuteStatementEnabledResponse
	GetBody() *GetExecuteStatementEnabledResponseBody
}

type GetExecuteStatementEnabledResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteStatementEnabledResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteStatementEnabledResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteStatementEnabledResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteStatementEnabledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecuteStatementEnabledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecuteStatementEnabledResponse) GetBody() *GetExecuteStatementEnabledResponseBody {
	return s.Body
}

func (s *GetExecuteStatementEnabledResponse) SetHeaders(v map[string]*string) *GetExecuteStatementEnabledResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteStatementEnabledResponse) SetStatusCode(v int32) *GetExecuteStatementEnabledResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteStatementEnabledResponse) SetBody(v *GetExecuteStatementEnabledResponseBody) *GetExecuteStatementEnabledResponse {
	s.Body = v
	return s
}

func (s *GetExecuteStatementEnabledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
