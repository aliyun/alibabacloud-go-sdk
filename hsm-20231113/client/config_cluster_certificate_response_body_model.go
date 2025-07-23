// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigClusterCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigClusterCertificateResponseBody
	GetRequestId() *string
}

type ConfigClusterCertificateResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigClusterCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigClusterCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigClusterCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigClusterCertificateResponseBody) SetRequestId(v string) *ConfigClusterCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigClusterCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
