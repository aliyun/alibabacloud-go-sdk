// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *Quota) *UpdateQuotaRequest
	GetBody() *Quota
}

type UpdateQuotaRequest struct {
	Body *Quota `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) GetBody() *Quota {
	return s.Body
}

func (s *UpdateQuotaRequest) SetBody(v *Quota) *UpdateQuotaRequest {
	s.Body = v
	return s
}

func (s *UpdateQuotaRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
