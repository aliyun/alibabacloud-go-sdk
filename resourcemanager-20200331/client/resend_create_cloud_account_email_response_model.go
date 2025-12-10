// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendCreateCloudAccountEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendCreateCloudAccountEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendCreateCloudAccountEmailResponse
	GetStatusCode() *int32
	SetBody(v *ResendCreateCloudAccountEmailResponseBody) *ResendCreateCloudAccountEmailResponse
	GetBody() *ResendCreateCloudAccountEmailResponseBody
}

type ResendCreateCloudAccountEmailResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendCreateCloudAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendCreateCloudAccountEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendCreateCloudAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendCreateCloudAccountEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendCreateCloudAccountEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendCreateCloudAccountEmailResponse) GetBody() *ResendCreateCloudAccountEmailResponseBody {
	return s.Body
}

func (s *ResendCreateCloudAccountEmailResponse) SetHeaders(v map[string]*string) *ResendCreateCloudAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) SetStatusCode(v int32) *ResendCreateCloudAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) SetBody(v *ResendCreateCloudAccountEmailResponseBody) *ResendCreateCloudAccountEmailResponse {
	s.Body = v
	return s
}

func (s *ResendCreateCloudAccountEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
