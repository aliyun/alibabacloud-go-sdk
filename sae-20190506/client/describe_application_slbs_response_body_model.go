// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationSlbsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApplicationSlbsResponseBody
	GetCode() *string
	SetData(v *DescribeApplicationSlbsResponseBodyData) *DescribeApplicationSlbsResponseBody
	GetData() *DescribeApplicationSlbsResponseBodyData
	SetErrorCode(v string) *DescribeApplicationSlbsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DescribeApplicationSlbsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeApplicationSlbsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeApplicationSlbsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DescribeApplicationSlbsResponseBody
	GetTraceId() *string
}

type DescribeApplicationSlbsResponseBody struct {
	// Indicates whether the information about the SLB instances that are associated with an application was obtained successfully. Valid values:
	//
	// 	- **true**: indicates that the information was obtained successfully.
	//
	// 	- **false**: indicates that the information failed to be obtained.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *DescribeApplicationSlbsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. Valid values:
	//
	// 	- **2xx**: indicates that the request was successful.
	//
	// 	- **3xx**: indicates that the request was redirected.
	//
	// 	- **4xx**: indicates that the request was invalid.
	//
	// 	- **5xx**: indicates that a server error occurred.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The ID of the trace. It can be used to query the details of a request.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned message.
	//
	// 	- **success*	- is returned when the request succeeds.
	//
	// 	- An error code is returned when the request fails.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApplicationSlbsResponseBody) GetData() *DescribeApplicationSlbsResponseBodyData {
	return s.Data
}

func (s *DescribeApplicationSlbsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeApplicationSlbsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApplicationSlbsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationSlbsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApplicationSlbsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DescribeApplicationSlbsResponseBody) SetCode(v string) *DescribeApplicationSlbsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetData(v *DescribeApplicationSlbsResponseBodyData) *DescribeApplicationSlbsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetErrorCode(v string) *DescribeApplicationSlbsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetMessage(v string) *DescribeApplicationSlbsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetRequestId(v string) *DescribeApplicationSlbsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetSuccess(v bool) *DescribeApplicationSlbsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) SetTraceId(v string) *DescribeApplicationSlbsResponseBody {
	s.TraceId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeApplicationSlbsResponseBodyData struct {
	AppId     *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppName   *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The configurations of the Internet-facing SLB instance.
	Internet []*DescribeApplicationSlbsResponseBodyDataInternet `json:"Internet,omitempty" xml:"Internet,omitempty" type:"Repeated"`
	// The ID of the Internet-facing SLB instance.
	//
	// example:
	//
	// ``59.74.**.**``
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// PayBySpec
	InternetSlbChargeType *string `json:"InternetSlbChargeType,omitempty" xml:"InternetSlbChargeType,omitempty"`
	// example:
	//
	// false
	InternetSlbExpired *bool `json:"InternetSlbExpired,omitempty" xml:"InternetSlbExpired,omitempty"`
	// Configurations of Internet-facing SLB instances.
	//
	// example:
	//
	// lb-uf6xc7wybefehnv45****
	InternetSlbId *string `json:"InternetSlbId,omitempty" xml:"InternetSlbId,omitempty"`
	// The configurations of the internal-facing SLB instance.
	Intranet []*DescribeApplicationSlbsResponseBodyDataIntranet `json:"Intranet,omitempty" xml:"Intranet,omitempty" type:"Repeated"`
	// The error code.
	//
	// 	- The **ErrorCode*	- parameter is not returned when the request succeeds.
	//
	// 	- The **ErrorCode*	- parameter is returned when the request fails. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// 192.168.0.0
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// PayBySpec
	IntranetSlbChargeType *string `json:"IntranetSlbChargeType,omitempty" xml:"IntranetSlbChargeType,omitempty"`
	// example:
	//
	// false
	IntranetSlbExpired *bool `json:"IntranetSlbExpired,omitempty" xml:"IntranetSlbExpired,omitempty"`
	// The IP address of the internal-facing SLB instance.
	//
	// example:
	//
	// lb-uf6xc7wybefehnv45****
	IntranetSlbId *string `json:"IntranetSlbId,omitempty" xml:"IntranetSlbId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyData) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationSlbsResponseBodyData) GetAppName() *string {
	return s.AppName
}

func (s *DescribeApplicationSlbsResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeApplicationSlbsResponseBodyData) GetInternet() []*DescribeApplicationSlbsResponseBodyDataInternet {
	return s.Internet
}

func (s *DescribeApplicationSlbsResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeApplicationSlbsResponseBodyData) GetInternetSlbChargeType() *string {
	return s.InternetSlbChargeType
}

func (s *DescribeApplicationSlbsResponseBodyData) GetInternetSlbExpired() *bool {
	return s.InternetSlbExpired
}

func (s *DescribeApplicationSlbsResponseBodyData) GetInternetSlbId() *string {
	return s.InternetSlbId
}

func (s *DescribeApplicationSlbsResponseBodyData) GetIntranet() []*DescribeApplicationSlbsResponseBodyDataIntranet {
	return s.Intranet
}

func (s *DescribeApplicationSlbsResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeApplicationSlbsResponseBodyData) GetIntranetSlbChargeType() *string {
	return s.IntranetSlbChargeType
}

func (s *DescribeApplicationSlbsResponseBodyData) GetIntranetSlbExpired() *bool {
	return s.IntranetSlbExpired
}

func (s *DescribeApplicationSlbsResponseBodyData) GetIntranetSlbId() *string {
	return s.IntranetSlbId
}

func (s *DescribeApplicationSlbsResponseBodyData) SetAppId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetAppName(v string) *DescribeApplicationSlbsResponseBodyData {
	s.AppName = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetClusterId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternet(v []*DescribeApplicationSlbsResponseBodyDataInternet) *DescribeApplicationSlbsResponseBodyData {
	s.Internet = v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetIp(v string) *DescribeApplicationSlbsResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetSlbChargeType(v string) *DescribeApplicationSlbsResponseBodyData {
	s.InternetSlbChargeType = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetSlbExpired(v bool) *DescribeApplicationSlbsResponseBodyData {
	s.InternetSlbExpired = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetInternetSlbId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.InternetSlbId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranet(v []*DescribeApplicationSlbsResponseBodyDataIntranet) *DescribeApplicationSlbsResponseBodyData {
	s.Intranet = v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetIp(v string) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetSlbChargeType(v string) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetSlbChargeType = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetSlbExpired(v bool) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetSlbExpired = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) SetIntranetSlbId(v string) *DescribeApplicationSlbsResponseBodyData {
	s.IntranetSlbId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyData) Validate() error {
	if s.Internet != nil {
		for _, item := range s.Internet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Intranet != nil {
		for _, item := range s.Intranet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationSlbsResponseBodyDataInternet struct {
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// example:
	//
	// wwe
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// example:
	//
	// 56
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The timestamp when the canary release rule was created.
	//
	// example:
	//
	// 1741247308294
	CreateTime            *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableConnectionDrain *bool  `json:"EnableConnectionDrain,omitempty" xml:"EnableConnectionDrain,omitempty"`
	// example:
	//
	// 1513561019707729_16f37aae5f3_-375882821_-169099****
	HttpsCaCertId *string `json:"HttpsCaCertId,omitempty" xml:"HttpsCaCertId,omitempty"`
	// The supported protocol.
	//
	// example:
	//
	// 1513561019707729_16f37aae5f3_-375882821_-169099****
	HttpsCertId *string `json:"HttpsCertId,omitempty" xml:"HttpsCertId,omitempty"`
	// The ID of the internal-facing SLB instance.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The container port.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// false
	StickySession *bool `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The port specified for the SLB listener.
	//
	// example:
	//
	// 8080
	TargetPort     *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyDataInternet) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyDataInternet) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetEnableConnectionDrain() *bool {
	return s.EnableConnectionDrain
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetHttpsCaCertId() *string {
	return s.HttpsCaCertId
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetHttpsCertId() *string {
	return s.HttpsCertId
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetPort() *int32 {
	return s.Port
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetStickySession() *bool {
	return s.StickySession
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetConnectionDrainTimeout(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetCookie(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.Cookie = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetCookieTimeout(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetCreateTime(v int64) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetEnableConnectionDrain(v bool) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.EnableConnectionDrain = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetHttpsCaCertId(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.HttpsCaCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetHttpsCertId(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.HttpsCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetPort(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.Port = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetProtocol(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.Protocol = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetStickySession(v bool) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.StickySession = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetStickySessionType(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.StickySessionType = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetTargetPort(v int32) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.TargetPort = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) SetVServerGroupId(v string) *DescribeApplicationSlbsResponseBodyDataInternet {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataInternet) Validate() error {
	return dara.Validate(s)
}

type DescribeApplicationSlbsResponseBodyDataIntranet struct {
	ConnectionDrainTimeout *int32 `json:"ConnectionDrainTimeout,omitempty" xml:"ConnectionDrainTimeout,omitempty"`
	// example:
	//
	// wwe
	Cookie *string `json:"Cookie,omitempty" xml:"Cookie,omitempty"`
	// example:
	//
	// 56
	CookieTimeout *int32 `json:"CookieTimeout,omitempty" xml:"CookieTimeout,omitempty"`
	// The timestamp when the canary release rule was created.
	//
	// example:
	//
	// 1741247308294
	CreateTime            *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	EnableConnectionDrain *bool  `json:"EnableConnectionDrain,omitempty" xml:"EnableConnectionDrain,omitempty"`
	// example:
	//
	// 1513561019707729_16f37aae5f3_-375882821_-169099****
	HttpsCaCertId *string `json:"HttpsCaCertId,omitempty" xml:"HttpsCaCertId,omitempty"`
	// The supported protocol.
	//
	// example:
	//
	// 1513561019707729_16f37aae5f3_-375882821_-169099****
	HttpsCertId *string `json:"HttpsCertId,omitempty" xml:"HttpsCertId,omitempty"`
	// The IP address of the Internet-facing SLB instance.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The container port.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// example:
	//
	// false
	StickySession *bool `json:"StickySession,omitempty" xml:"StickySession,omitempty"`
	// example:
	//
	// insert
	StickySessionType *string `json:"StickySessionType,omitempty" xml:"StickySessionType,omitempty"`
	// The port specified for the SLB listener.
	//
	// example:
	//
	// 8080
	TargetPort     *int32  `json:"TargetPort,omitempty" xml:"TargetPort,omitempty"`
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeApplicationSlbsResponseBodyDataIntranet) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationSlbsResponseBodyDataIntranet) GoString() string {
	return s.String()
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetConnectionDrainTimeout() *int32 {
	return s.ConnectionDrainTimeout
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetCookie() *string {
	return s.Cookie
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetCookieTimeout() *int32 {
	return s.CookieTimeout
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetEnableConnectionDrain() *bool {
	return s.EnableConnectionDrain
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetHttpsCaCertId() *string {
	return s.HttpsCaCertId
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetHttpsCertId() *string {
	return s.HttpsCertId
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetPort() *int32 {
	return s.Port
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetStickySession() *bool {
	return s.StickySession
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetStickySessionType() *string {
	return s.StickySessionType
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetTargetPort() *int32 {
	return s.TargetPort
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetConnectionDrainTimeout(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.ConnectionDrainTimeout = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetCookie(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.Cookie = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetCookieTimeout(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.CookieTimeout = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetCreateTime(v int64) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.CreateTime = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetEnableConnectionDrain(v bool) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.EnableConnectionDrain = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetHttpsCaCertId(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.HttpsCaCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetHttpsCertId(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.HttpsCertId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetPort(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.Port = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetProtocol(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.Protocol = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetStickySession(v bool) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.StickySession = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetStickySessionType(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.StickySessionType = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetTargetPort(v int32) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.TargetPort = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) SetVServerGroupId(v string) *DescribeApplicationSlbsResponseBodyDataIntranet {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeApplicationSlbsResponseBodyDataIntranet) Validate() error {
	return dara.Validate(s)
}
