// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListExecutorsResponseBody
	GetCode() *int32
	SetData(v []*ListExecutorsResponseBodyData) *ListExecutorsResponseBody
	GetData() []*ListExecutorsResponseBodyData
	SetMessage(v string) *ListExecutorsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExecutorsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListExecutorsResponseBody
	GetSuccess() *bool
}

type ListExecutorsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data []*ListExecutorsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListExecutorsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListExecutorsResponseBody) GetData() []*ListExecutorsResponseBodyData {
	return s.Data
}

func (s *ListExecutorsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExecutorsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutorsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListExecutorsResponseBody) SetCode(v int32) *ListExecutorsResponseBody {
	s.Code = &v
	return s
}

func (s *ListExecutorsResponseBody) SetData(v []*ListExecutorsResponseBodyData) *ListExecutorsResponseBody {
	s.Data = v
	return s
}

func (s *ListExecutorsResponseBody) SetMessage(v string) *ListExecutorsResponseBody {
	s.Message = &v
	return s
}

func (s *ListExecutorsResponseBody) SetRequestId(v string) *ListExecutorsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutorsResponseBody) SetSuccess(v bool) *ListExecutorsResponseBody {
	s.Success = &v
	return s
}

func (s *ListExecutorsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsResponseBodyData struct {
	// example:
	//
	// http://192.168.1.10:9999/
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 192.168.1.10
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// true
	IsDesignated *bool `json:"IsDesignated,omitempty" xml:"IsDesignated,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// true
	Online *bool `json:"Online,omitempty" xml:"Online,omitempty"`
	// example:
	//
	// 9999
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// 2.0.2
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s ListExecutorsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExecutorsResponseBodyData) GetAddress() *string {
	return s.Address
}

func (s *ListExecutorsResponseBodyData) GetIp() *string {
	return s.Ip
}

func (s *ListExecutorsResponseBodyData) GetIsDesignated() *bool {
	return s.IsDesignated
}

func (s *ListExecutorsResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *ListExecutorsResponseBodyData) GetOnline() *bool {
	return s.Online
}

func (s *ListExecutorsResponseBodyData) GetPort() *int32 {
	return s.Port
}

func (s *ListExecutorsResponseBodyData) GetVersion() *string {
	return s.Version
}

func (s *ListExecutorsResponseBodyData) GetWeight() *int32 {
	return s.Weight
}

func (s *ListExecutorsResponseBodyData) SetAddress(v string) *ListExecutorsResponseBodyData {
	s.Address = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetIp(v string) *ListExecutorsResponseBodyData {
	s.Ip = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetIsDesignated(v bool) *ListExecutorsResponseBodyData {
	s.IsDesignated = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetLabel(v string) *ListExecutorsResponseBodyData {
	s.Label = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetOnline(v bool) *ListExecutorsResponseBodyData {
	s.Online = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetPort(v int32) *ListExecutorsResponseBodyData {
	s.Port = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetVersion(v string) *ListExecutorsResponseBodyData {
	s.Version = &v
	return s
}

func (s *ListExecutorsResponseBodyData) SetWeight(v int32) *ListExecutorsResponseBodyData {
	s.Weight = &v
	return s
}

func (s *ListExecutorsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
