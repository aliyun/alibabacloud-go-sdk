// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTitleGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTitleGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTitleGenerationResponse
	GetStatusCode() *int32
	SetId(v string) *RunTitleGenerationResponse
	GetId() *string
	SetEvent(v string) *RunTitleGenerationResponse
	GetEvent() *string
	SetBody(v *RunTitleGenerationResponseBody) *RunTitleGenerationResponse
	GetBody() *RunTitleGenerationResponseBody
}

type RunTitleGenerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                         `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                         `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunTitleGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTitleGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTitleGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTitleGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunTitleGenerationResponse) GetEvent() *string {
	return s.Event
}

func (s *RunTitleGenerationResponse) GetBody() *RunTitleGenerationResponseBody {
	return s.Body
}

func (s *RunTitleGenerationResponse) SetHeaders(v map[string]*string) *RunTitleGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunTitleGenerationResponse) SetStatusCode(v int32) *RunTitleGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTitleGenerationResponse) SetId(v string) *RunTitleGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunTitleGenerationResponse) SetEvent(v string) *RunTitleGenerationResponse {
	s.Event = &v
	return s
}

func (s *RunTitleGenerationResponse) SetBody(v *RunTitleGenerationResponseBody) *RunTitleGenerationResponse {
	s.Body = v
	return s
}

func (s *RunTitleGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
