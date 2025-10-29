// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *RestartLivePullToPushResponseBody) *RestartLivePullToPushResponse
	GetBody() *RestartLivePullToPushResponseBody
}

type RestartLivePullToPushResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *RestartLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartLivePullToPushResponse) GetBody() *RestartLivePullToPushResponseBody {
	return s.Body
}

func (s *RestartLivePullToPushResponse) SetHeaders(v map[string]*string) *RestartLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *RestartLivePullToPushResponse) SetStatusCode(v int32) *RestartLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartLivePullToPushResponse) SetBody(v *RestartLivePullToPushResponseBody) *RestartLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *RestartLivePullToPushResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
