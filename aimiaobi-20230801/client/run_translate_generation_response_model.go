// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTranslateGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTranslateGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTranslateGenerationResponse
	GetStatusCode() *int32
	SetId(v string) *RunTranslateGenerationResponse
	GetId() *string
	SetEvent(v string) *RunTranslateGenerationResponse
	GetEvent() *string
	SetBody(v *RunTranslateGenerationResponseBody) *RunTranslateGenerationResponse
	GetBody() *RunTranslateGenerationResponseBody
}

type RunTranslateGenerationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                             `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunTranslateGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTranslateGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTranslateGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTranslateGenerationResponse) GetId() *string {
	return s.Id
}

func (s *RunTranslateGenerationResponse) GetEvent() *string {
	return s.Event
}

func (s *RunTranslateGenerationResponse) GetBody() *RunTranslateGenerationResponseBody {
	return s.Body
}

func (s *RunTranslateGenerationResponse) SetHeaders(v map[string]*string) *RunTranslateGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunTranslateGenerationResponse) SetStatusCode(v int32) *RunTranslateGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTranslateGenerationResponse) SetId(v string) *RunTranslateGenerationResponse {
	s.Id = &v
	return s
}

func (s *RunTranslateGenerationResponse) SetEvent(v string) *RunTranslateGenerationResponse {
	s.Event = &v
	return s
}

func (s *RunTranslateGenerationResponse) SetBody(v *RunTranslateGenerationResponseBody) *RunTranslateGenerationResponse {
	s.Body = v
	return s
}

func (s *RunTranslateGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
