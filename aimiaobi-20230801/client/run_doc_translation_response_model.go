// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocTranslationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunDocTranslationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunDocTranslationResponse
	GetStatusCode() *int32
	SetId(v string) *RunDocTranslationResponse
	GetId() *string
	SetEvent(v string) *RunDocTranslationResponse
	GetEvent() *string
	SetBody(v *RunDocTranslationResponseBody) *RunDocTranslationResponse
	GetBody() *RunDocTranslationResponseBody
}

type RunDocTranslationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Id         *string                        `json:"id,omitempty" xml:"id,omitempty"`
	Event      *string                        `json:"event,omitempty" xml:"event,omitempty"`
	Body       *RunDocTranslationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunDocTranslationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunDocTranslationResponse) GoString() string {
	return s.String()
}

func (s *RunDocTranslationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunDocTranslationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunDocTranslationResponse) GetId() *string {
	return s.Id
}

func (s *RunDocTranslationResponse) GetEvent() *string {
	return s.Event
}

func (s *RunDocTranslationResponse) GetBody() *RunDocTranslationResponseBody {
	return s.Body
}

func (s *RunDocTranslationResponse) SetHeaders(v map[string]*string) *RunDocTranslationResponse {
	s.Headers = v
	return s
}

func (s *RunDocTranslationResponse) SetStatusCode(v int32) *RunDocTranslationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunDocTranslationResponse) SetId(v string) *RunDocTranslationResponse {
	s.Id = &v
	return s
}

func (s *RunDocTranslationResponse) SetEvent(v string) *RunDocTranslationResponse {
	s.Event = &v
	return s
}

func (s *RunDocTranslationResponse) SetBody(v *RunDocTranslationResponseBody) *RunDocTranslationResponse {
	s.Body = v
	return s
}

func (s *RunDocTranslationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
