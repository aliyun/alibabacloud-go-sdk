// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerRecordResponse
	GetStatusCode() *int32
	SetBody(v *TriggerRecordResponseBody) *TriggerRecordResponse
	GetBody() *TriggerRecordResponseBody
}

type TriggerRecordResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerRecordResponse) GoString() string {
	return s.String()
}

func (s *TriggerRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerRecordResponse) GetBody() *TriggerRecordResponseBody {
	return s.Body
}

func (s *TriggerRecordResponse) SetHeaders(v map[string]*string) *TriggerRecordResponse {
	s.Headers = v
	return s
}

func (s *TriggerRecordResponse) SetStatusCode(v int32) *TriggerRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerRecordResponse) SetBody(v *TriggerRecordResponseBody) *TriggerRecordResponse {
	s.Body = v
	return s
}

func (s *TriggerRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
