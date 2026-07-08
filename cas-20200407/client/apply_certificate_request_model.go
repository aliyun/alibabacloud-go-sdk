// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ApplyCertificateRequest
	GetInstanceId() *string
}

type ApplyCertificateRequest struct {
	// The ID of the instance. Before you request a certificate, call the UpdateInstance operation to update and complete the application information for the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cas_dv-cn-rp643r82d0i1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ApplyCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateRequest) GoString() string {
	return s.String()
}

func (s *ApplyCertificateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ApplyCertificateRequest) SetInstanceId(v string) *ApplyCertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *ApplyCertificateRequest) Validate() error {
	return dara.Validate(s)
}
