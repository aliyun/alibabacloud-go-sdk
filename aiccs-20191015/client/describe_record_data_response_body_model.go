// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *DescribeRecordDataResponseBody
	GetAcid() *string
	SetAgentId(v string) *DescribeRecordDataResponseBody
	GetAgentId() *string
	SetCode(v string) *DescribeRecordDataResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeRecordDataResponseBody
	GetMessage() *string
	SetOssLink(v string) *DescribeRecordDataResponseBody
	GetOssLink() *string
	SetRequestId(v string) *DescribeRecordDataResponseBody
	GetRequestId() *string
}

type DescribeRecordDataResponseBody struct {
	// example:
	//
	// 1004849****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// example:
	//
	// 1212****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// https://****
	OssLink *string `json:"OssLink,omitempty" xml:"OssLink,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRecordDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordDataResponseBody) GetAcid() *string {
	return s.Acid
}

func (s *DescribeRecordDataResponseBody) GetAgentId() *string {
	return s.AgentId
}

func (s *DescribeRecordDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRecordDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRecordDataResponseBody) GetOssLink() *string {
	return s.OssLink
}

func (s *DescribeRecordDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordDataResponseBody) SetAcid(v string) *DescribeRecordDataResponseBody {
	s.Acid = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetAgentId(v string) *DescribeRecordDataResponseBody {
	s.AgentId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetCode(v string) *DescribeRecordDataResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetMessage(v string) *DescribeRecordDataResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetOssLink(v string) *DescribeRecordDataResponseBody {
	s.OssLink = &v
	return s
}

func (s *DescribeRecordDataResponseBody) SetRequestId(v string) *DescribeRecordDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordDataResponseBody) Validate() error {
	return dara.Validate(s)
}
