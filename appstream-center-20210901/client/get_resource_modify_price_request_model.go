// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceModifyPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v int64) *GetResourceModifyPriceRequest
	GetAmount() *int64
	SetAppInstanceGroupId(v string) *GetResourceModifyPriceRequest
	GetAppInstanceGroupId() *string
	SetProductType(v string) *GetResourceModifyPriceRequest
	GetProductType() *string
}

type GetResourceModifyPriceRequest struct {
	// The total number of subscription resource nodes after the scale-out. Unit: nodes.
	//
	// **The value must be greater than the current number of purchased nodes in the delivery group and cannot exceed the node quota of the account.*	- This parameter is not the number of added nodes or the number of sessions.
	//
	// For example, if there is currently 1 node and you want to scale out to 2 nodes, specify 2 to indicate a final total of 2 nodes.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The ID of the delivery group to be scaled out. You can query only one delivery group at a time.
	//
	// The delivery group must belong to the current account and match ProductType. It must be billed by node on a subscription basis, have a normal order status, and be associated with a resource node pool.
	//
	// You can obtain the cloud app delivery group ID by calling [ListAppInstanceGroup](https://api.aliyun.com/document/appstream-center/2021-09-01/ListAppInstanceGroup). You can call [GetAppInstanceGroup](https://api.aliyun.com/document/appstream-center/2021-09-01/GetAppInstanceGroup) to further verify the current configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The product type, which must match the product to which the target delivery group belongs. Valid values:
	//
	// - CloudApp: WUYING Cloud App. Queries the resource scale-out price of a cloud app delivery group.
	//
	// - CloudBrowser: Cloud Browser. Applies only to subscription browser delivery groups that are billed by node.
	//
	// This operation does not apply to pay-as-you-go or browser groups billed by monthly active user (MAU).
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s GetResourceModifyPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceModifyPriceRequest) GoString() string {
	return s.String()
}

func (s *GetResourceModifyPriceRequest) GetAmount() *int64 {
	return s.Amount
}

func (s *GetResourceModifyPriceRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetResourceModifyPriceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *GetResourceModifyPriceRequest) SetAmount(v int64) *GetResourceModifyPriceRequest {
	s.Amount = &v
	return s
}

func (s *GetResourceModifyPriceRequest) SetAppInstanceGroupId(v string) *GetResourceModifyPriceRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetResourceModifyPriceRequest) SetProductType(v string) *GetResourceModifyPriceRequest {
	s.ProductType = &v
	return s
}

func (s *GetResourceModifyPriceRequest) Validate() error {
	return dara.Validate(s)
}
