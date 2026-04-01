// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateClusterCertificateRequest
	GetClusterId() *string
}

type UpdateClusterCertificateRequest struct {
	// example:
	//
	// eck-xxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s UpdateClusterCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterCertificateRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterCertificateRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateClusterCertificateRequest) SetClusterId(v string) *UpdateClusterCertificateRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateClusterCertificateRequest) Validate() error {
	return dara.Validate(s)
}
