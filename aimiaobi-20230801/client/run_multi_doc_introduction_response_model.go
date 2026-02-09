// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMultiDocIntroductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunMultiDocIntroductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunMultiDocIntroductionResponse
	GetStatusCode() *int32
	SetId(v string) *RunMultiDocIntroductionResponse
	GetId() *string
	SetEvent(v string) *RunMultiDocIntroductionResponse
	GetEvent() *string
	SetBody(v *RunMultiDocIntroductionResponseBody) *RunMultiDocIntroductionResponse
	GetBody() *RunMultiDocIntroductionResponseBody
}

type RunMultiDocIntroductionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                              `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                              `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunMultiDocIntroductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunMultiDocIntroductionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunMultiDocIntroductionResponse) GoString() string {
	return s.String()
}

func (s *RunMultiDocIntroductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunMultiDocIntroductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunMultiDocIntroductionResponse) GetId() *string {
	return s.Id
}

func (s *RunMultiDocIntroductionResponse) GetEvent() *string {
	return s.Event
}

func (s *RunMultiDocIntroductionResponse) GetBody() *RunMultiDocIntroductionResponseBody {
	return s.Body
}

func (s *RunMultiDocIntroductionResponse) SetHeaders(v map[string]*string) *RunMultiDocIntroductionResponse {
	s.Headers = v
	return s
}

func (s *RunMultiDocIntroductionResponse) SetStatusCode(v int32) *RunMultiDocIntroductionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunMultiDocIntroductionResponse) SetId(v string) *RunMultiDocIntroductionResponse {
	s.Id = &v
	return s
}

func (s *RunMultiDocIntroductionResponse) SetEvent(v string) *RunMultiDocIntroductionResponse {
	s.Event = &v
	return s
}

func (s *RunMultiDocIntroductionResponse) SetBody(v *RunMultiDocIntroductionResponseBody) *RunMultiDocIntroductionResponse {
	s.Body = v
	return s
}

func (s *RunMultiDocIntroductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
