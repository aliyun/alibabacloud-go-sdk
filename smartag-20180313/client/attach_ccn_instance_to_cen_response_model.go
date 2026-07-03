// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachCcnInstanceToCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AttachCcnInstanceToCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AttachCcnInstanceToCenResponse
	GetStatusCode() *int32
	SetBody(v *AttachCcnInstanceToCenResponseBody) *AttachCcnInstanceToCenResponse
	GetBody() *AttachCcnInstanceToCenResponseBody
}

type AttachCcnInstanceToCenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AttachCcnInstanceToCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AttachCcnInstanceToCenResponse) String() string {
	return dara.Prettify(s)
}

func (s AttachCcnInstanceToCenResponse) GoString() string {
	return s.String()
}

func (s *AttachCcnInstanceToCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AttachCcnInstanceToCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AttachCcnInstanceToCenResponse) GetBody() *AttachCcnInstanceToCenResponseBody {
	return s.Body
}

func (s *AttachCcnInstanceToCenResponse) SetHeaders(v map[string]*string) *AttachCcnInstanceToCenResponse {
	s.Headers = v
	return s
}

func (s *AttachCcnInstanceToCenResponse) SetStatusCode(v int32) *AttachCcnInstanceToCenResponse {
	s.StatusCode = &v
	return s
}

func (s *AttachCcnInstanceToCenResponse) SetBody(v *AttachCcnInstanceToCenResponseBody) *AttachCcnInstanceToCenResponse {
	s.Body = v
	return s
}

func (s *AttachCcnInstanceToCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
