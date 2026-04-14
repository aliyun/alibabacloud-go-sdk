// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCpuHighAgentStreamResponseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CpuHighAgentStreamResponseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CpuHighAgentStreamResponseResponse
	GetStatusCode() *int32
	SetId(v string) *CpuHighAgentStreamResponseResponse
	GetId() *string
	SetEvent(v string) *CpuHighAgentStreamResponseResponse
	GetEvent() *string
	SetBody(v *CpuHighAgentStreamResponseResponseBody) *CpuHighAgentStreamResponseResponse
	GetBody() *CpuHighAgentStreamResponseResponseBody
}

type CpuHighAgentStreamResponseResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                                 `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                                 `json:"event,omitempty" xml:"event,omitempty"`
	Body       *CpuHighAgentStreamResponseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CpuHighAgentStreamResponseResponse) String() string {
	return dara.Prettify(s)
}

func (s CpuHighAgentStreamResponseResponse) GoString() string {
	return s.String()
}

func (s *CpuHighAgentStreamResponseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CpuHighAgentStreamResponseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CpuHighAgentStreamResponseResponse) GetId() *string {
	return s.Id
}

func (s *CpuHighAgentStreamResponseResponse) GetEvent() *string {
	return s.Event
}

func (s *CpuHighAgentStreamResponseResponse) GetBody() *CpuHighAgentStreamResponseResponseBody {
	return s.Body
}

func (s *CpuHighAgentStreamResponseResponse) SetHeaders(v map[string]*string) *CpuHighAgentStreamResponseResponse {
	s.Headers = v
	return s
}

func (s *CpuHighAgentStreamResponseResponse) SetStatusCode(v int32) *CpuHighAgentStreamResponseResponse {
	s.StatusCode = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponse) SetId(v string) *CpuHighAgentStreamResponseResponse {
	s.Id = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponse) SetEvent(v string) *CpuHighAgentStreamResponseResponse {
	s.Event = &v
	return s
}

func (s *CpuHighAgentStreamResponseResponse) SetBody(v *CpuHighAgentStreamResponseResponseBody) *CpuHighAgentStreamResponseResponse {
	s.Body = v
	return s
}

func (s *CpuHighAgentStreamResponseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
