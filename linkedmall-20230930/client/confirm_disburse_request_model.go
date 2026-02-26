// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDisburseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ConfirmDisburseCmd) *ConfirmDisburseRequest
	GetBody() *ConfirmDisburseCmd
}

type ConfirmDisburseRequest struct {
	// This parameter is required.
	Body *ConfirmDisburseCmd `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfirmDisburseRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDisburseRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDisburseRequest) GetBody() *ConfirmDisburseCmd {
	return s.Body
}

func (s *ConfirmDisburseRequest) SetBody(v *ConfirmDisburseCmd) *ConfirmDisburseRequest {
	s.Body = v
	return s
}

func (s *ConfirmDisburseRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
