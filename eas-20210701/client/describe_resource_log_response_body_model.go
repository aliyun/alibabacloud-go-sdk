// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStore(v string) *DescribeResourceLogResponseBody
	GetLogStore() *string
	SetMessage(v string) *DescribeResourceLogResponseBody
	GetMessage() *string
	SetProjectName(v string) *DescribeResourceLogResponseBody
	GetProjectName() *string
	SetRequestId(v string) *DescribeResourceLogResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeResourceLogResponseBody
	GetStatus() *string
}

type DescribeResourceLogResponseBody struct {
	// The Logstore of Log Service.
	//
	// example:
	//
	// access_log
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Network interfaces are updating
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Log Service project that is associated with the resource group.
	//
	// example:
	//
	// eas-r-asdasdasd-sls
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the resource group.
	//
	// example:
	//
	// ResourceReady
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeResourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResourceLogResponseBody) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeResourceLogResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResourceLogResponseBody) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeResourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResourceLogResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeResourceLogResponseBody) SetLogStore(v string) *DescribeResourceLogResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetMessage(v string) *DescribeResourceLogResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetProjectName(v string) *DescribeResourceLogResponseBody {
	s.ProjectName = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetRequestId(v string) *DescribeResourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResourceLogResponseBody) SetStatus(v string) *DescribeResourceLogResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeResourceLogResponseBody) Validate() error {
	return dara.Validate(s)
}
