// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerSophonPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerSophonPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerSophonPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *TriggerSophonPlaybookResponseBody) *TriggerSophonPlaybookResponse
	GetBody() *TriggerSophonPlaybookResponseBody
}

type TriggerSophonPlaybookResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerSophonPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerSophonPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerSophonPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerSophonPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerSophonPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerSophonPlaybookResponse) GetBody() *TriggerSophonPlaybookResponseBody {
	return s.Body
}

func (s *TriggerSophonPlaybookResponse) SetHeaders(v map[string]*string) *TriggerSophonPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetStatusCode(v int32) *TriggerSophonPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerSophonPlaybookResponse) SetBody(v *TriggerSophonPlaybookResponseBody) *TriggerSophonPlaybookResponse {
	s.Body = v
	return s
}

func (s *TriggerSophonPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
