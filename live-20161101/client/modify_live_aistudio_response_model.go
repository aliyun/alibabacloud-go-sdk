// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveAIStudioResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLiveAIStudioResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLiveAIStudioResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLiveAIStudioResponseBody) *ModifyLiveAIStudioResponse
	GetBody() *ModifyLiveAIStudioResponseBody
}

type ModifyLiveAIStudioResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLiveAIStudioResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLiveAIStudioResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveAIStudioResponse) GoString() string {
	return s.String()
}

func (s *ModifyLiveAIStudioResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLiveAIStudioResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLiveAIStudioResponse) GetBody() *ModifyLiveAIStudioResponseBody {
	return s.Body
}

func (s *ModifyLiveAIStudioResponse) SetHeaders(v map[string]*string) *ModifyLiveAIStudioResponse {
	s.Headers = v
	return s
}

func (s *ModifyLiveAIStudioResponse) SetStatusCode(v int32) *ModifyLiveAIStudioResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLiveAIStudioResponse) SetBody(v *ModifyLiveAIStudioResponseBody) *ModifyLiveAIStudioResponse {
	s.Body = v
	return s
}

func (s *ModifyLiveAIStudioResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
