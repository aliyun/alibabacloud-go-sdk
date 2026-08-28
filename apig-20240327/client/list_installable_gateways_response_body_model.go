// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstallableGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListInstallableGatewaysResponseBody
	GetCode() *string
	SetData(v *ListInstallableGatewaysResponseBodyData) *ListInstallableGatewaysResponseBody
	GetData() *ListInstallableGatewaysResponseBodyData
	SetMessage(v string) *ListInstallableGatewaysResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInstallableGatewaysResponseBody
	GetRequestId() *string
}

type ListInstallableGatewaysResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListInstallableGatewaysResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FA163-3664-5D34-81D1-5FFFC94AD7D5
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListInstallableGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstallableGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstallableGatewaysResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListInstallableGatewaysResponseBody) GetData() *ListInstallableGatewaysResponseBodyData {
	return s.Data
}

func (s *ListInstallableGatewaysResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInstallableGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstallableGatewaysResponseBody) SetCode(v string) *ListInstallableGatewaysResponseBody {
	s.Code = &v
	return s
}

func (s *ListInstallableGatewaysResponseBody) SetData(v *ListInstallableGatewaysResponseBodyData) *ListInstallableGatewaysResponseBody {
	s.Data = v
	return s
}

func (s *ListInstallableGatewaysResponseBody) SetMessage(v string) *ListInstallableGatewaysResponseBody {
	s.Message = &v
	return s
}

func (s *ListInstallableGatewaysResponseBody) SetRequestId(v string) *ListInstallableGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstallableGatewaysResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstallableGatewaysResponseBodyData struct {
	Items []*ListInstallableGatewaysResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *string `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 43
	TotalSize *string `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListInstallableGatewaysResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInstallableGatewaysResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInstallableGatewaysResponseBodyData) GetItems() []*ListInstallableGatewaysResponseBodyDataItems {
	return s.Items
}

func (s *ListInstallableGatewaysResponseBodyData) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListInstallableGatewaysResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *ListInstallableGatewaysResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *ListInstallableGatewaysResponseBodyData) SetItems(v []*ListInstallableGatewaysResponseBodyDataItems) *ListInstallableGatewaysResponseBodyData {
	s.Items = v
	return s
}

func (s *ListInstallableGatewaysResponseBodyData) SetPageNumber(v string) *ListInstallableGatewaysResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyData) SetPageSize(v string) *ListInstallableGatewaysResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyData) SetTotalSize(v string) *ListInstallableGatewaysResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstallableGatewaysResponseBodyDataItems struct {
	// example:
	//
	// 2.1.10
	EngineVersion *string `json:"engineVersion,omitempty" xml:"engineVersion,omitempty"`
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// Running
	GatewayPhase *string `json:"gatewayPhase,omitempty" xml:"gatewayPhase,omitempty"`
	// example:
	//
	// true
	Installable *bool `json:"installable,omitempty" xml:"installable,omitempty"`
	// example:
	//
	// VERSION_NOT_MATCH
	InstallableFalseReasonType *string `json:"installableFalseReasonType,omitempty" xml:"installableFalseReasonType,omitempty"`
	// example:
	//
	// 1.0.0
	InstalledPluginVersion *string `json:"installedPluginVersion,omitempty" xml:"installedPluginVersion,omitempty"`
	// example:
	//
	// my-gateway
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListInstallableGatewaysResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListInstallableGatewaysResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetEngineVersion() *string {
	return s.EngineVersion
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetGatewayPhase() *string {
	return s.GatewayPhase
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetInstallable() *bool {
	return s.Installable
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetInstallableFalseReasonType() *string {
	return s.InstallableFalseReasonType
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetInstalledPluginVersion() *string {
	return s.InstalledPluginVersion
}

func (s *ListInstallableGatewaysResponseBodyDataItems) GetName() *string {
	return s.Name
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetEngineVersion(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.EngineVersion = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetGatewayId(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetGatewayPhase(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.GatewayPhase = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetInstallable(v bool) *ListInstallableGatewaysResponseBodyDataItems {
	s.Installable = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetInstallableFalseReasonType(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.InstallableFalseReasonType = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetInstalledPluginVersion(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.InstalledPluginVersion = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) SetName(v string) *ListInstallableGatewaysResponseBodyDataItems {
	s.Name = &v
	return s
}

func (s *ListInstallableGatewaysResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
