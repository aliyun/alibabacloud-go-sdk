// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeKernelVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeKernelVersionResponseBody
	GetRequestId() *string
}

type UpgradeKernelVersionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A444FFFFFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpgradeKernelVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeKernelVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeKernelVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeKernelVersionResponseBody) SetRequestId(v string) *UpgradeKernelVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeKernelVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
