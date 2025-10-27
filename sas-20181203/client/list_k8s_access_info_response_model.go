// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sAccessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListK8sAccessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListK8sAccessInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListK8sAccessInfoResponseBody) *ListK8sAccessInfoResponse
	GetBody() *ListK8sAccessInfoResponseBody
}

type ListK8sAccessInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListK8sAccessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListK8sAccessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListK8sAccessInfoResponse) GoString() string {
	return s.String()
}

func (s *ListK8sAccessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListK8sAccessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListK8sAccessInfoResponse) GetBody() *ListK8sAccessInfoResponseBody {
	return s.Body
}

func (s *ListK8sAccessInfoResponse) SetHeaders(v map[string]*string) *ListK8sAccessInfoResponse {
	s.Headers = v
	return s
}

func (s *ListK8sAccessInfoResponse) SetStatusCode(v int32) *ListK8sAccessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListK8sAccessInfoResponse) SetBody(v *ListK8sAccessInfoResponseBody) *ListK8sAccessInfoResponse {
	s.Body = v
	return s
}

func (s *ListK8sAccessInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
