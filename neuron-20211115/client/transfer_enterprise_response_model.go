// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferEnterpriseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TransferEnterpriseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TransferEnterpriseResponse
	GetStatusCode() *int32
	SetBody(v *Enterprise) *TransferEnterpriseResponse
	GetBody() *Enterprise
}

type TransferEnterpriseResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Enterprise        `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferEnterpriseResponse) String() string {
	return dara.Prettify(s)
}

func (s TransferEnterpriseResponse) GoString() string {
	return s.String()
}

func (s *TransferEnterpriseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TransferEnterpriseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TransferEnterpriseResponse) GetBody() *Enterprise {
	return s.Body
}

func (s *TransferEnterpriseResponse) SetHeaders(v map[string]*string) *TransferEnterpriseResponse {
	s.Headers = v
	return s
}

func (s *TransferEnterpriseResponse) SetStatusCode(v int32) *TransferEnterpriseResponse {
	s.StatusCode = &v
	return s
}

func (s *TransferEnterpriseResponse) SetBody(v *Enterprise) *TransferEnterpriseResponse {
	s.Body = v
	return s
}

func (s *TransferEnterpriseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
