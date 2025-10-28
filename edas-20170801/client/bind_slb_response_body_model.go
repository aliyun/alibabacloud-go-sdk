// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSlbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *BindSlbResponseBody
	GetCode() *int32
	SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody
	GetData() *BindSlbResponseBodyData
	SetMessage(v string) *BindSlbResponseBody
	GetMessage() *string
	SetRequestId(v string) *BindSlbResponseBody
	GetRequestId() *string
}

type BindSlbResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data that is returned.
	Data *BindSlbResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// bind slb success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 23DR4FDXXXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BindSlbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BindSlbResponseBody) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *BindSlbResponseBody) GetData() *BindSlbResponseBodyData {
	return s.Data
}

func (s *BindSlbResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BindSlbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BindSlbResponseBody) SetCode(v int32) *BindSlbResponseBody {
	s.Code = &v
	return s
}

func (s *BindSlbResponseBody) SetData(v *BindSlbResponseBodyData) *BindSlbResponseBody {
	s.Data = v
	return s
}

func (s *BindSlbResponseBody) SetMessage(v string) *BindSlbResponseBody {
	s.Message = &v
	return s
}

func (s *BindSlbResponseBody) SetRequestId(v string) *BindSlbResponseBody {
	s.RequestId = &v
	return s
}

func (s *BindSlbResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BindSlbResponseBodyData struct {
	// The ID of the Internet-facing SLB instance.
	//
	// example:
	//
	// “”
	ExtSlbId *string `json:"ExtSlbId,omitempty" xml:"ExtSlbId,omitempty"`
	// The IP address of the Internet-facing SLB instance.
	//
	// example:
	//
	// “”
	ExtSlbIp *string `json:"ExtSlbIp,omitempty" xml:"ExtSlbIp,omitempty"`
	// The name of the Internet-facing SLB instance.
	//
	// example:
	//
	// “”
	ExtSlbName *string `json:"ExtSlbName,omitempty" xml:"ExtSlbName,omitempty"`
	// The ID of the vServer group for the Internet-facing SLB instance.
	//
	// example:
	//
	// “”
	ExtVServerGroupId *string `json:"ExtVServerGroupId,omitempty" xml:"ExtVServerGroupId,omitempty"`
	// The ID of the internal-facing SLB instance.
	//
	// example:
	//
	// lb-wz96ph63r************
	SlbId *string `json:"SlbId,omitempty" xml:"SlbId,omitempty"`
	// The IP address of the internal-facing SLB instance.
	//
	// example:
	//
	// 192.16*.*.*
	SlbIp *string `json:"SlbIp,omitempty" xml:"SlbIp,omitempty"`
	// The name of the internal-facing SLB instance.
	//
	// example:
	//
	// test**********
	SlbName *string `json:"SlbName,omitempty" xml:"SlbName,omitempty"`
	// The listener port for the SLB instance.
	//
	// example:
	//
	// 80
	SlbPort *int32 `json:"SlbPort,omitempty" xml:"SlbPort,omitempty"`
	// The ID of the vServer group for the internal-facing SLB instance.
	//
	// example:
	//
	// “”
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s BindSlbResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BindSlbResponseBodyData) GoString() string {
	return s.String()
}

func (s *BindSlbResponseBodyData) GetExtSlbId() *string {
	return s.ExtSlbId
}

func (s *BindSlbResponseBodyData) GetExtSlbIp() *string {
	return s.ExtSlbIp
}

func (s *BindSlbResponseBodyData) GetExtSlbName() *string {
	return s.ExtSlbName
}

func (s *BindSlbResponseBodyData) GetExtVServerGroupId() *string {
	return s.ExtVServerGroupId
}

func (s *BindSlbResponseBodyData) GetSlbId() *string {
	return s.SlbId
}

func (s *BindSlbResponseBodyData) GetSlbIp() *string {
	return s.SlbIp
}

func (s *BindSlbResponseBodyData) GetSlbName() *string {
	return s.SlbName
}

func (s *BindSlbResponseBodyData) GetSlbPort() *int32 {
	return s.SlbPort
}

func (s *BindSlbResponseBodyData) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *BindSlbResponseBodyData) SetExtSlbId(v string) *BindSlbResponseBodyData {
	s.ExtSlbId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtSlbIp(v string) *BindSlbResponseBodyData {
	s.ExtSlbIp = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtSlbName(v string) *BindSlbResponseBodyData {
	s.ExtSlbName = &v
	return s
}

func (s *BindSlbResponseBodyData) SetExtVServerGroupId(v string) *BindSlbResponseBodyData {
	s.ExtVServerGroupId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbId(v string) *BindSlbResponseBodyData {
	s.SlbId = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbIp(v string) *BindSlbResponseBodyData {
	s.SlbIp = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbName(v string) *BindSlbResponseBodyData {
	s.SlbName = &v
	return s
}

func (s *BindSlbResponseBodyData) SetSlbPort(v int32) *BindSlbResponseBodyData {
	s.SlbPort = &v
	return s
}

func (s *BindSlbResponseBodyData) SetVServerGroupId(v string) *BindSlbResponseBodyData {
	s.VServerGroupId = &v
	return s
}

func (s *BindSlbResponseBodyData) Validate() error {
	return dara.Validate(s)
}
