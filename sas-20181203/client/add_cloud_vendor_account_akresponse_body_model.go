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
	// The AK information that is added.
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
	// The AK type. Valid values:
	//
	// - **primary**: Primary account.
	//
	// - **sub**: Sub-account.
	//
	// example:
	//
	// sub
	AkType *string `json:"AkType,omitempty" xml:"AkType,omitempty"`
	// The unique ID of the AK.
	//
	// example:
	//
	// 2158
	AuthId *int64 `json:"AuthId,omitempty" xml:"AuthId,omitempty"`
	// The list of AK-associated modules.
	AuthModules []*AddCloudVendorAccountAKResponseBodyDataAuthModules `json:"AuthModules,omitempty" xml:"AuthModules,omitempty" type:"Repeated"`
	// The account ID.
	//
	// > The account ID of the connected cloud vendor.
	//
	// example:
	//
	// azure_demo_1
	CtdrCloudUserId *string `json:"CtdrCloudUserId,omitempty" xml:"CtdrCloudUserId,omitempty"`
	// The AK exception information.
	//
	// example:
	//
	// The IAM user is forbidden in the currently selected region
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The AK parameter ID.
	//
	// example:
	//
	// AE6SLd****
	SecretId *string `json:"SecretId,omitempty" xml:"SecretId,omitempty"`
	// The AK usage status. Valid values:
	//
	// - **0**: In use.
	//
	// - **1**: Abnormal.
	//
	// - **2**: Validity verification in progress.
	//
	// - **3**: Validity verification timed out.
	//
	// example:
	//
	// 0
	ServiceStatus *int32 `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	// The AK status. Valid values:
	//
	// - **0**: Enabled.
	//
	// - **1**: Not enabled.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The cloud asset vendor. Valid values:
	//
	// - **Tencent**: Tencent Cloud
	//
	// - **HUAWEICLOUD**: Huawei Cloud
	//
	// - **Azure**: Azure
	//
	// - **AWS**: AWS
	//
	// - **VOLCENGINE**: Volcengine
	//
	// - **google**: Google Cloud
	//
	// - **CHAITIN**: Chaitin Technology
	//
	// - **FORTINET**: Fortinet
	//
	// - **THREATBOOK**: ThreatBook
	//
	// example:
	//
	// Tencent
	Vendor *string `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
	// The AK account name.
	//
	// >Used to identify the account to which third-party host assets belong.
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
	// The module exception information.
	//
	// example:
	//
	// ak_domain_error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The module code. Valid values:
	//
	// - **HOST**: Host
	//
	// - **CSPM**: Cloud product configuration check
	//
	// - **SIEM**: CloudSiem
	//
	// - **TRIAL**: Log audit
	//
	// example:
	//
	// HOST
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// The cloud asset description associated with the module.
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
	// The module status. Valid values:
	//
	// - **0**: In use.
	//
	// - **1**: Abnormal.
	//
	// - **2**: Validity verification in progress.
	//
	// - **3**: Validity verification timed out.
	//
	// example:
	//
	// 0
	ModuleServiceStatus *int32 `json:"ModuleServiceStatus,omitempty" xml:"ModuleServiceStatus,omitempty"`
	// The associate permission description for the module.
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
