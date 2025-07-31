// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceKernelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody
	GetRequestId() *string
}

type UpgradeDBInstanceKernelVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 27B9A130-7C4B-40D9-84E8-2FC081097AAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeDBInstanceKernelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) SetRequestId(v string) *UpgradeDBInstanceKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeDBInstanceKernelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
