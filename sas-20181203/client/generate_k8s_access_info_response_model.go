// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateK8sAccessInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateK8sAccessInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateK8sAccessInfoResponse
	GetStatusCode() *int32
	SetBody(v *GenerateK8sAccessInfoResponseBody) *GenerateK8sAccessInfoResponse
	GetBody() *GenerateK8sAccessInfoResponseBody
}

type GenerateK8sAccessInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateK8sAccessInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateK8sAccessInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateK8sAccessInfoResponse) GoString() string {
	return s.String()
}

func (s *GenerateK8sAccessInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateK8sAccessInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateK8sAccessInfoResponse) GetBody() *GenerateK8sAccessInfoResponseBody {
	return s.Body
}

func (s *GenerateK8sAccessInfoResponse) SetHeaders(v map[string]*string) *GenerateK8sAccessInfoResponse {
	s.Headers = v
	return s
}

func (s *GenerateK8sAccessInfoResponse) SetStatusCode(v int32) *GenerateK8sAccessInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateK8sAccessInfoResponse) SetBody(v *GenerateK8sAccessInfoResponseBody) *GenerateK8sAccessInfoResponse {
	s.Body = v
	return s
}

func (s *GenerateK8sAccessInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
