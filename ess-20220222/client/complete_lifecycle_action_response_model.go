// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCompleteLifecycleActionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CompleteLifecycleActionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CompleteLifecycleActionResponse
	GetStatusCode() *int32
	SetBody(v *CompleteLifecycleActionResponseBody) *CompleteLifecycleActionResponse
	GetBody() *CompleteLifecycleActionResponseBody
}

type CompleteLifecycleActionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CompleteLifecycleActionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteLifecycleActionResponse) String() string {
	return dara.Prettify(s)
}

func (s CompleteLifecycleActionResponse) GoString() string {
	return s.String()
}

func (s *CompleteLifecycleActionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CompleteLifecycleActionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CompleteLifecycleActionResponse) GetBody() *CompleteLifecycleActionResponseBody {
	return s.Body
}

func (s *CompleteLifecycleActionResponse) SetHeaders(v map[string]*string) *CompleteLifecycleActionResponse {
	s.Headers = v
	return s
}

func (s *CompleteLifecycleActionResponse) SetStatusCode(v int32) *CompleteLifecycleActionResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteLifecycleActionResponse) SetBody(v *CompleteLifecycleActionResponseBody) *CompleteLifecycleActionResponse {
	s.Body = v
	return s
}

func (s *CompleteLifecycleActionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
