// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDialogAnalysisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDialogAnalysisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDialogAnalysisResponse
	GetStatusCode() *int32
	SetId(v string) *RunDialogAnalysisResponse
	GetId() *string
	SetEvent(v string) *RunDialogAnalysisResponse
	GetEvent() *string
	SetBody(v *RunDialogAnalysisResponseBody) *RunDialogAnalysisResponse
	GetBody() *RunDialogAnalysisResponseBody
}

type RunDialogAnalysisResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDialogAnalysisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDialogAnalysisResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDialogAnalysisResponse) GoString() string {
	return s.String()
}

func (s *RunDialogAnalysisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDialogAnalysisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDialogAnalysisResponse) GetId() *string {
	return s.Id
}

func (s *RunDialogAnalysisResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDialogAnalysisResponse) GetBody() *RunDialogAnalysisResponseBody {
	return s.Body
}

func (s *RunDialogAnalysisResponse) SetHeaders(v map[string]*string) *RunDialogAnalysisResponse {
	s.Headers = v
	return s
}

func (s *RunDialogAnalysisResponse) SetStatusCode(v int32) *RunDialogAnalysisResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDialogAnalysisResponse) SetId(v string) *RunDialogAnalysisResponse {
	s.Id = &v
	return s
}

func (s *RunDialogAnalysisResponse) SetEvent(v string) *RunDialogAnalysisResponse {
	s.Event = &v
	return s
}

func (s *RunDialogAnalysisResponse) SetBody(v *RunDialogAnalysisResponseBody) *RunDialogAnalysisResponse {
	s.Body = v
	return s
}

func (s *RunDialogAnalysisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
