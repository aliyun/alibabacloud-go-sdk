// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbSlsInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuditInfo(v *DescribeInstDbSlsInfoResponseBodyAuditInfo) *DescribeInstDbSlsInfoResponseBody
	GetAuditInfo() *DescribeInstDbSlsInfoResponseBodyAuditInfo
	SetRequestId(v string) *DescribeInstDbSlsInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeInstDbSlsInfoResponseBody
	GetSuccess() *bool
}

type DescribeInstDbSlsInfoResponseBody struct {
	// The details of the SQL audit.
	AuditInfo *DescribeInstDbSlsInfoResponseBodyAuditInfo `json:"AuditInfo,omitempty" xml:"AuditInfo,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// DC3ABA3E-0F8A-4596-9104-F5155C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBody) GetAuditInfo() *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	return s.AuditInfo
}

func (s *DescribeInstDbSlsInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstDbSlsInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeInstDbSlsInfoResponseBody) SetAuditInfo(v *DescribeInstDbSlsInfoResponseBodyAuditInfo) *DescribeInstDbSlsInfoResponseBody {
	s.AuditInfo = v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetRequestId(v string) *DescribeInstDbSlsInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) SetSuccess(v bool) *DescribeInstDbSlsInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstDbSlsInfoResponseBodyAuditInfo struct {
	// The name of the LogStore.
	//
	// example:
	//
	// test
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The name of the Log Service project.
	//
	// example:
	//
	// test
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponseBodyAuditInfo) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) GetProject() *string {
	return s.Project
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetLogStore(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.LogStore = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) SetProject(v string) *DescribeInstDbSlsInfoResponseBodyAuditInfo {
	s.Project = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponseBodyAuditInfo) Validate() error {
	return dara.Validate(s)
}
