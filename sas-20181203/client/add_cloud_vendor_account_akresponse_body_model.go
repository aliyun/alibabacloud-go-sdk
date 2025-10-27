// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCloudVendorAccountAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddCloudVendorAccountAKResponseBodyData) *AddCloudVendorAccountAKResponseBody
	GetData() *AddCloudVendorAccountAKResponseBodyData
	SetRequestId(v string) *AddCloudVendorAccountAKResponseBody
	GetRequestId() *string
}

type AddCloudVendorAccountAKResponseBody struct {
	// The information about the AccessKey pair that is added.
	Data *AddCloudVendorAccountAKResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A60DA4EC-7CD8-577D-AD73-***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddCloudVendorAccountAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKResponseBody) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKResponseBody) GetData() *AddCloudVendorAccountAKResponseBodyData {
	return s.Data
}

func (s *AddCloudVendorAccountAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddCloudVendorAccountAKResponseBody) SetData(v *AddCloudVendorAccountAKResponseBodyData) *AddCloudVendorAccountAKResponseBody {
	s.Data = v
	return s
}

func (s *AddCloudVendorAccountAKResponseBody) SetRequestId(v string) *AddCloudVendorAccountAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddCloudVendorAccountAKResponseBodyData struct {
	// The type of the account to which the AccessKey pair belongs. Valid values:
	//
	// 	- **primary**: a primary account
	//
	// 	- **sub**: a sub-account
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
	AuthModules     []*AddCloudVendorAccountAKResponseBodyDataAuthModules `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	CtdrCloudUserId *string                                               `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The error message of the AccessKey pair.
	//
	// example:
	//
	// The IAM user is forbidden in the currently selected region
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The AccessKey ID.
	//
	// example:
	//
	// AE6SLd****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The service status of the AccessKey pair. Valid values:
	//
	// 	- **0**: being used
	//
	// 	- **1**: exception occurred
	//
	// 	- **2**: being validated
	//
	// 	- **3**: validation timed out
	//
	// example:
	//
	// 0
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- **0**: enabled
	//
	// 	- **1**: disabled
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud service provider. Valid values:
	//
	// 	- **Tencent**: Tencent Cloud
	//
	// 	- **HUAWEICLOUD**: Huawei Cloud
	//
	// 	- **Azure**: Microsoft Azure
	//
	// 	- **AWS**: AWS
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

func (s AddCloudVendorAccountAKResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetAkType() *string {
	return s.AkType
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetAuthId() *int64 {
	return s.AuthId
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetAuthModules() []*AddCloudVendorAccountAKResponseBodyDataAuthModules {
	return s.AuthModules
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetCtdrCloudUserId() *string {
	return s.CtdrCloudUserId
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetSecretId() *string {
	return s.SecretId
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetServiceStatus() *int32 {
	return s.ServiceStatus
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetVendor() *string {
	return s.Vendor
}

func (s *AddCloudVendorAccountAKResponseBodyData) GetVendorAuthAlias() *string {
	return s.VendorAuthAlias
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetAkType(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.AkType = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetAuthId(v int64) *AddCloudVendorAccountAKResponseBodyData {
	s.AuthId = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetAuthModules(v []*AddCloudVendorAccountAKResponseBodyDataAuthModules) *AddCloudVendorAccountAKResponseBodyData {
	s.AuthModules = v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetCtdrCloudUserId(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.CtdrCloudUserId = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetMessage(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.Message = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetSecretId(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.SecretId = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetServiceStatus(v int32) *AddCloudVendorAccountAKResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetStatus(v int32) *AddCloudVendorAccountAKResponseBodyData {
	s.Status = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetVendor(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.Vendor = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) SetVendorAuthAlias(v string) *AddCloudVendorAccountAKResponseBodyData {
	s.VendorAuthAlias = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyData) Validate() error {
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

type AddCloudVendorAccountAKResponseBodyDataAuthModules struct {
	// The error message of the module.
	//
	// example:
	//
	// ak_domain_error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The code of the module. Valid values:
	//
	// 	- **HOST**: host
	//
	// 	- **CSPM**: configuration assessment
	//
	// 	- **SIEM**: CloudSiem
	//
	// 	- **TRIAL**: log audit
	//
	// example:
	//
	// HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The cloud asset that is associated with the module.
	//
	// example:
	//
	// Cloud server or virtual machine
	ModuleAssetType *string `json:"ModuleAssetType,omitempty" xml:"ModuleAssetType,omitempty"`
	// The display name of the module.
	//
	// example:
	//
	// Host Assets
	ModuleDisp *string `json:"ModuleDisp,omitempty" xml:"ModuleDisp,omitempty"`
	// The service status of the module. Valid values:
	//
	// 	- **0**: being used
	//
	// 	- **1**: exception occurred
	//
	// 	- **2**: being validated
	//
	// 	- **3**: validation timed out
	//
	// example:
	//
	// 0
	ModuleServiceStatus *int32 `json:"ModuleServiceStatus,omitempty" xml:"ModuleServiceStatus,omitempty"`
	// The permission description of the module.
	//
	// example:
	//
	// Read permission of the cloud server or virtual machine
	ModuleStatement *string `json:"ModuleStatement,omitempty" xml:"ModuleStatement,omitempty"`
}

func (s AddCloudVendorAccountAKResponseBodyDataAuthModules) String() string {
	return dara.Prettify(s)
}

func (s AddCloudVendorAccountAKResponseBodyDataAuthModules) GoString() string {
	return s.String()
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetMessage() *string {
	return s.Message
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetModule() *string {
	return s.Module
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleAssetType() *string {
	return s.ModuleAssetType
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleDisp() *string {
	return s.ModuleDisp
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleServiceStatus() *int32 {
	return s.ModuleServiceStatus
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) GetModuleStatement() *string {
	return s.ModuleStatement
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetMessage(v string) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.Message = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetModule(v string) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.Module = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleAssetType(v string) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleAssetType = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleDisp(v string) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleDisp = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleServiceStatus(v int32) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleServiceStatus = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) SetModuleStatement(v string) *AddCloudVendorAccountAKResponseBodyDataAuthModules {
	s.ModuleStatement = &v
	return s
}

func (s *AddCloudVendorAccountAKResponseBodyDataAuthModules) Validate() error {
	return dara.Validate(s)
}
