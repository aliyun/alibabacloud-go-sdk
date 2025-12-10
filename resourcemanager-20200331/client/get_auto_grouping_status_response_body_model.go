// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAutoGroupingStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnableExistedResourcesTransfer(v bool) *GetAutoGroupingStatusResponseBody
	GetEnableExistedResourcesTransfer() *bool
	SetEnableStatus(v string) *GetAutoGroupingStatusResponseBody
	GetEnableStatus() *string
	SetRequestId(v string) *GetAutoGroupingStatusResponseBody
	GetRequestId() *string
}

type GetAutoGroupingStatusResponseBody struct {
	// Indicates whether the Transfer Existing Associated Resources feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	EnableExistedResourcesTransfer *bool `json:"EnableExistedResourcesTransfer,omitempty" xml:"EnableExistedResourcesTransfer,omitempty"`
	// The status of the Automatic Resource Transfer feature. Valid values:
	//
	// 	- Enabling: The feature is being enabled.
	//
	// 	- Enable: The feature is enabled.
	//
	// 	- Partial_Enable: The transfer of associated resources is enabled, but custom transfer rule-based resource transfer is disabled. You can call the [EnableAutoGrouping](https://help.aliyun.com/document_detail/2870380.html) operation to enable custom transfer rule-based resource transfer.
	//
	// 	- Disable: The feature is disabled.
	//
	// example:
	//
	// Enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0217AFEB-5318-56D4-B167-1933D83EDF3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAutoGroupingStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAutoGroupingStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetAutoGroupingStatusResponseBody) GetEnableExistedResourcesTransfer() *bool {
	return s.EnableExistedResourcesTransfer
}

func (s *GetAutoGroupingStatusResponseBody) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *GetAutoGroupingStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAutoGroupingStatusResponseBody) SetEnableExistedResourcesTransfer(v bool) *GetAutoGroupingStatusResponseBody {
	s.EnableExistedResourcesTransfer = &v
	return s
}

func (s *GetAutoGroupingStatusResponseBody) SetEnableStatus(v string) *GetAutoGroupingStatusResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *GetAutoGroupingStatusResponseBody) SetRequestId(v string) *GetAutoGroupingStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAutoGroupingStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
