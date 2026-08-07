// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloneVoiceModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloneVoiceModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloneVoiceModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListCloneVoiceModelsResponseBody) *ListCloneVoiceModelsResponse
	GetBody() *ListCloneVoiceModelsResponseBody
}

type ListCloneVoiceModelsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloneVoiceModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloneVoiceModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloneVoiceModelsResponse) GoString() string {
	return s.String()
}

func (s *ListCloneVoiceModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloneVoiceModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloneVoiceModelsResponse) GetBody() *ListCloneVoiceModelsResponseBody {
	return s.Body
}

func (s *ListCloneVoiceModelsResponse) SetHeaders(v map[string]*string) *ListCloneVoiceModelsResponse {
	s.Headers = v
	return s
}

func (s *ListCloneVoiceModelsResponse) SetStatusCode(v int32) *ListCloneVoiceModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloneVoiceModelsResponse) SetBody(v *ListCloneVoiceModelsResponseBody) *ListCloneVoiceModelsResponse {
	s.Body = v
	return s
}

func (s *ListCloneVoiceModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
