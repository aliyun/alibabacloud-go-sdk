// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecuteOperationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExecuteOperationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExecuteOperationResultResponse
	GetStatusCode() *int32
	SetBody(v *GetExecuteOperationResultResponseBody) *GetExecuteOperationResultResponse
	GetBody() *GetExecuteOperationResultResponseBody
}

type GetExecuteOperationResultResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExecuteOperationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExecuteOperationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExecuteOperationResultResponse) GoString() string {
	return s.String()
}

func (s *GetExecuteOperationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExecuteOperationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExecuteOperationResultResponse) GetBody() *GetExecuteOperationResultResponseBody {
	return s.Body
}

func (s *GetExecuteOperationResultResponse) SetHeaders(v map[string]*string) *GetExecuteOperationResultResponse {
	s.Headers = v
	return s
}

func (s *GetExecuteOperationResultResponse) SetStatusCode(v int32) *GetExecuteOperationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExecuteOperationResultResponse) SetBody(v *GetExecuteOperationResultResponseBody) *GetExecuteOperationResultResponse {
	s.Body = v
	return s
}

func (s *GetExecuteOperationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
