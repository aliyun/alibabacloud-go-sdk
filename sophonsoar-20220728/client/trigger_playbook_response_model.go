// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerPlaybookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerPlaybookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerPlaybookResponse
	GetStatusCode() *int32
	SetBody(v *TriggerPlaybookResponseBody) *TriggerPlaybookResponse
	GetBody() *TriggerPlaybookResponseBody
}

type TriggerPlaybookResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerPlaybookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerPlaybookResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerPlaybookResponse) GoString() string {
	return s.String()
}

func (s *TriggerPlaybookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerPlaybookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerPlaybookResponse) GetBody() *TriggerPlaybookResponseBody {
	return s.Body
}

func (s *TriggerPlaybookResponse) SetHeaders(v map[string]*string) *TriggerPlaybookResponse {
	s.Headers = v
	return s
}

func (s *TriggerPlaybookResponse) SetStatusCode(v int32) *TriggerPlaybookResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerPlaybookResponse) SetBody(v *TriggerPlaybookResponseBody) *TriggerPlaybookResponse {
	s.Body = v
	return s
}

func (s *TriggerPlaybookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
