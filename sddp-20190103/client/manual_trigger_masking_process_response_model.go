// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualTriggerMaskingProcessResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManualTriggerMaskingProcessResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManualTriggerMaskingProcessResponse
	GetStatusCode() *int32
	SetBody(v *ManualTriggerMaskingProcessResponseBody) *ManualTriggerMaskingProcessResponse
	GetBody() *ManualTriggerMaskingProcessResponseBody
}

type ManualTriggerMaskingProcessResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualTriggerMaskingProcessResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualTriggerMaskingProcessResponse) String() string {
	return dara.Prettify(s)
}

func (s ManualTriggerMaskingProcessResponse) GoString() string {
	return s.String()
}

func (s *ManualTriggerMaskingProcessResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManualTriggerMaskingProcessResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManualTriggerMaskingProcessResponse) GetBody() *ManualTriggerMaskingProcessResponseBody {
	return s.Body
}

func (s *ManualTriggerMaskingProcessResponse) SetHeaders(v map[string]*string) *ManualTriggerMaskingProcessResponse {
	s.Headers = v
	return s
}

func (s *ManualTriggerMaskingProcessResponse) SetStatusCode(v int32) *ManualTriggerMaskingProcessResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualTriggerMaskingProcessResponse) SetBody(v *ManualTriggerMaskingProcessResponseBody) *ManualTriggerMaskingProcessResponse {
	s.Body = v
	return s
}

func (s *ManualTriggerMaskingProcessResponse) Validate() error {
	return dara.Validate(s)
}
