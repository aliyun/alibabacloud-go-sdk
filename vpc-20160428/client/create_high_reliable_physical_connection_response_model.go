// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHighReliablePhysicalConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHighReliablePhysicalConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHighReliablePhysicalConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CreateHighReliablePhysicalConnectionResponseBody) *CreateHighReliablePhysicalConnectionResponse
	GetBody() *CreateHighReliablePhysicalConnectionResponseBody
}

type CreateHighReliablePhysicalConnectionResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHighReliablePhysicalConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHighReliablePhysicalConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHighReliablePhysicalConnectionResponse) GoString() string {
	return s.String()
}

func (s *CreateHighReliablePhysicalConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHighReliablePhysicalConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHighReliablePhysicalConnectionResponse) GetBody() *CreateHighReliablePhysicalConnectionResponseBody {
	return s.Body
}

func (s *CreateHighReliablePhysicalConnectionResponse) SetHeaders(v map[string]*string) *CreateHighReliablePhysicalConnectionResponse {
	s.Headers = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponse) SetStatusCode(v int32) *CreateHighReliablePhysicalConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponse) SetBody(v *CreateHighReliablePhysicalConnectionResponseBody) *CreateHighReliablePhysicalConnectionResponse {
	s.Body = v
	return s
}

func (s *CreateHighReliablePhysicalConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
