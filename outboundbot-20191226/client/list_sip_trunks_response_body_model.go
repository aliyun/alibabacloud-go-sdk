// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSipTrunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSipTrunksResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListSipTrunksResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListSipTrunksResponseBodyList) *ListSipTrunksResponseBody
	GetList() []*ListSipTrunksResponseBodyList
	SetMessage(v string) *ListSipTrunksResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListSipTrunksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSipTrunksResponseBody
	GetSuccess() *bool
}

type ListSipTrunksResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The list of SIP trunks.
	List []*ListSipTrunksResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListSipTrunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSipTrunksResponseBody) GoString() string {
	return s.String()
}

func (s *ListSipTrunksResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSipTrunksResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListSipTrunksResponseBody) GetList() []*ListSipTrunksResponseBodyList {
	return s.List
}

func (s *ListSipTrunksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSipTrunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSipTrunksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSipTrunksResponseBody) SetCode(v string) *ListSipTrunksResponseBody {
	s.Code = &v
	return s
}

func (s *ListSipTrunksResponseBody) SetHttpStatusCode(v int32) *ListSipTrunksResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListSipTrunksResponseBody) SetList(v []*ListSipTrunksResponseBodyList) *ListSipTrunksResponseBody {
	s.List = v
	return s
}

func (s *ListSipTrunksResponseBody) SetMessage(v string) *ListSipTrunksResponseBody {
	s.Message = &v
	return s
}

func (s *ListSipTrunksResponseBody) SetRequestId(v string) *ListSipTrunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSipTrunksResponseBody) SetSuccess(v bool) *ListSipTrunksResponseBody {
	s.Success = &v
	return s
}

func (s *ListSipTrunksResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSipTrunksResponseBodyList struct {
	// The description of the SIP trunk.
	//
	// example:
	//
	// Test environment SIP trunk
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The creation timestamp, in milliseconds.
	//
	// example:
	//
	// 1727654400000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The list of IP addresses of the SIP trunk.
	Ips []*ListSipTrunksResponseBodyListIps `json:"Ips,omitempty" xml:"Ips,omitempty" type:"Repeated"`
	// The unique identifier of the SIP trunk.
	//
	// example:
	//
	// 6bf8a1d2-3c4e-4f5a-9b8c-7d6e5f4a3b2c
	TrunkId *string `json:"TrunkId,omitempty" xml:"TrunkId,omitempty"`
	// The name of the SIP trunk.
	//
	// example:
	//
	// sip-trunk-demo
	TrunkName *string `json:"TrunkName,omitempty" xml:"TrunkName,omitempty"`
}

func (s ListSipTrunksResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListSipTrunksResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListSipTrunksResponseBodyList) GetDescription() *string {
	return s.Description
}

func (s *ListSipTrunksResponseBodyList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListSipTrunksResponseBodyList) GetIps() []*ListSipTrunksResponseBodyListIps {
	return s.Ips
}

func (s *ListSipTrunksResponseBodyList) GetTrunkId() *string {
	return s.TrunkId
}

func (s *ListSipTrunksResponseBodyList) GetTrunkName() *string {
	return s.TrunkName
}

func (s *ListSipTrunksResponseBodyList) SetDescription(v string) *ListSipTrunksResponseBodyList {
	s.Description = &v
	return s
}

func (s *ListSipTrunksResponseBodyList) SetGmtCreate(v int64) *ListSipTrunksResponseBodyList {
	s.GmtCreate = &v
	return s
}

func (s *ListSipTrunksResponseBodyList) SetIps(v []*ListSipTrunksResponseBodyListIps) *ListSipTrunksResponseBodyList {
	s.Ips = v
	return s
}

func (s *ListSipTrunksResponseBodyList) SetTrunkId(v string) *ListSipTrunksResponseBodyList {
	s.TrunkId = &v
	return s
}

func (s *ListSipTrunksResponseBodyList) SetTrunkName(v string) *ListSipTrunksResponseBodyList {
	s.TrunkName = &v
	return s
}

func (s *ListSipTrunksResponseBodyList) Validate() error {
	if s.Ips != nil {
		for _, item := range s.Ips {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSipTrunksResponseBodyListIps struct {
	// The IP address of the SIP trunk.
	//
	// example:
	//
	// 192.0.2.1
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ListSipTrunksResponseBodyListIps) String() string {
	return dara.Prettify(s)
}

func (s ListSipTrunksResponseBodyListIps) GoString() string {
	return s.String()
}

func (s *ListSipTrunksResponseBodyListIps) GetIp() *string {
	return s.Ip
}

func (s *ListSipTrunksResponseBodyListIps) SetIp(v string) *ListSipTrunksResponseBodyListIps {
	s.Ip = &v
	return s
}

func (s *ListSipTrunksResponseBodyListIps) Validate() error {
	return dara.Validate(s)
}
