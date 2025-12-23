// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpcOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPurchaseStatus(v string) *CreateIpcOrderResponseBody
	GetPurchaseStatus() *string
	SetRequestId(v string) *CreateIpcOrderResponseBody
	GetRequestId() *string
}

type CreateIpcOrderResponseBody struct {
	// example:
	//
	// Success
	PurchaseStatus *string `json:"PurchaseStatus,omitempty" xml:"PurchaseStatus,omitempty"`
	// example:
	//
	// ****36-3C1E-4417-BDB2-1E034F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpcOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpcOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpcOrderResponseBody) GetPurchaseStatus() *string {
	return s.PurchaseStatus
}

func (s *CreateIpcOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpcOrderResponseBody) SetPurchaseStatus(v string) *CreateIpcOrderResponseBody {
	s.PurchaseStatus = &v
	return s
}

func (s *CreateIpcOrderResponseBody) SetRequestId(v string) *CreateIpcOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpcOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
