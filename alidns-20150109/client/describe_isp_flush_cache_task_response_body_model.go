// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspFlushCacheTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeIspFlushCacheTaskResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeIspFlushCacheTaskResponseBody
	GetCreateTimestamp() *int64
	SetDomainName(v string) *DescribeIspFlushCacheTaskResponseBody
	GetDomainName() *string
	SetFlushCacheResults(v []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) *DescribeIspFlushCacheTaskResponseBody
	GetFlushCacheResults() []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResults
	SetInstanceId(v string) *DescribeIspFlushCacheTaskResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeIspFlushCacheTaskResponseBody
	GetInstanceName() *string
	SetIsp(v string) *DescribeIspFlushCacheTaskResponseBody
	GetIsp() *string
	SetRequestId(v string) *DescribeIspFlushCacheTaskResponseBody
	GetRequestId() *string
	SetTaskId(v string) *DescribeIspFlushCacheTaskResponseBody
	GetTaskId() *string
	SetTaskStatus(v string) *DescribeIspFlushCacheTaskResponseBody
	GetTaskStatus() *string
}

type DescribeIspFlushCacheTaskResponseBody struct {
	CreateTime        *string                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp   *int64                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	DomainName        *string                                                   `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	FlushCacheResults []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResults `json:"FlushCacheResults,omitempty" xml:"FlushCacheResults,omitempty" type:"Repeated"`
	InstanceId        *string                                                   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                                                   `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Isp               *string                                                   `json:"Isp,omitempty" xml:"Isp,omitempty"`
	RequestId         *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskId            *string                                                   `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TaskStatus        *string                                                   `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
}

func (s DescribeIspFlushCacheTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetFlushCacheResults() []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResults {
	return s.FlushCacheResults
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetIsp() *string {
	return s.Isp
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeIspFlushCacheTaskResponseBody) GetTaskStatus() *string {
	return s.TaskStatus
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetCreateTime(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetCreateTimestamp(v int64) *DescribeIspFlushCacheTaskResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetDomainName(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetFlushCacheResults(v []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) *DescribeIspFlushCacheTaskResponseBody {
	s.FlushCacheResults = v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetInstanceId(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetInstanceName(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetIsp(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.Isp = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetRequestId(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetTaskId(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) SetTaskStatus(v string) *DescribeIspFlushCacheTaskResponseBody {
	s.TaskStatus = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheTaskResponseBodyFlushCacheResults struct {
	DnsNodes []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes `json:"DnsNodes,omitempty" xml:"DnsNodes,omitempty" type:"Repeated"`
	Province *string                                                           `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) GetDnsNodes() []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes {
	return s.DnsNodes
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) GetProvince() *string {
	return s.Province
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) SetDnsNodes(v []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults {
	s.DnsNodes = v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) SetProvince(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults {
	s.Province = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResults) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes struct {
	Answers []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers `json:"Answers,omitempty" xml:"Answers,omitempty" type:"Repeated"`
	NodeIp  *string                                                                  `json:"NodeIp,omitempty" xml:"NodeIp,omitempty"`
	SpName  *string                                                                  `json:"SpName,omitempty" xml:"SpName,omitempty"`
	Status  *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) GetAnswers() []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers {
	return s.Answers
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) GetNodeIp() *string {
	return s.NodeIp
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) GetSpName() *string {
	return s.SpName
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) GetStatus() *string {
	return s.Status
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) SetAnswers(v []*DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes {
	s.Answers = v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) SetNodeIp(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes {
	s.NodeIp = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) SetSpName(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes {
	s.SpName = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) SetStatus(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes {
	s.Status = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers struct {
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Record *string `json:"Record,omitempty" xml:"Record,omitempty"`
	Ttl    *int64  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) GoString() string {
	return s.String()
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) GetName() *string {
	return s.Name
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) GetRecord() *string {
	return s.Record
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) GetTtl() *int64 {
	return s.Ttl
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) GetType() *string {
	return s.Type
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) SetName(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers {
	s.Name = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) SetRecord(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers {
	s.Record = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) SetTtl(v int64) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers {
	s.Ttl = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) SetType(v string) *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers {
	s.Type = &v
	return s
}

func (s *DescribeIspFlushCacheTaskResponseBodyFlushCacheResultsDnsNodesAnswers) Validate() error {
	return dara.Validate(s)
}
