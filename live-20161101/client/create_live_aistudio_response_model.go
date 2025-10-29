// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveAIStudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveAIStudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveAIStudioResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveAIStudioResponseBody) *CreateLiveAIStudioResponse
	GetBody() *CreateLiveAIStudioResponseBody
}

type CreateLiveAIStudioResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveAIStudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveAIStudioResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveAIStudioResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveAIStudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveAIStudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveAIStudioResponse) GetBody() *CreateLiveAIStudioResponseBody {
	return s.Body
}

func (s *CreateLiveAIStudioResponse) SetHeaders(v map[string]*string) *CreateLiveAIStudioResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveAIStudioResponse) SetStatusCode(v int32) *CreateLiveAIStudioResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveAIStudioResponse) SetBody(v *CreateLiveAIStudioResponseBody) *CreateLiveAIStudioResponse {
	s.Body = v
	return s
}

func (s *CreateLiveAIStudioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
