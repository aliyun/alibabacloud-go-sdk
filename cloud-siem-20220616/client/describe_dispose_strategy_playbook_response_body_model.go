// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisposeStrategyPlaybookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeDisposeStrategyPlaybookResponseBody
	GetCode() *int32
	SetData(v []*DescribeDisposeStrategyPlaybookResponseBodyData) *DescribeDisposeStrategyPlaybookResponseBody
	GetData() []*DescribeDisposeStrategyPlaybookResponseBodyData
	SetMessage(v string) *DescribeDisposeStrategyPlaybookResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDisposeStrategyPlaybookResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDisposeStrategyPlaybookResponseBody
	GetSuccess() *bool
}

type DescribeDisposeStrategyPlaybookResponseBody struct {
	// The HTTP status code that is returned.
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
	Data []*DescribeDisposeStrategyPlaybookResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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

func (s DescribeDisposeStrategyPlaybookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) GetData() []*DescribeDisposeStrategyPlaybookResponseBodyData {
	return s.Data
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetCode(v int32) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetData(v []*DescribeDisposeStrategyPlaybookResponseBodyData) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetMessage(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetRequestId(v string) *DescribeDisposeStrategyPlaybookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) SetSuccess(v bool) *DescribeDisposeStrategyPlaybookResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDisposeStrategyPlaybookResponseBodyData struct {
	// The playbook name, which is the unique identifier of the playbook.
	//
	// example:
	//
	// WafBlockIP
	PlaybookName *string `json:"PlaybookName,omitempty" xml:"PlaybookName,omitempty"`
	// The UUID of the playbook.
	//
	// example:
	//
	// system_aliyun_clb_process_book
	PlaybookUuid *string `json:"PlaybookUuid,omitempty" xml:"PlaybookUuid,omitempty"`
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisposeStrategyPlaybookResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) GetPlaybookName() *string {
	return s.PlaybookName
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) GetPlaybookUuid() *string {
	return s.PlaybookUuid
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookName(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookName = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) SetPlaybookUuid(v string) *DescribeDisposeStrategyPlaybookResponseBodyData {
	s.PlaybookUuid = &v
	return s
}

func (s *DescribeDisposeStrategyPlaybookResponseBodyData) Validate() error {
	return dara.Validate(s)
}
