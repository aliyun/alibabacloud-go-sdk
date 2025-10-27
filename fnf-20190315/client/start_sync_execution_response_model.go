// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSyncExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartSyncExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartSyncExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartSyncExecutionResponseBody) *StartSyncExecutionResponse
	GetBody() *StartSyncExecutionResponseBody
}

type StartSyncExecutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartSyncExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartSyncExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartSyncExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartSyncExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartSyncExecutionResponse) GetBody() *StartSyncExecutionResponseBody {
	return s.Body
}

func (s *StartSyncExecutionResponse) SetHeaders(v map[string]*string) *StartSyncExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartSyncExecutionResponse) SetStatusCode(v int32) *StartSyncExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartSyncExecutionResponse) SetBody(v *StartSyncExecutionResponseBody) *StartSyncExecutionResponse {
	s.Body = v
	return s
}

func (s *StartSyncExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
