// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessAliMeCallbackOfStagingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProcessAliMeCallbackOfStagingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProcessAliMeCallbackOfStagingResponse
	GetStatusCode() *int32
	SetBody(v *ProcessAliMeCallbackOfStagingResponseBody) *ProcessAliMeCallbackOfStagingResponse
	GetBody() *ProcessAliMeCallbackOfStagingResponseBody
}

type ProcessAliMeCallbackOfStagingResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProcessAliMeCallbackOfStagingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProcessAliMeCallbackOfStagingResponse) String() string {
	return dara.Prettify(s)
}

func (s ProcessAliMeCallbackOfStagingResponse) GoString() string {
	return s.String()
}

func (s *ProcessAliMeCallbackOfStagingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProcessAliMeCallbackOfStagingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProcessAliMeCallbackOfStagingResponse) GetBody() *ProcessAliMeCallbackOfStagingResponseBody {
	return s.Body
}

func (s *ProcessAliMeCallbackOfStagingResponse) SetHeaders(v map[string]*string) *ProcessAliMeCallbackOfStagingResponse {
	s.Headers = v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponse) SetStatusCode(v int32) *ProcessAliMeCallbackOfStagingResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponse) SetBody(v *ProcessAliMeCallbackOfStagingResponseBody) *ProcessAliMeCallbackOfStagingResponse {
	s.Body = v
	return s
}

func (s *ProcessAliMeCallbackOfStagingResponse) Validate() error {
	return dara.Validate(s)
}
