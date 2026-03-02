// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferEnterpriseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *EnterpriseTransferCmd) *TransferEnterpriseRequest
	GetBody() *EnterpriseTransferCmd
}

type TransferEnterpriseRequest struct {
	// This parameter is required.
	Body *EnterpriseTransferCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TransferEnterpriseRequest) String() string {
	return dara.Prettify(s)
}

func (s TransferEnterpriseRequest) GoString() string {
	return s.String()
}

func (s *TransferEnterpriseRequest) GetBody() *EnterpriseTransferCmd {
	return s.Body
}

func (s *TransferEnterpriseRequest) SetBody(v *EnterpriseTransferCmd) *TransferEnterpriseRequest {
	s.Body = v
	return s
}

func (s *TransferEnterpriseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
