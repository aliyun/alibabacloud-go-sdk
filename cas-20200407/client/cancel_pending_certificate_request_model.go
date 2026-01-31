// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelPendingCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CancelPendingCertificateRequest
	GetInstanceId() *string
}

type CancelPendingCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cas_dv-cn-rp643r82d0i1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CancelPendingCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelPendingCertificateRequest) GoString() string {
	return s.String()
}

func (s *CancelPendingCertificateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelPendingCertificateRequest) SetInstanceId(v string) *CancelPendingCertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelPendingCertificateRequest) Validate() error {
	return dara.Validate(s)
}
