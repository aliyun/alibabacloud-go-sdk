// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResendPromoteResourceAccountEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResendPromoteResourceAccountEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResendPromoteResourceAccountEmailResponse
	GetStatusCode() *int32
	SetBody(v *ResendPromoteResourceAccountEmailResponseBody) *ResendPromoteResourceAccountEmailResponse
	GetBody() *ResendPromoteResourceAccountEmailResponseBody
}

type ResendPromoteResourceAccountEmailResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResendPromoteResourceAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResendPromoteResourceAccountEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ResendPromoteResourceAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *ResendPromoteResourceAccountEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResendPromoteResourceAccountEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResendPromoteResourceAccountEmailResponse) GetBody() *ResendPromoteResourceAccountEmailResponseBody {
	return s.Body
}

func (s *ResendPromoteResourceAccountEmailResponse) SetHeaders(v map[string]*string) *ResendPromoteResourceAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) SetStatusCode(v int32) *ResendPromoteResourceAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) SetBody(v *ResendPromoteResourceAccountEmailResponseBody) *ResendPromoteResourceAccountEmailResponse {
	s.Body = v
	return s
}

func (s *ResendPromoteResourceAccountEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
