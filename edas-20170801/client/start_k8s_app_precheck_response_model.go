// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartK8sAppPrecheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartK8sAppPrecheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartK8sAppPrecheckResponse
	GetStatusCode() *int32
	SetBody(v *StartK8sAppPrecheckResponseBody) *StartK8sAppPrecheckResponse
	GetBody() *StartK8sAppPrecheckResponseBody
}

type StartK8sAppPrecheckResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartK8sAppPrecheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartK8sAppPrecheckResponse) String() string {
	return dara.Prettify(s)
}

func (s StartK8sAppPrecheckResponse) GoString() string {
	return s.String()
}

func (s *StartK8sAppPrecheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartK8sAppPrecheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartK8sAppPrecheckResponse) GetBody() *StartK8sAppPrecheckResponseBody {
	return s.Body
}

func (s *StartK8sAppPrecheckResponse) SetHeaders(v map[string]*string) *StartK8sAppPrecheckResponse {
	s.Headers = v
	return s
}

func (s *StartK8sAppPrecheckResponse) SetStatusCode(v int32) *StartK8sAppPrecheckResponse {
	s.StatusCode = &v
	return s
}

func (s *StartK8sAppPrecheckResponse) SetBody(v *StartK8sAppPrecheckResponseBody) *StartK8sAppPrecheckResponse {
	s.Body = v
	return s
}

func (s *StartK8sAppPrecheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
