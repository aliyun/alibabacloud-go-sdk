// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v int64) *GetProxyResponseBody
	GetCreatorId() *int64
	SetCreatorName(v string) *GetProxyResponseBody
	GetCreatorName() *string
	SetErrorCode(v string) *GetProxyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetProxyResponseBody
	GetErrorMessage() *string
	SetHttpsPort(v int32) *GetProxyResponseBody
	GetHttpsPort() *int32
	SetInstanceId(v int64) *GetProxyResponseBody
	GetInstanceId() *int64
	SetPrivateEnable(v bool) *GetProxyResponseBody
	GetPrivateEnable() *bool
	SetPrivateHost(v string) *GetProxyResponseBody
	GetPrivateHost() *string
	SetProtocolPort(v int32) *GetProxyResponseBody
	GetProtocolPort() *int32
	SetProtocolType(v string) *GetProxyResponseBody
	GetProtocolType() *string
	SetProxyId(v int64) *GetProxyResponseBody
	GetProxyId() *int64
	SetPublicEnable(v bool) *GetProxyResponseBody
	GetPublicEnable() *bool
	SetPublicHost(v string) *GetProxyResponseBody
	GetPublicHost() *string
	SetRegionId(v string) *GetProxyResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetProxyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProxyResponseBody
	GetSuccess() *bool
}

type GetProxyResponseBody struct {
	// The ID of the user who enabled the secure access proxy feature.
	//
	// example:
	//
	// 12****
	CreatorId *int64 `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The nickname of the user who enabled the secure access proxy feature.
	//
	// example:
	//
	// test_name
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The condition cannot be empty!
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The port number used by the HTTPS protocol.
	//
	// example:
	//
	// 443
	HttpsPort *int32 `json:"HttpsPort,omitempty" xml:"HttpsPort,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// 183****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Indicates whether the internal endpoint was enabled. Default value: **true**.
	//
	// example:
	//
	// true
	PrivateEnable *bool `json:"PrivateEnable,omitempty" xml:"PrivateEnable,omitempty"`
	// The internal endpoint.
	//
	// example:
	//
	// dphzmy-5j8oimjsz6ze****.proxy.dms.aliyuncs.com
	PrivateHost *string `json:"PrivateHost,omitempty" xml:"PrivateHost,omitempty"`
	// The port number used by the protocol.
	//
	// example:
	//
	// 3306
	ProtocolPort *int32 `json:"ProtocolPort,omitempty" xml:"ProtocolPort,omitempty"`
	// The protocol type of the database. Example: MYSQL.
	//
	// example:
	//
	// MYSQL
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The ID of the secure access proxy.
	//
	// example:
	//
	// 4**
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// Indicates whether the public endpoint was enabled. Valid values:
	//
	// 	- **true**: The public endpoint was enabled.
	//
	// 	- **false**: The public endpoint was disabled.
	//
	// example:
	//
	// true
	PublicEnable *bool `json:"PublicEnable,omitempty" xml:"PublicEnable,omitempty"`
	// The public endpoint. A public endpoint is returned no matter whether the public endpoint is enabled or disabled.
	//
	// >
	//
	// 	- If the value of the PublicEnable parameter is **true**, a valid public endpoint that can be resolved by using Alibaba Cloud DNS is returned.
	//
	// 	- If the value of the PublicEnable parameter is **false**, an invalid public endpoint that cannot be resolved by using Alibaba Cloud DNS is returned.
	//
	// example:
	//
	// dphzmy-5j8oimjsz6zed7k****.proxy.dms.aliyuncs.com
	PublicHost *string `json:"PublicHost,omitempty" xml:"PublicHost,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F49D4598-2B3C-5723-865E-2CCB818E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProxyResponseBody) GoString() string {
	return s.String()
}

func (s *GetProxyResponseBody) GetCreatorId() *int64 {
	return s.CreatorId
}

func (s *GetProxyResponseBody) GetCreatorName() *string {
	return s.CreatorName
}

func (s *GetProxyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProxyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetProxyResponseBody) GetHttpsPort() *int32 {
	return s.HttpsPort
}

func (s *GetProxyResponseBody) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetProxyResponseBody) GetPrivateEnable() *bool {
	return s.PrivateEnable
}

func (s *GetProxyResponseBody) GetPrivateHost() *string {
	return s.PrivateHost
}

func (s *GetProxyResponseBody) GetProtocolPort() *int32 {
	return s.ProtocolPort
}

func (s *GetProxyResponseBody) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetProxyResponseBody) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *GetProxyResponseBody) GetPublicEnable() *bool {
	return s.PublicEnable
}

func (s *GetProxyResponseBody) GetPublicHost() *string {
	return s.PublicHost
}

func (s *GetProxyResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProxyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProxyResponseBody) SetCreatorId(v int64) *GetProxyResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetProxyResponseBody) SetCreatorName(v string) *GetProxyResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetProxyResponseBody) SetErrorCode(v string) *GetProxyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProxyResponseBody) SetErrorMessage(v string) *GetProxyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProxyResponseBody) SetHttpsPort(v int32) *GetProxyResponseBody {
	s.HttpsPort = &v
	return s
}

func (s *GetProxyResponseBody) SetInstanceId(v int64) *GetProxyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetProxyResponseBody) SetPrivateEnable(v bool) *GetProxyResponseBody {
	s.PrivateEnable = &v
	return s
}

func (s *GetProxyResponseBody) SetPrivateHost(v string) *GetProxyResponseBody {
	s.PrivateHost = &v
	return s
}

func (s *GetProxyResponseBody) SetProtocolPort(v int32) *GetProxyResponseBody {
	s.ProtocolPort = &v
	return s
}

func (s *GetProxyResponseBody) SetProtocolType(v string) *GetProxyResponseBody {
	s.ProtocolType = &v
	return s
}

func (s *GetProxyResponseBody) SetProxyId(v int64) *GetProxyResponseBody {
	s.ProxyId = &v
	return s
}

func (s *GetProxyResponseBody) SetPublicEnable(v bool) *GetProxyResponseBody {
	s.PublicEnable = &v
	return s
}

func (s *GetProxyResponseBody) SetPublicHost(v string) *GetProxyResponseBody {
	s.PublicHost = &v
	return s
}

func (s *GetProxyResponseBody) SetRegionId(v string) *GetProxyResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetProxyResponseBody) SetRequestId(v string) *GetProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProxyResponseBody) SetSuccess(v bool) *GetProxyResponseBody {
	s.Success = &v
	return s
}

func (s *GetProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
