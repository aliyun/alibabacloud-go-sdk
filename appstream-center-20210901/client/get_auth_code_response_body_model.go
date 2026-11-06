// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAuthCodeResponseBody
	GetCode() *string
	SetData(v *GetAuthCodeResponseBodyData) *GetAuthCodeResponseBody
	GetData() *GetAuthCodeResponseBodyData
	SetMessage(v string) *GetAuthCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAuthCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAuthCodeResponseBody
	GetSuccess() *bool
}

type GetAuthCodeResponseBody struct {
	// The error code. It is returned when the request fails. A success response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned by the operation.
	Data *GetAuthCodeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error or result message. It describes the cause when the request fails. A success response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was processed successfully. A value of true indicates that the request was processed successfully. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAuthCodeResponseBody) GetData() *GetAuthCodeResponseBodyData {
	return s.Data
}

func (s *GetAuthCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAuthCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAuthCodeResponseBody) SetCode(v string) *GetAuthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *GetAuthCodeResponseBody) SetData(v *GetAuthCodeResponseBodyData) *GetAuthCodeResponseBody {
	s.Data = v
	return s
}

func (s *GetAuthCodeResponseBody) SetMessage(v string) *GetAuthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *GetAuthCodeResponseBody) SetRequestId(v string) *GetAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAuthCodeResponseBody) SetSuccess(v bool) *GetAuthCodeResponseBody {
	s.Success = &v
	return s
}

func (s *GetAuthCodeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAuthCodeResponseBodyData struct {
	// The network access type supported by the connection resource.
	//
	// - **INTERNET**: Supports Internet access.
	//
	// - **VPC**: Supports virtual private cloud (VPC) access.
	//
	// - **ANY**: Supports both Internet and VPC access.
	//
	// example:
	//
	// INTERNET
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// The application ID, which is used for subsequent application queries or unbinding the image.
	//
	// example:
	//
	// app-bp1example****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The delivery group ID of the creation or validation resource. It is used together with the user and application information in this response during connection.
	//
	// example:
	//
	// aig-bp1example****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The application version ID, which is used to locate a specific version of the application. This field is not the image version ID.
	//
	// example:
	//
	// apv-bp1example****
	AppVersionId *string `json:"AppVersionId,omitempty" xml:"AppVersionId,omitempty"`
	// The authorization code used to connect to the creation or validation resource. **The validity period is 180 seconds (3 minutes) from the time of issuance. Use it promptly after obtaining it.*	- Use it together with the user, delivery group, and application information in the same response, and keep it confidential as a connection credential.
	//
	// Do not reuse the authorization code after it is redeemed. If it has been used or expired, call this operation again to obtain a new one.
	//
	// example:
	//
	// e4e169bea1cc48e8afac53**********
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// The purpose type of the connection resource.
	//
	// - **ImageBuilder**: Image or application creation resource.
	//
	// - **ImageValidate**: Image or application validation resource.
	//
	// example:
	//
	// ImageBuilder
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The region ID in this response.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID used when connecting to the creation or validation environment. Use the value returned in this response.
	//
	// example:
	//
	// __vhd_builder__
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetAuthCodeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAuthCodeResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAuthCodeResponseBodyData) GetAccessType() *string {
	return s.AccessType
}

func (s *GetAuthCodeResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *GetAuthCodeResponseBodyData) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *GetAuthCodeResponseBodyData) GetAppVersionId() *string {
	return s.AppVersionId
}

func (s *GetAuthCodeResponseBodyData) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetAuthCodeResponseBodyData) GetProductType() *string {
	return s.ProductType
}

func (s *GetAuthCodeResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetAuthCodeResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *GetAuthCodeResponseBodyData) SetAccessType(v string) *GetAuthCodeResponseBodyData {
	s.AccessType = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetAppId(v string) *GetAuthCodeResponseBodyData {
	s.AppId = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetAppInstanceGroupId(v string) *GetAuthCodeResponseBodyData {
	s.AppInstanceGroupId = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetAppVersionId(v string) *GetAuthCodeResponseBodyData {
	s.AppVersionId = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetAuthCode(v string) *GetAuthCodeResponseBodyData {
	s.AuthCode = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetProductType(v string) *GetAuthCodeResponseBodyData {
	s.ProductType = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetRegionId(v string) *GetAuthCodeResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) SetUserId(v string) *GetAuthCodeResponseBodyData {
	s.UserId = &v
	return s
}

func (s *GetAuthCodeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
