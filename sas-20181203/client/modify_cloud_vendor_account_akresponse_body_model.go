// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCloudVendorAccountAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyCloudVendorAccountAKResponseBodyData) *ModifyCloudVendorAccountAKResponseBody
	GetData() *ModifyCloudVendorAccountAKResponseBodyData
	SetRequestId(v string) *ModifyCloudVendorAccountAKResponseBody
	GetRequestId() *string
}

type ModifyCloudVendorAccountAKResponseBody struct {
	// The information about the AccessKey pair that is added.
	Data *ModifyCloudVendorAccountAKResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6635CED5-4B20-5D2D-94EC-A1C8F9C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCloudVendorAccountAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorAccountAKResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorAccountAKResponseBody) GetData() *ModifyCloudVendorAccountAKResponseBodyData {
	return s.Data
}

func (s *ModifyCloudVendorAccountAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCloudVendorAccountAKResponseBody) SetData(v *ModifyCloudVendorAccountAKResponseBodyData) *ModifyCloudVendorAccountAKResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBody) SetRequestId(v string) *ModifyCloudVendorAccountAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCloudVendorAccountAKResponseBodyData struct {
	// The type of the account to which the AccessKey pair belongs. Valid values:
	//
	// 	- **primary**
	//
	// 	- **sub**
	//
	// example:
	//
	// sub
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The unique ID of the AccessKey pair.
	//
	// example:
	//
	// 2158
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The modules that are associated with the AccessKey pair.
	AuthModules     []*ModifyCloudVendorAccountAKResponseBodyDataAuthModules `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	CtdrCloudUserId *string                                                  `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The error message of the AccessKey pair.
	//
	// example:
	//
	// The IAM user is forbidden in the currently selected region
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The AccessKey ID.
	//
	// >  If AkType is set to **primary**, the value of SecretId is AccessKey ID of the third-party master account. If AkType is set to **sub**, the value of SecretId is the AccessKey ID of the third-party sub-account. This parameter value does not change for a **Microsoft Azure account**. For an Azure account, this parameter value is the **app ID*	- that is used for authentication.
	//
	// example:
	//
	// AE6SLd****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The service status of the AccessKey pair. Valid values:
	//
	// 	- **0**: being used.
	//
	// 	- **1**: exception occurred.
	//
	// 	- **2**: being validated.
	//
	// 	- **3**: validation timed out.
	//
	// example:
	//
	// 0
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- **0**: enabled.
	//
	// 	- **1**: disabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud.
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud.
	//
	// 	- **Azure**: Microsoft Azure.
	//
	// 	- **AWS**: Amazon Web Services (AWS).
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The name of the AccessKey pair.
	//
	// >  The account information of the third-party cloud servers.
	//
	// example:
	//
	// test
	VendorAuthAlias *string `json:"VendorAuthAlias,omitempty" xml:"VendorAuthAlias,omitempty"`
}

func (s ModifyCloudVendorAccountAKResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorAccountAKResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetAkType() *string {
	return s.AkType
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetAuthId() *int64 {
	return s.AuthId
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetAuthModules() []*ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	return s.AuthModules
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetSecretId() *string {
	return s.SecretId
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetVendor() *string {
	return s.Vendor
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetAkType(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.AkType = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetAuthId(v int64) *ModifyCloudVendorAccountAKResponseBodyData {
	s.AuthId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetAuthModules(v []*ModifyCloudVendorAccountAKResponseBodyDataAuthModules) *ModifyCloudVendorAccountAKResponseBodyData {
	s.AuthModules = v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetCtdrCloudUserId(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.CtdrCloudUserId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetMessage(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.Message = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetSecretId(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetServiceStatus(v int32) *ModifyCloudVendorAccountAKResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetStatus(v int32) *ModifyCloudVendorAccountAKResponseBodyData {
	s.Status = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetVendor(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.Vendor = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) SetVendorAuthAlias(v string) *ModifyCloudVendorAccountAKResponseBodyData {
	s.VendorAuthAlias = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyData) Validate() error {
	if s.AuthModules != nil {
		for _, item := range s.AuthModules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyCloudVendorAccountAKResponseBodyDataAuthModules struct {
	// The error message of the module.
	//
	// example:
	//
	// ak_domain_error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The code of the module. Valid values:
	//
	// 	- **HOST**: host.
	//
	// 	- **CSPM**: configuration assessment.
	//
	// 	- **SIEM**: CTDR.
	//
	// 	- **TRIAL**: log audit.
	//
	// example:
	//
	// HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The type of the cloud asset that is associated with the module.
	//
	// example:
	//
	// Module.
	ModuleAssetType *string `json:"ModuleAssetType,omitempty" xml:"ModuleAssetType,omitempty"`
	// The display name of the module.
	//
	// example:
	//
	// Host
	ModuleDisp *string `json:"ModuleDisp,omitempty" xml:"ModuleDisp,omitempty"`
	// The service status of the module. Valid values:
	//
	// 	- **0**: being used.
	//
	// 	- **1**: exception occurred.
	//
	// 	- **2**: being validated.
	//
	// 	- **3**: validation timed out.
	//
	// example:
	//
	// 0
	ModuleServiceStatus *int32 `json:"ModuleServiceStatus,omitempty" xml:"ModuleServiceStatus,omitempty"`
	// The permission description of the module.
	//
	// example:
	//
	// Host
	ModuleStatement *string `json:"ModuleStatement,omitempty" xml:"ModuleStatement,omitempty"`
}

func (s ModifyCloudVendorAccountAKResponseBodyDataAuthModules) String() string {
	return dara.Prettify(s)
}

func (s ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GoString() string {
	return s.String()
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetMessage() *string {
	return s.Message
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetModule() *string {
	return s.Module
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleAssetType() *string {
	return s.ModuleAssetType
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleDisp() *string {
	return s.ModuleDisp
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleServiceStatus() *int32 {
	return s.ModuleServiceStatus
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleStatement() *string {
	return s.ModuleStatement
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetMessage(v string) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.Message = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetModule(v string) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.Module = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleAssetType(v string) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleAssetType = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleDisp(v string) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleDisp = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleServiceStatus(v int32) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleServiceStatus = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleStatement(v string) *ModifyCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleStatement = &v
	return s
}

func (s *ModifyCloudVendorAccountAKResponseBodyDataAuthModules) Validate() error {
	return dara.Validate(s)
}
