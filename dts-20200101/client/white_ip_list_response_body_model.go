// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iWhiteIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDestIpList(v string) *WhiteIpListResponseBody
	GetDestIpList() *string
	SetDynamicCode(v string) *WhiteIpListResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *WhiteIpListResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *WhiteIpListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *WhiteIpListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *WhiteIpListResponseBody
	GetHttpStatusCode() *int32
	SetIpList(v string) *WhiteIpListResponseBody
	GetIpList() *string
	SetRequestId(v string) *WhiteIpListResponseBody
	GetRequestId() *string
	SetSrcIpList(v string) *WhiteIpListResponseBody
	GetSrcIpList() *string
	SetSuccess(v bool) *WhiteIpListResponseBody
	GetSuccess() *bool
}

type WhiteIpListResponseBody struct {
	// Target end adaptation to VPCNAT IP whitelist
	//
	// example:
	//
	// 127.0.0.1
	DestIpList *string `json:"DestIpList,omitempty" xml:"DestIpList,omitempty"`
	// The dynamic error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 403
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The dynamic error message. This parameter will be removed in the future.
	//
	// example:
	//
	// Type
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// IP address.
	//
	// example:
	//
	// 10.151.12.0/24,47.102.181.0/24,47.101.109.0/24,120.55.129.0/24,11.115.103.0/24,47.102.234.0/24
	IpList *string `json:"IpList,omitempty" xml:"IpList,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD823BD3-1BA6-4117-A536-165CB280****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Source adaptation to VPC NAT IP whitelist
	//
	// example:
	//
	// 127.0.0.1
	SrcIpList *string `json:"SrcIpList,omitempty" xml:"SrcIpList,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s WhiteIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s WhiteIpListResponseBody) GoString() string {
	return s.String()
}

func (s *WhiteIpListResponseBody) GetDestIpList() *string {
	return s.DestIpList
}

func (s *WhiteIpListResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *WhiteIpListResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *WhiteIpListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *WhiteIpListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *WhiteIpListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *WhiteIpListResponseBody) GetIpList() *string {
	return s.IpList
}

func (s *WhiteIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *WhiteIpListResponseBody) GetSrcIpList() *string {
	return s.SrcIpList
}

func (s *WhiteIpListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *WhiteIpListResponseBody) SetDestIpList(v string) *WhiteIpListResponseBody {
	s.DestIpList = &v
	return s
}

func (s *WhiteIpListResponseBody) SetDynamicCode(v string) *WhiteIpListResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *WhiteIpListResponseBody) SetDynamicMessage(v string) *WhiteIpListResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *WhiteIpListResponseBody) SetErrCode(v string) *WhiteIpListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *WhiteIpListResponseBody) SetErrMessage(v string) *WhiteIpListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *WhiteIpListResponseBody) SetHttpStatusCode(v int32) *WhiteIpListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *WhiteIpListResponseBody) SetIpList(v string) *WhiteIpListResponseBody {
	s.IpList = &v
	return s
}

func (s *WhiteIpListResponseBody) SetRequestId(v string) *WhiteIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *WhiteIpListResponseBody) SetSrcIpList(v string) *WhiteIpListResponseBody {
	s.SrcIpList = &v
	return s
}

func (s *WhiteIpListResponseBody) SetSuccess(v bool) *WhiteIpListResponseBody {
	s.Success = &v
	return s
}

func (s *WhiteIpListResponseBody) Validate() error {
	return dara.Validate(s)
}
