// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateYikeProductionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetProductionId(v string) *CreateYikeProductionResponseBody
	GetProductionId() *string
	SetRequestId(v string) *CreateYikeProductionResponseBody
	GetRequestId() *string
}

type CreateYikeProductionResponseBody struct {
	// example:
	//
	// ProductionId
	ProductionId *string `json:"ProductionId,omitempty" xml:"ProductionId,omitempty"`
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateYikeProductionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateYikeProductionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateYikeProductionResponseBody) GetProductionId() *string {
	return s.ProductionId
}

func (s *CreateYikeProductionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateYikeProductionResponseBody) SetProductionId(v string) *CreateYikeProductionResponseBody {
	s.ProductionId = &v
	return s
}

func (s *CreateYikeProductionResponseBody) SetRequestId(v string) *CreateYikeProductionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateYikeProductionResponseBody) Validate() error {
	return dara.Validate(s)
}
