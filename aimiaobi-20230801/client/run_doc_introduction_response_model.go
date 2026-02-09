// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocIntroductionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocIntroductionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocIntroductionResponse
	GetStatusCode() *int32
	SetId(v string) *RunDocIntroductionResponse
	GetId() *string
	SetEvent(v string) *RunDocIntroductionResponse
	GetEvent() *string
	SetBody(v *RunDocIntroductionResponseBody) *RunDocIntroductionResponse
	GetBody() *RunDocIntroductionResponseBody
}

type RunDocIntroductionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDocIntroductionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocIntroductionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocIntroductionResponse) GoString() string {
	return s.String()
}

func (s *RunDocIntroductionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocIntroductionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocIntroductionResponse) GetId() *string {
	return s.Id
}

func (s *RunDocIntroductionResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDocIntroductionResponse) GetBody() *RunDocIntroductionResponseBody {
	return s.Body
}

func (s *RunDocIntroductionResponse) SetHeaders(v map[string]*string) *RunDocIntroductionResponse {
	s.Headers = v
	return s
}

func (s *RunDocIntroductionResponse) SetStatusCode(v int32) *RunDocIntroductionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocIntroductionResponse) SetId(v string) *RunDocIntroductionResponse {
	s.Id = &v
	return s
}

func (s *RunDocIntroductionResponse) SetEvent(v string) *RunDocIntroductionResponse {
	s.Event = &v
	return s
}

func (s *RunDocIntroductionResponse) SetBody(v *RunDocIntroductionResponseBody) *RunDocIntroductionResponse {
	s.Body = v
	return s
}

func (s *RunDocIntroductionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
