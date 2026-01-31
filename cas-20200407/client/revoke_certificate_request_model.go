// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *RevokeCertificateRequest
	GetInstanceId() *string
}

type RevokeCertificateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cas-cn-68n1mm16****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RevokeCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeCertificateRequest) GoString() string {
	return s.String()
}

func (s *RevokeCertificateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RevokeCertificateRequest) SetInstanceId(v string) *RevokeCertificateRequest {
	s.InstanceId = &v
	return s
}

func (s *RevokeCertificateRequest) Validate() error {
	return dara.Validate(s)
}
