// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerProcessTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerProcessTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerProcessTaskResponse
	GetStatusCode() *int32
	SetBody(v *TriggerProcessTaskResponseBody) *TriggerProcessTaskResponse
	GetBody() *TriggerProcessTaskResponseBody
}

type TriggerProcessTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerProcessTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerProcessTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerProcessTaskResponse) GoString() string {
	return s.String()
}

func (s *TriggerProcessTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerProcessTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerProcessTaskResponse) GetBody() *TriggerProcessTaskResponseBody {
	return s.Body
}

func (s *TriggerProcessTaskResponse) SetHeaders(v map[string]*string) *TriggerProcessTaskResponse {
	s.Headers = v
	return s
}

func (s *TriggerProcessTaskResponse) SetStatusCode(v int32) *TriggerProcessTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerProcessTaskResponse) SetBody(v *TriggerProcessTaskResponseBody) *TriggerProcessTaskResponse {
	s.Body = v
	return s
}

func (s *TriggerProcessTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
