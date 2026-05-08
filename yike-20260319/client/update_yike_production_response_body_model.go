// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateYikeProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *UpdateYikeProductionResponseBody
	GetProductionId() *string
	SetRequestId(v string) *UpdateYikeProductionResponseBody
	GetRequestId() *string
}

type UpdateYikeProductionResponseBody struct {
	// example:
	//
	// pd_12334**
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateYikeProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateYikeProductionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateYikeProductionResponseBody) GetProductionId() *string {
	return s.ProductionId
}

func (s *UpdateYikeProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateYikeProductionResponseBody) SetProductionId(v string) *UpdateYikeProductionResponseBody {
	s.ProductionId = &v
	return s
}

func (s *UpdateYikeProductionResponseBody) SetRequestId(v string) *UpdateYikeProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateYikeProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
