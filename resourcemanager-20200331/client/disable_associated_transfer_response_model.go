// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAssociatedTransferResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableAssociatedTransferResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableAssociatedTransferResponse
	GetStatusCode() *int32
	SetBody(v *DisableAssociatedTransferResponseBody) *DisableAssociatedTransferResponse
	GetBody() *DisableAssociatedTransferResponseBody
}

type DisableAssociatedTransferResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableAssociatedTransferResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableAssociatedTransferResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableAssociatedTransferResponse) GoString() string {
	return s.String()
}

func (s *DisableAssociatedTransferResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableAssociatedTransferResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableAssociatedTransferResponse) GetBody() *DisableAssociatedTransferResponseBody {
	return s.Body
}

func (s *DisableAssociatedTransferResponse) SetHeaders(v map[string]*string) *DisableAssociatedTransferResponse {
	s.Headers = v
	return s
}

func (s *DisableAssociatedTransferResponse) SetStatusCode(v int32) *DisableAssociatedTransferResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableAssociatedTransferResponse) SetBody(v *DisableAssociatedTransferResponseBody) *DisableAssociatedTransferResponse {
	s.Body = v
	return s
}

func (s *DisableAssociatedTransferResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
