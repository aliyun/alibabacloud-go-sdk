// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLicense(v *GetInstanceLicenseResponseBodyLicense) *GetInstanceLicenseResponseBody
	GetLicense() *GetInstanceLicenseResponseBodyLicense
	SetRequestId(v string) *GetInstanceLicenseResponseBody
	GetRequestId() *string
}

type GetInstanceLicenseResponseBody struct {
	// The license details.
	License *GetInstanceLicenseResponseBodyLicense `json:"License,omitempty" xml:"License,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponseBody) GetLicense() *GetInstanceLicenseResponseBodyLicense {
	return s.License
}

func (s *GetInstanceLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceLicenseResponseBody) SetLicense(v *GetInstanceLicenseResponseBodyLicense) *GetInstanceLicenseResponseBody {
	s.License = v
	return s
}

func (s *GetInstanceLicenseResponseBody) SetRequestId(v string) *GetInstanceLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceLicenseResponseBody) Validate() error {
	if s.License != nil {
		if err := s.License.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceLicenseResponseBodyLicense struct {
	// The license edition.
	//
	// example:
	//
	// free
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The UNIX timestamp indicating the end of the license validity period.
	//
	// example:
	//
	// 1723996800000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The details of the instance license.
	InstanceLicenseDetail *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail `json:"InstanceLicenseDetail,omitempty" xml:"InstanceLicenseDetail,omitempty" type:"Struct"`
	// The billing method for the license.
	//
	// example:
	//
	// prepay
	LicenseChargeType *string `json:"LicenseChargeType,omitempty" xml:"LicenseChargeType,omitempty"`
	// The detailed configurations of the license, formatted as a JSON string.
	//
	// example:
	//
	// {"modules":[{"features":[{"name":"urn:alibaba:idaas:license:module:ud:customField","status":"enabled"}]……{"name":"urn:alibaba:idaas:license:tag:enterprise","status":"enabled"}],"version":"1.0"}
	LicenseConfigJson *string `json:"LicenseConfigJson,omitempty" xml:"LicenseConfigJson,omitempty"`
	// The UNIX timestamp indicating when the license was created.
	//
	// example:
	//
	// 1720509699000
	LicenseCreateTime *int64 `json:"LicenseCreateTime,omitempty" xml:"LicenseCreateTime,omitempty"`
	// The unique identifier for the license.
	//
	// example:
	//
	// license_1234xxxx
	LicenseId *string `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	// The license status.
	//
	// example:
	//
	// valid
	LicenseStatus *string `json:"LicenseStatus,omitempty" xml:"LicenseStatus,omitempty"`
	// The channel used to purchase the license.
	//
	// example:
	//
	// alibaba_cloud
	PurchaseChannel *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	// The unique identifier of the external service associated with the license.
	//
	// example:
	//
	// eiam-cn-xxxxx
	PurchaseInstanceId *string `json:"PurchaseInstanceId,omitempty" xml:"PurchaseInstanceId,omitempty"`
	// The UNIX timestamp indicating the start of the license validity period.
	//
	// example:
	//
	// 1720509699000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The user quota included with the license.
	//
	// example:
	//
	// 100
	UserQuota *int64 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetInstanceLicenseResponseBodyLicense) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponseBodyLicense) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponseBodyLicense) GetEdition() *string {
	return s.Edition
}

func (s *GetInstanceLicenseResponseBodyLicense) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetInstanceLicenseDetail() *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	return s.InstanceLicenseDetail
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseChargeType() *string {
	return s.LicenseChargeType
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseConfigJson() *string {
	return s.LicenseConfigJson
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseCreateTime() *int64 {
	return s.LicenseCreateTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseId() *string {
	return s.LicenseId
}

func (s *GetInstanceLicenseResponseBodyLicense) GetLicenseStatus() *string {
	return s.LicenseStatus
}

func (s *GetInstanceLicenseResponseBodyLicense) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *GetInstanceLicenseResponseBodyLicense) GetPurchaseInstanceId() *string {
	return s.PurchaseInstanceId
}

func (s *GetInstanceLicenseResponseBodyLicense) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetInstanceLicenseResponseBodyLicense) GetUserQuota() *int64 {
	return s.UserQuota
}

func (s *GetInstanceLicenseResponseBodyLicense) SetEdition(v string) *GetInstanceLicenseResponseBodyLicense {
	s.Edition = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetEndTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.EndTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetInstanceLicenseDetail(v *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) *GetInstanceLicenseResponseBodyLicense {
	s.InstanceLicenseDetail = v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseChargeType(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseChargeType = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseConfigJson(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseConfigJson = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseCreateTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseCreateTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseId(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseId = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetLicenseStatus(v string) *GetInstanceLicenseResponseBodyLicense {
	s.LicenseStatus = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetPurchaseChannel(v string) *GetInstanceLicenseResponseBodyLicense {
	s.PurchaseChannel = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetPurchaseInstanceId(v string) *GetInstanceLicenseResponseBodyLicense {
	s.PurchaseInstanceId = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetStartTime(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.StartTime = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) SetUserQuota(v int64) *GetInstanceLicenseResponseBodyLicense {
	s.UserQuota = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicense) Validate() error {
	if s.InstanceLicenseDetail != nil {
		if err := s.InstanceLicenseDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail struct {
	// The status of the license for the Conditional Access feature.
	//
	// example:
	//
	// enabled
	ConditionalAccessPolicyLicenseStatus *string `json:"ConditionalAccessPolicyLicenseStatus,omitempty" xml:"ConditionalAccessPolicyLicenseStatus,omitempty"`
	// The status of the license for machine-to-machine (M2M) applications.
	//
	// example:
	//
	// enabled
	M2mApplicationLicenseStatus *string `json:"M2mApplicationLicenseStatus,omitempty" xml:"M2mApplicationLicenseStatus,omitempty"`
	// The quota for machine-to-machine (M2M) applications.
	//
	// example:
	//
	// 2
	M2mApplicationQuota *int64 `json:"M2mApplicationQuota,omitempty" xml:"M2mApplicationQuota,omitempty"`
	// The machine identity status.
	//
	// example:
	//
	// enabled
	MimApplicationLicenseStatus *string `json:"MimApplicationLicenseStatus,omitempty" xml:"MimApplicationLicenseStatus,omitempty"`
	// The quota for network access endpoints.
	//
	// example:
	//
	// 1
	NetworkAccessEndpointQuota *int64 `json:"NetworkAccessEndpointQuota,omitempty" xml:"NetworkAccessEndpointQuota,omitempty"`
	// The number of active subscription accounts.
	//
	// example:
	//
	// 100
	PrepaidActiveUserNumber *int64 `json:"PrepaidActiveUserNumber,omitempty" xml:"PrepaidActiveUserNumber,omitempty"`
	// The user quota included with the license.
	//
	// example:
	//
	// 100
	UserQuota *int64 `json:"UserQuota,omitempty" xml:"UserQuota,omitempty"`
}

func (s GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GoString() string {
	return s.String()
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetConditionalAccessPolicyLicenseStatus() *string {
	return s.ConditionalAccessPolicyLicenseStatus
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetM2mApplicationLicenseStatus() *string {
	return s.M2mApplicationLicenseStatus
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetM2mApplicationQuota() *int64 {
	return s.M2mApplicationQuota
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetMimApplicationLicenseStatus() *string {
	return s.MimApplicationLicenseStatus
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetNetworkAccessEndpointQuota() *int64 {
	return s.NetworkAccessEndpointQuota
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetPrepaidActiveUserNumber() *int64 {
	return s.PrepaidActiveUserNumber
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) GetUserQuota() *int64 {
	return s.UserQuota
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetConditionalAccessPolicyLicenseStatus(v string) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.ConditionalAccessPolicyLicenseStatus = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetM2mApplicationLicenseStatus(v string) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.M2mApplicationLicenseStatus = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetM2mApplicationQuota(v int64) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.M2mApplicationQuota = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetMimApplicationLicenseStatus(v string) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.MimApplicationLicenseStatus = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetNetworkAccessEndpointQuota(v int64) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.NetworkAccessEndpointQuota = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetPrepaidActiveUserNumber(v int64) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.PrepaidActiveUserNumber = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) SetUserQuota(v int64) *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail {
	s.UserQuota = &v
	return s
}

func (s *GetInstanceLicenseResponseBodyLicenseInstanceLicenseDetail) Validate() error {
	return dara.Validate(s)
}
