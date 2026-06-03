// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmTransferInEmailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmTransferInEmailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmTransferInEmailResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmTransferInEmailResponseBody) *ConfirmTransferInEmailResponse
	GetBody() *ConfirmTransferInEmailResponseBody
}

type ConfirmTransferInEmailResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmTransferInEmailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmTransferInEmailResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmTransferInEmailResponse) GoString() string {
	return s.String()
}

func (s *ConfirmTransferInEmailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmTransferInEmailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmTransferInEmailResponse) GetBody() *ConfirmTransferInEmailResponseBody {
	return s.Body
}

func (s *ConfirmTransferInEmailResponse) SetHeaders(v map[string]*string) *ConfirmTransferInEmailResponse {
	s.Headers = v
	return s
}

func (s *ConfirmTransferInEmailResponse) SetStatusCode(v int32) *ConfirmTransferInEmailResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmTransferInEmailResponse) SetBody(v *ConfirmTransferInEmailResponseBody) *ConfirmTransferInEmailResponse {
	s.Body = v
	return s
}

func (s *ConfirmTransferInEmailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
