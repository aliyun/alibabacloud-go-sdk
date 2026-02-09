// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSummaryGenerateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSummaryGenerateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSummaryGenerateResponse
	GetStatusCode() *int32
	SetId(v string) *RunSummaryGenerateResponse
	GetId() *string
	SetEvent(v string) *RunSummaryGenerateResponse
	GetEvent() *string
	SetBody(v *RunSummaryGenerateResponseBody) *RunSummaryGenerateResponse
	GetBody() *RunSummaryGenerateResponseBody
}

type RunSummaryGenerateResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunSummaryGenerateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSummaryGenerateResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSummaryGenerateResponse) GoString() string {
	return s.String()
}

func (s *RunSummaryGenerateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSummaryGenerateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSummaryGenerateResponse) GetId() *string {
	return s.Id
}

func (s *RunSummaryGenerateResponse) GetEvent() *string {
	return s.Event
}

func (s *RunSummaryGenerateResponse) GetBody() *RunSummaryGenerateResponseBody {
	return s.Body
}

func (s *RunSummaryGenerateResponse) SetHeaders(v map[string]*string) *RunSummaryGenerateResponse {
	s.Headers = v
	return s
}

func (s *RunSummaryGenerateResponse) SetStatusCode(v int32) *RunSummaryGenerateResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSummaryGenerateResponse) SetId(v string) *RunSummaryGenerateResponse {
	s.Id = &v
	return s
}

func (s *RunSummaryGenerateResponse) SetEvent(v string) *RunSummaryGenerateResponse {
	s.Event = &v
	return s
}

func (s *RunSummaryGenerateResponse) SetBody(v *RunSummaryGenerateResponseBody) *RunSummaryGenerateResponse {
	s.Body = v
	return s
}

func (s *RunSummaryGenerateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
