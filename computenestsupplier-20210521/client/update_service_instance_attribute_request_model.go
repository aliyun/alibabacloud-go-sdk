// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *UpdateServiceInstanceAttributeRequest
	GetEndTime() *string
	SetLicenseData(v *UpdateServiceInstanceAttributeRequestLicenseData) *UpdateServiceInstanceAttributeRequest
	GetLicenseData() *UpdateServiceInstanceAttributeRequestLicenseData
	SetReason(v string) *UpdateServiceInstanceAttributeRequest
	GetReason() *string
	SetRegionId(v string) *UpdateServiceInstanceAttributeRequest
	GetRegionId() *string
	SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeRequest
	GetServiceInstanceId() *string
}

type UpdateServiceInstanceAttributeRequest struct {
	// The time when the service instance expires.
	//
	// Use the UTC time format: yyyy-MM-ddTHH:mmZ
	//
	// example:
	//
	// 2023-12-25T02:28:40Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The License Data
	LicenseData *UpdateServiceInstanceAttributeRequestLicenseData `json:"LicenseData,omitempty" xml:"LicenseData,omitempty" type:"Struct"`
	// Application reason, currently used for trial application extension.
	//
	// example:
	//
	// \\"\\"
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// si-3df88e962cdexxxxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
}

func (s UpdateServiceInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateServiceInstanceAttributeRequest) GetLicenseData() *UpdateServiceInstanceAttributeRequestLicenseData {
	return s.LicenseData
}

func (s *UpdateServiceInstanceAttributeRequest) GetReason() *string {
	return s.Reason
}

func (s *UpdateServiceInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateServiceInstanceAttributeRequest) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpdateServiceInstanceAttributeRequest) SetEndTime(v string) *UpdateServiceInstanceAttributeRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetLicenseData(v *UpdateServiceInstanceAttributeRequestLicenseData) *UpdateServiceInstanceAttributeRequest {
	s.LicenseData = v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetReason(v string) *UpdateServiceInstanceAttributeRequest {
	s.Reason = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetRegionId(v string) *UpdateServiceInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) SetServiceInstanceId(v string) *UpdateServiceInstanceAttributeRequest {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceInstanceAttributeRequestLicenseData struct {
	// The Custom Data
	//
	// example:
	//
	// {"Test"}
	CustomData *string `json:"CustomData,omitempty" xml:"CustomData,omitempty"`
	// Mock response info.
	ResponseInfo *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo `json:"ResponseInfo,omitempty" xml:"ResponseInfo,omitempty" type:"Struct"`
}

func (s UpdateServiceInstanceAttributeRequestLicenseData) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequestLicenseData) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) GetCustomData() *string {
	return s.CustomData
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) GetResponseInfo() *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	return s.ResponseInfo
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) SetCustomData(v string) *UpdateServiceInstanceAttributeRequestLicenseData {
	s.CustomData = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) SetResponseInfo(v *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) *UpdateServiceInstanceAttributeRequestLicenseData {
	s.ResponseInfo = v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseData) Validate() error {
	return dara.Validate(s)
}

type UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo struct {
	// Mock error code.
	//
	// example:
	//
	// EntityNotExist.Service
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Mock error message.
	//
	// example:
	//
	// The provided parameter "ServiceId" is invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// if you want mock response, please open this option.
	//
	// example:
	//
	// true
	UpdateResponse *bool `json:"UpdateResponse,omitempty" xml:"UpdateResponse,omitempty"`
}

func (s UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) GoString() string {
	return s.String()
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) GetUpdateResponse() *bool {
	return s.UpdateResponse
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetErrorCode(v string) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.ErrorCode = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetErrorMessage(v string) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) SetUpdateResponse(v bool) *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo {
	s.UpdateResponse = &v
	return s
}

func (s *UpdateServiceInstanceAttributeRequestLicenseDataResponseInfo) Validate() error {
	return dara.Validate(s)
}
