// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPatrolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerPatrolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerPatrolResponse
	GetStatusCode() *int32
	SetBody(v *TriggerPatrolResponseBody) *TriggerPatrolResponse
	GetBody() *TriggerPatrolResponseBody
}

type TriggerPatrolResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerPatrolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerPatrolResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerPatrolResponse) GoString() string {
	return s.String()
}

func (s *TriggerPatrolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerPatrolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerPatrolResponse) GetBody() *TriggerPatrolResponseBody {
	return s.Body
}

func (s *TriggerPatrolResponse) SetHeaders(v map[string]*string) *TriggerPatrolResponse {
	s.Headers = v
	return s
}

func (s *TriggerPatrolResponse) SetStatusCode(v int32) *TriggerPatrolResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerPatrolResponse) SetBody(v *TriggerPatrolResponseBody) *TriggerPatrolResponse {
	s.Body = v
	return s
}

func (s *TriggerPatrolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
