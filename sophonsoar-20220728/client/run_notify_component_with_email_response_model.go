// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNotifyComponentWithEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunNotifyComponentWithEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunNotifyComponentWithEmailResponse
	GetStatusCode() *int32
	SetBody(v *RunNotifyComponentWithEmailResponseBody) *RunNotifyComponentWithEmailResponse
	GetBody() *RunNotifyComponentWithEmailResponseBody
}

type RunNotifyComponentWithEmailResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNotifyComponentWithEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNotifyComponentWithEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s RunNotifyComponentWithEmailResponse) GoString() string {
	return s.String()
}

func (s *RunNotifyComponentWithEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunNotifyComponentWithEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunNotifyComponentWithEmailResponse) GetBody() *RunNotifyComponentWithEmailResponseBody {
	return s.Body
}

func (s *RunNotifyComponentWithEmailResponse) SetHeaders(v map[string]*string) *RunNotifyComponentWithEmailResponse {
	s.Headers = v
	return s
}

func (s *RunNotifyComponentWithEmailResponse) SetStatusCode(v int32) *RunNotifyComponentWithEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNotifyComponentWithEmailResponse) SetBody(v *RunNotifyComponentWithEmailResponseBody) *RunNotifyComponentWithEmailResponse {
	s.Body = v
	return s
}

func (s *RunNotifyComponentWithEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
