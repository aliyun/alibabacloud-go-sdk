// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAccelerateTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEnterpriseAccelerateTargetResponseBody
	GetRequestId() *string
}

type CreateEnterpriseAccelerateTargetResponseBody struct {
	// example:
	//
	// D1AE33DD-0D46-59CD-8340-92BEA2BDD0F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnterpriseAccelerateTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAccelerateTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAccelerateTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnterpriseAccelerateTargetResponseBody) SetRequestId(v string) *CreateEnterpriseAccelerateTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnterpriseAccelerateTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
