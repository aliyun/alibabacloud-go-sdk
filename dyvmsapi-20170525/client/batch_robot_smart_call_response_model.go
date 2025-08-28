// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRobotSmartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchRobotSmartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchRobotSmartCallResponse
	GetStatusCode() *int32
	SetBody(v *BatchRobotSmartCallResponseBody) *BatchRobotSmartCallResponse
	GetBody() *BatchRobotSmartCallResponseBody
}

type BatchRobotSmartCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchRobotSmartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchRobotSmartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchRobotSmartCallResponse) GoString() string {
	return s.String()
}

func (s *BatchRobotSmartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchRobotSmartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchRobotSmartCallResponse) GetBody() *BatchRobotSmartCallResponseBody {
	return s.Body
}

func (s *BatchRobotSmartCallResponse) SetHeaders(v map[string]*string) *BatchRobotSmartCallResponse {
	s.Headers = v
	return s
}

func (s *BatchRobotSmartCallResponse) SetStatusCode(v int32) *BatchRobotSmartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchRobotSmartCallResponse) SetBody(v *BatchRobotSmartCallResponseBody) *BatchRobotSmartCallResponse {
	s.Body = v
	return s
}

func (s *BatchRobotSmartCallResponse) Validate() error {
	return dara.Validate(s)
}
