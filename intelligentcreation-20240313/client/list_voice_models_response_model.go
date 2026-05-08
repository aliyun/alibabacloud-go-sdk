// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListVoiceModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVoiceModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVoiceModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListVoiceModelsResponseBody) *ListVoiceModelsResponse
	GetBody() *ListVoiceModelsResponseBody
}

type ListVoiceModelsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVoiceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVoiceModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVoiceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListVoiceModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVoiceModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVoiceModelsResponse) GetBody() *ListVoiceModelsResponseBody {
	return s.Body
}

func (s *ListVoiceModelsResponse) SetHeaders(v map[string]*string) *ListVoiceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListVoiceModelsResponse) SetStatusCode(v int32) *ListVoiceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVoiceModelsResponse) SetBody(v *ListVoiceModelsResponseBody) *ListVoiceModelsResponse {
	s.Body = v
	return s
}

func (s *ListVoiceModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
