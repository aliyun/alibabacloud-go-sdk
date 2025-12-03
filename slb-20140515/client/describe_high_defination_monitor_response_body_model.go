// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHighDefinationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogProject(v string) *DescribeHighDefinationMonitorResponseBody
	GetLogProject() *string
	SetLogStore(v string) *DescribeHighDefinationMonitorResponseBody
	GetLogStore() *string
	SetRequestId(v string) *DescribeHighDefinationMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeHighDefinationMonitorResponseBody
	GetSuccess() *string
}

type DescribeHighDefinationMonitorResponseBody struct {
	// The name of the Log Service project.
	//
	// example:
	//
	// my-project
	LogProject *string `json:"LogProject,omitempty" xml:"LogProject,omitempty"`
	// The name of the Logstore.
	//
	// example:
	//
	// my-log-store
	LogStore *string `json:"LogStore,omitempty" xml:"LogStore,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2F398FF5-B349-5C01-8638-8E9A0BF1DBE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeHighDefinationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHighDefinationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHighDefinationMonitorResponseBody) GetLogProject() *string {
	return s.LogProject
}

func (s *DescribeHighDefinationMonitorResponseBody) GetLogStore() *string {
	return s.LogStore
}

func (s *DescribeHighDefinationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHighDefinationMonitorResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeHighDefinationMonitorResponseBody) SetLogProject(v string) *DescribeHighDefinationMonitorResponseBody {
	s.LogProject = &v
	return s
}

func (s *DescribeHighDefinationMonitorResponseBody) SetLogStore(v string) *DescribeHighDefinationMonitorResponseBody {
	s.LogStore = &v
	return s
}

func (s *DescribeHighDefinationMonitorResponseBody) SetRequestId(v string) *DescribeHighDefinationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHighDefinationMonitorResponseBody) SetSuccess(v string) *DescribeHighDefinationMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeHighDefinationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
