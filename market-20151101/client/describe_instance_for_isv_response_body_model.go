// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceForIsvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActiveAddress(v string) *DescribeInstanceForIsvResponseBody
	GetActiveAddress() *string
	SetAppJson(v string) *DescribeInstanceForIsvResponseBody
	GetAppJson() *string
	SetAutoRenewal(v string) *DescribeInstanceForIsvResponseBody
	GetAutoRenewal() *string
	SetBeganOn(v int64) *DescribeInstanceForIsvResponseBody
	GetBeganOn() *int64
	SetComponentJson(v string) *DescribeInstanceForIsvResponseBody
	GetComponentJson() *string
	SetCreatedOn(v int64) *DescribeInstanceForIsvResponseBody
	GetCreatedOn() *int64
	SetEndOn(v int64) *DescribeInstanceForIsvResponseBody
	GetEndOn() *int64
	SetExtendJson(v string) *DescribeInstanceForIsvResponseBody
	GetExtendJson() *string
	SetHostJson(v string) *DescribeInstanceForIsvResponseBody
	GetHostJson() *string
	SetImageJson(v string) *DescribeInstanceForIsvResponseBody
	GetImageJson() *string
	SetInstanceId(v int64) *DescribeInstanceForIsvResponseBody
	GetInstanceId() *int64
	SetIsTrial(v bool) *DescribeInstanceForIsvResponseBody
	GetIsTrial() *bool
	SetLicenseCode(v string) *DescribeInstanceForIsvResponseBody
	GetLicenseCode() *string
	SetOrderId(v int64) *DescribeInstanceForIsvResponseBody
	GetOrderId() *int64
	SetProductCode(v string) *DescribeInstanceForIsvResponseBody
	GetProductCode() *string
	SetProductName(v string) *DescribeInstanceForIsvResponseBody
	GetProductName() *string
	SetProductSkuCode(v string) *DescribeInstanceForIsvResponseBody
	GetProductSkuCode() *string
	SetProductType(v string) *DescribeInstanceForIsvResponseBody
	GetProductType() *string
	SetRelationalData(v *DescribeInstanceForIsvResponseBodyRelationalData) *DescribeInstanceForIsvResponseBody
	GetRelationalData() *DescribeInstanceForIsvResponseBodyRelationalData
	SetRequestId(v string) *DescribeInstanceForIsvResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeInstanceForIsvResponseBody
	GetStatus() *string
	SetSupplierName(v string) *DescribeInstanceForIsvResponseBody
	GetSupplierName() *string
}

type DescribeInstanceForIsvResponseBody struct {
	// The activation URL.
	//
	// example:
	//
	// https://*****.com
	ActiveAddress *string `json:"ActiveAddress,omitempty" xml:"ActiveAddress,omitempty"`
	// The application information. The metric description of the parameters in the example is as follows:
	//
	// username: the username.
	//
	// passowrd: the password.
	//
	// frontEndUrl: the frontend URL.
	//
	// adminUrl: the administration URL.
	//
	// example:
	//
	// {"frontEndUrl":"https://****.aliyundoc.com","password":"Sjtv***","adminUrl":"https://****.aliyundoc.com","username":"aliyun***"}
	AppJson *string `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	// The automatic payment information for epoch-based products. A non-empty value indicates that you enable auto-renewal.
	//
	// autoRenewCycle: Day, Month, or Year.
	//
	// autoRenewDuration: the renewal epoch.
	//
	// example:
	//
	// {"autoRenewCycle":"Year","autoRenewDuration":1}
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// The time when the instance was activated (timestamp).
	//
	// example:
	//
	// 1570634021000
	BeganOn *int64 `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	// The business module. The metric description of the parameters in the example is as follows:
	//
	// 	- package_version: the version.
	//
	// example:
	//
	// {"package_version":"yuncode000111"}
	ComponentJson *string `json:"ComponentJson,omitempty" xml:"ComponentJson,omitempty"`
	// The creation time (timestamp).
	//
	// example:
	//
	// 1570634018000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// The expiration time (timestamp).
	//
	// example:
	//
	// 1602259200000
	EndOn *int64 `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {"description": "If you have any problems, contact us..."}
	ExtendJson *string `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	// The host information. The metric description of the parameters in the example is as follows:
	//
	// username: the username.
	//
	// passowrd: the password.
	//
	// ip: the host IP address.
	//
	// innerIp: the internal network IP address of the host.
	//
	// example:
	//
	// {"password":"***","ip":"118.31.***.41","innerIp":"118.31.***.41","region":"","username":"***","beianInfo":""}
	HostJson *string `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	// The image instance information.
	//
	// imageId: the image ID.
	//
	// instanceId: the instance ID.
	//
	// ecsInstanceId: the ECS instance ID.
	//
	// example:
	//
	// {"imageId":"m-*******","instanceId":********}
	ImageJson *string `json:"ImageJson,omitempty" xml:"ImageJson,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1551111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the instance is a trial instance.
	//
	// example:
	//
	// true
	IsTrial *bool `json:"IsTrial,omitempty" xml:"IsTrial,omitempty"`
	// The activation code.
	//
	// example:
	//
	// xxxxxx
	LicenseCode *string `json:"LicenseCode,omitempty" xml:"LicenseCode,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 204211111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The commodity code.
	//
	// example:
	//
	// cmgj00**11
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product name.
	//
	// example:
	//
	// ***Collaborative Office
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// The SKU code.
	//
	// example:
	//
	// cmgj00**11-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// The product type. Valid values:
	//
	// 	- APP: application.
	//
	// 	- SERVIC: service.
	//
	// 	- MIRROR: image.
	//
	// 	- DOWNLOAD: download.
	//
	// 	- API_SERVICE: API.
	//
	// 	- DOCKER: Docker.
	//
	// example:
	//
	// APP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The relational information.
	RelationalData *DescribeInstanceForIsvResponseBodyRelationalData `json:"RelationalData,omitempty" xml:"RelationalData,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6EF60BEC-****-****-****-270359FB54A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The business status. Different values are set based on the product type.
	//
	// When ProductType is set to MIRROR:
	//
	// 	- PRODUCE: being provisioned.
	//
	// 	- USING: in use.
	//
	// 	- UNUSED: not in use.
	//
	// 	- EXPIRED: expired.
	//
	// 	- CLOSED: released.
	//
	// When ProductType is set to SERVICE:
	//
	// 	- READYING: not started.
	//
	// 	- STARTED: in service.
	//
	// 	- CONFIRM: completed.
	//
	// 	- EXPIRED: expired.
	//
	// 	- CLOSED: released.
	//
	// When ProductType is set to APP, DOWNLOAD, API_SERVICE, or DOCKER:
	//
	// 	- OPENING: being activated.
	//
	// 	- OPENED: activated.
	//
	// 	- EXPIRED: expired.
	//
	// 	- CLOSED: released.
	//
	// example:
	//
	// OPENED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The service provider name.
	//
	// example:
	//
	// ****	- Co., Ltd.
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstanceForIsvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvResponseBody) GetActiveAddress() *string {
	return s.ActiveAddress
}

func (s *DescribeInstanceForIsvResponseBody) GetAppJson() *string {
	return s.AppJson
}

func (s *DescribeInstanceForIsvResponseBody) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *DescribeInstanceForIsvResponseBody) GetBeganOn() *int64 {
	return s.BeganOn
}

func (s *DescribeInstanceForIsvResponseBody) GetComponentJson() *string {
	return s.ComponentJson
}

func (s *DescribeInstanceForIsvResponseBody) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeInstanceForIsvResponseBody) GetEndOn() *int64 {
	return s.EndOn
}

func (s *DescribeInstanceForIsvResponseBody) GetExtendJson() *string {
	return s.ExtendJson
}

func (s *DescribeInstanceForIsvResponseBody) GetHostJson() *string {
	return s.HostJson
}

func (s *DescribeInstanceForIsvResponseBody) GetImageJson() *string {
	return s.ImageJson
}

func (s *DescribeInstanceForIsvResponseBody) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeInstanceForIsvResponseBody) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *DescribeInstanceForIsvResponseBody) GetLicenseCode() *string {
	return s.LicenseCode
}

func (s *DescribeInstanceForIsvResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeInstanceForIsvResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstanceForIsvResponseBody) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstanceForIsvResponseBody) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeInstanceForIsvResponseBody) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstanceForIsvResponseBody) GetRelationalData() *DescribeInstanceForIsvResponseBodyRelationalData {
	return s.RelationalData
}

func (s *DescribeInstanceForIsvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceForIsvResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstanceForIsvResponseBody) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeInstanceForIsvResponseBody) SetActiveAddress(v string) *DescribeInstanceForIsvResponseBody {
	s.ActiveAddress = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetAppJson(v string) *DescribeInstanceForIsvResponseBody {
	s.AppJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetAutoRenewal(v string) *DescribeInstanceForIsvResponseBody {
	s.AutoRenewal = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetBeganOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetComponentJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ComponentJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetCreatedOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetEndOn(v int64) *DescribeInstanceForIsvResponseBody {
	s.EndOn = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetExtendJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetHostJson(v string) *DescribeInstanceForIsvResponseBody {
	s.HostJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetImageJson(v string) *DescribeInstanceForIsvResponseBody {
	s.ImageJson = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetInstanceId(v int64) *DescribeInstanceForIsvResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetIsTrial(v bool) *DescribeInstanceForIsvResponseBody {
	s.IsTrial = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetLicenseCode(v string) *DescribeInstanceForIsvResponseBody {
	s.LicenseCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetOrderId(v int64) *DescribeInstanceForIsvResponseBody {
	s.OrderId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductCode(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductName(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductName = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductSkuCode(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetProductType(v string) *DescribeInstanceForIsvResponseBody {
	s.ProductType = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetRelationalData(v *DescribeInstanceForIsvResponseBodyRelationalData) *DescribeInstanceForIsvResponseBody {
	s.RelationalData = v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetRequestId(v string) *DescribeInstanceForIsvResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetStatus(v string) *DescribeInstanceForIsvResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) SetSupplierName(v string) *DescribeInstanceForIsvResponseBody {
	s.SupplierName = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBody) Validate() error {
	if s.RelationalData != nil {
		if err := s.RelationalData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceForIsvResponseBodyRelationalData struct {
	// The service status. Valid values:
	//
	// 	- READYING: not started.
	//
	// 	- STARTED: in service.
	//
	// 	- CONFIRM: completed.
	//
	// 	- EXPIRED: expired.
	//
	// 	- CLOSED: released.
	//
	// example:
	//
	// STARTED
	ServiceStatus *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
}

func (s DescribeInstanceForIsvResponseBodyRelationalData) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceForIsvResponseBodyRelationalData) GoString() string {
	return s.String()
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) SetServiceStatus(v string) *DescribeInstanceForIsvResponseBodyRelationalData {
	s.ServiceStatus = &v
	return s
}

func (s *DescribeInstanceForIsvResponseBodyRelationalData) Validate() error {
	return dara.Validate(s)
}
