// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProvisionedProductPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProvisionedProductPlanResponseBody
	GetRequestId() *string
}

type DeleteProvisionedProductPlanResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F81BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteProvisionedProductPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProvisionedProductPlanResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProvisionedProductPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProvisionedProductPlanResponseBody) SetRequestId(v string) *DeleteProvisionedProductPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProvisionedProductPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
