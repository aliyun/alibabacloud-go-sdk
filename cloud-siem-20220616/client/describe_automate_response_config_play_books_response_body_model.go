// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutomateResponseConfigPlayBooksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponseBody
	GetCode() *int32
	SetData(v []*DescribeAutomateResponseConfigPlayBooksResponseBodyData) *DescribeAutomateResponseConfigPlayBooksResponseBody
	GetData() []*DescribeAutomateResponseConfigPlayBooksResponseBodyData
	SetMessage(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAutomateResponseConfigPlayBooksResponseBody
	GetSuccess() *bool
}

type DescribeAutomateResponseConfigPlayBooksResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeAutomateResponseConfigPlayBooksResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) GetData() []*DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	return s.Data
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetCode(v int32) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetData(v []*DescribeAutomateResponseConfigPlayBooksResponseBodyData) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetMessage(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetRequestId(v string) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) SetSuccess(v bool) *DescribeAutomateResponseConfigPlayBooksResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAutomateResponseConfigPlayBooksResponseBodyData struct {
	// The description of the playbook.
	//
	// example:
	//
	// Waf Block IP
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The unique identifier name of the playbook.
	//
	// example:
	//
	// WafBlockIP
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The input parameter template of the playbook. Valid values:
	//
	// 	- template-ip: IP address
	//
	// 	- template-process: process
	//
	// 	- template-filee: file
	//
	// example:
	//
	// template-ip
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutomateResponseConfigPlayBooksResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) GetName() *string {
	return s.Name
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) GetParamType() *string {
	return s.ParamType
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDescription(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Description = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetDisplayName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetName(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Name = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetParamType(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.ParamType = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) SetUuid(v string) *DescribeAutomateResponseConfigPlayBooksResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *DescribeAutomateResponseConfigPlayBooksResponseBodyData) Validate() error {
	return dara.Validate(s)
}
