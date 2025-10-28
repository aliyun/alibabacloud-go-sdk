// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEccInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryEccInfoResponseBody
	GetCode() *int32
	SetEccInfo(v *QueryEccInfoResponseBodyEccInfo) *QueryEccInfoResponseBody
	GetEccInfo() *QueryEccInfoResponseBodyEccInfo
	SetMessage(v string) *QueryEccInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryEccInfoResponseBody
	GetRequestId() *string
}

type QueryEccInfoResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the ECC.
	EccInfo *QueryEccInfoResponseBodyEccInfo `json:"EccInfo,omitempty" xml:"EccInfo,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEccInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEccInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryEccInfoResponseBody) GetEccInfo() *QueryEccInfoResponseBodyEccInfo {
	return s.EccInfo
}

func (s *QueryEccInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryEccInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEccInfoResponseBody) SetCode(v int32) *QueryEccInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryEccInfoResponseBody) SetEccInfo(v *QueryEccInfoResponseBodyEccInfo) *QueryEccInfoResponseBody {
	s.EccInfo = v
	return s
}

func (s *QueryEccInfoResponseBody) SetMessage(v string) *QueryEccInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryEccInfoResponseBody) SetRequestId(v string) *QueryEccInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEccInfoResponseBody) Validate() error {
	if s.EccInfo != nil {
		if err := s.EccInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryEccInfoResponseBodyEccInfo struct {
	// The ID of the application.
	//
	// example:
	//
	// e809****-43d7-4c6b-8e01-b0d9d1db****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// ECC ID
	//
	// example:
	//
	// a5b9****-40b4-4d7b-9c2a-55d6c1c0****
	EccId *string `json:"EccId,omitempty" xml:"EccId,omitempty"`
	// ECU ID
	//
	// example:
	//
	// 0d4e****-6d87-401f-ba81-13df9877****
	EcuId *string `json:"EcuId,omitempty" xml:"EcuId,omitempty"`
	// The ID of the ECC group.
	//
	// example:
	//
	// 57cd****-2d3b-496f-bcce-646d0a4d****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the ECC group.
	//
	// example:
	//
	// _DEFAULT_GROUP
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The MD5 hash value of the deployment package version.
	//
	// example:
	//
	// bab6****7a090e41ca9445c9b3cd****
	PackageMd5 *string `json:"PackageMd5,omitempty" xml:"PackageMd5,omitempty"`
	// The version of the deployment package.
	//
	// example:
	//
	// 20210209.153400
	PackageVersion *string `json:"PackageVersion,omitempty" xml:"PackageVersion,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-23727****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s QueryEccInfoResponseBodyEccInfo) String() string {
	return dara.Prettify(s)
}

func (s QueryEccInfoResponseBodyEccInfo) GoString() string {
	return s.String()
}

func (s *QueryEccInfoResponseBodyEccInfo) GetAppId() *string {
	return s.AppId
}

func (s *QueryEccInfoResponseBodyEccInfo) GetEccId() *string {
	return s.EccId
}

func (s *QueryEccInfoResponseBodyEccInfo) GetEcuId() *string {
	return s.EcuId
}

func (s *QueryEccInfoResponseBodyEccInfo) GetGroupId() *string {
	return s.GroupId
}

func (s *QueryEccInfoResponseBodyEccInfo) GetGroupName() *string {
	return s.GroupName
}

func (s *QueryEccInfoResponseBodyEccInfo) GetPackageMd5() *string {
	return s.PackageMd5
}

func (s *QueryEccInfoResponseBodyEccInfo) GetPackageVersion() *string {
	return s.PackageVersion
}

func (s *QueryEccInfoResponseBodyEccInfo) GetVpcId() *string {
	return s.VpcId
}

func (s *QueryEccInfoResponseBodyEccInfo) SetAppId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.AppId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetEccId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.EccId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetEcuId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.EcuId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetGroupId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.GroupId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetGroupName(v string) *QueryEccInfoResponseBodyEccInfo {
	s.GroupName = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetPackageMd5(v string) *QueryEccInfoResponseBodyEccInfo {
	s.PackageMd5 = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetPackageVersion(v string) *QueryEccInfoResponseBodyEccInfo {
	s.PackageVersion = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) SetVpcId(v string) *QueryEccInfoResponseBodyEccInfo {
	s.VpcId = &v
	return s
}

func (s *QueryEccInfoResponseBodyEccInfo) Validate() error {
	return dara.Validate(s)
}
