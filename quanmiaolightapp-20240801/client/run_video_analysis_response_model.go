// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunVideoAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunVideoAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunVideoAnalysisResponse
	GetStatusCode() *int32
	SetId(v string) *RunVideoAnalysisResponse
	GetId() *string
	SetEvent(v string) *RunVideoAnalysisResponse
	GetEvent() *string
	SetBody(v *RunVideoAnalysisResponseBody) *RunVideoAnalysisResponse
	GetBody() *RunVideoAnalysisResponseBody
}

type RunVideoAnalysisResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                       `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                       `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunVideoAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunVideoAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunVideoAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunVideoAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunVideoAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunVideoAnalysisResponse) GetId() *string {
	return s.Id
}

func (s *RunVideoAnalysisResponse) GetEvent() *string {
	return s.Event
}

func (s *RunVideoAnalysisResponse) GetBody() *RunVideoAnalysisResponseBody {
	return s.Body
}

func (s *RunVideoAnalysisResponse) SetHeaders(v map[string]*string) *RunVideoAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunVideoAnalysisResponse) SetStatusCode(v int32) *RunVideoAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunVideoAnalysisResponse) SetId(v string) *RunVideoAnalysisResponse {
	s.Id = &v
	return s
}

func (s *RunVideoAnalysisResponse) SetEvent(v string) *RunVideoAnalysisResponse {
	s.Event = &v
	return s
}

func (s *RunVideoAnalysisResponse) SetBody(v *RunVideoAnalysisResponseBody) *RunVideoAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunVideoAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
