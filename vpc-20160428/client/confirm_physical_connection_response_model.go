// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmPhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfirmPhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfirmPhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *ConfirmPhysicalConnectionResponseBody) *ConfirmPhysicalConnectionResponse
	GetBody() *ConfirmPhysicalConnectionResponseBody
}

type ConfirmPhysicalConnectionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfirmPhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmPhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfirmPhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *ConfirmPhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfirmPhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfirmPhysicalConnectionResponse) GetBody() *ConfirmPhysicalConnectionResponseBody {
	return s.Body
}

func (s *ConfirmPhysicalConnectionResponse) SetHeaders(v map[string]*string) *ConfirmPhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *ConfirmPhysicalConnectionResponse) SetStatusCode(v int32) *ConfirmPhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfirmPhysicalConnectionResponse) SetBody(v *ConfirmPhysicalConnectionResponseBody) *ConfirmPhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *ConfirmPhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
