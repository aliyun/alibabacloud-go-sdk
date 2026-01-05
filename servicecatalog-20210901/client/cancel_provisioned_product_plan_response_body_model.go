// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelProvisionedProductPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelProvisionedProductPlanResponseBody
	GetRequestId() *string
}

type CancelProvisionedProductPlanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelProvisionedProductPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CancelProvisionedProductPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelProvisionedProductPlanResponseBody) SetRequestId(v string) *CancelProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelProvisionedProductPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
