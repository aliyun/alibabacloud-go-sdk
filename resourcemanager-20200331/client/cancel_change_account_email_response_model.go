// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChangeAccountEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelChangeAccountEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelChangeAccountEmailResponse
	GetStatusCode() *int32
	SetBody(v *CancelChangeAccountEmailResponseBody) *CancelChangeAccountEmailResponse
	GetBody() *CancelChangeAccountEmailResponseBody
}

type CancelChangeAccountEmailResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelChangeAccountEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelChangeAccountEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelChangeAccountEmailResponse) GoString() string {
	return s.String()
}

func (s *CancelChangeAccountEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelChangeAccountEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelChangeAccountEmailResponse) GetBody() *CancelChangeAccountEmailResponseBody {
	return s.Body
}

func (s *CancelChangeAccountEmailResponse) SetHeaders(v map[string]*string) *CancelChangeAccountEmailResponse {
	s.Headers = v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetStatusCode(v int32) *CancelChangeAccountEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelChangeAccountEmailResponse) SetBody(v *CancelChangeAccountEmailResponseBody) *CancelChangeAccountEmailResponse {
	s.Body = v
	return s
}

func (s *CancelChangeAccountEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
