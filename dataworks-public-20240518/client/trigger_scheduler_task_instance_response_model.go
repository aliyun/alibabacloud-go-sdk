// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSchedulerTaskInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerSchedulerTaskInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerSchedulerTaskInstanceResponse
	GetStatusCode() *int32
	SetBody(v *TriggerSchedulerTaskInstanceResponseBody) *TriggerSchedulerTaskInstanceResponse
	GetBody() *TriggerSchedulerTaskInstanceResponseBody
}

type TriggerSchedulerTaskInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerSchedulerTaskInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerSchedulerTaskInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerSchedulerTaskInstanceResponse) GoString() string {
	return s.String()
}

func (s *TriggerSchedulerTaskInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerSchedulerTaskInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerSchedulerTaskInstanceResponse) GetBody() *TriggerSchedulerTaskInstanceResponseBody {
	return s.Body
}

func (s *TriggerSchedulerTaskInstanceResponse) SetHeaders(v map[string]*string) *TriggerSchedulerTaskInstanceResponse {
	s.Headers = v
	return s
}

func (s *TriggerSchedulerTaskInstanceResponse) SetStatusCode(v int32) *TriggerSchedulerTaskInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerSchedulerTaskInstanceResponse) SetBody(v *TriggerSchedulerTaskInstanceResponseBody) *TriggerSchedulerTaskInstanceResponse {
	s.Body = v
	return s
}

func (s *TriggerSchedulerTaskInstanceResponse) Validate() error {
	return dara.Validate(s)
}
