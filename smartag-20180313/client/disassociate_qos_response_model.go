// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateQosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisassociateQosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisassociateQosResponse
	GetStatusCode() *int32
	SetBody(v *DisassociateQosResponseBody) *DisassociateQosResponse
	GetBody() *DisassociateQosResponseBody
}

type DisassociateQosResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisassociateQosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisassociateQosResponse) String() string {
	return dara.Prettify(s)
}

func (s DisassociateQosResponse) GoString() string {
	return s.String()
}

func (s *DisassociateQosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisassociateQosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisassociateQosResponse) GetBody() *DisassociateQosResponseBody {
	return s.Body
}

func (s *DisassociateQosResponse) SetHeaders(v map[string]*string) *DisassociateQosResponse {
	s.Headers = v
	return s
}

func (s *DisassociateQosResponse) SetStatusCode(v int32) *DisassociateQosResponse {
	s.StatusCode = &v
	return s
}

func (s *DisassociateQosResponse) SetBody(v *DisassociateQosResponseBody) *DisassociateQosResponse {
	s.Body = v
	return s
}

func (s *DisassociateQosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
