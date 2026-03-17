// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTagMiningAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTagMiningAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTagMiningAnalysisResponse
	GetStatusCode() *int32
	SetId(v string) *RunTagMiningAnalysisResponse
	GetId() *string
	SetEvent(v string) *RunTagMiningAnalysisResponse
	GetEvent() *string
	SetBody(v *RunTagMiningAnalysisResponseBody) *RunTagMiningAnalysisResponse
	GetBody() *RunTagMiningAnalysisResponseBody
}

type RunTagMiningAnalysisResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                           `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                           `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunTagMiningAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTagMiningAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTagMiningAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunTagMiningAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTagMiningAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTagMiningAnalysisResponse) GetId() *string {
	return s.Id
}

func (s *RunTagMiningAnalysisResponse) GetEvent() *string {
	return s.Event
}

func (s *RunTagMiningAnalysisResponse) GetBody() *RunTagMiningAnalysisResponseBody {
	return s.Body
}

func (s *RunTagMiningAnalysisResponse) SetHeaders(v map[string]*string) *RunTagMiningAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetStatusCode(v int32) *RunTagMiningAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetId(v string) *RunTagMiningAnalysisResponse {
	s.Id = &v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetEvent(v string) *RunTagMiningAnalysisResponse {
	s.Event = &v
	return s
}

func (s *RunTagMiningAnalysisResponse) SetBody(v *RunTagMiningAnalysisResponseBody) *RunTagMiningAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunTagMiningAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
