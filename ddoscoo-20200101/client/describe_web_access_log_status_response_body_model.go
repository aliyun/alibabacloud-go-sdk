// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebAccessLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeWebAccessLogStatusResponseBody
	GetRequestId() *string
	SetSlsLogstore(v string) *DescribeWebAccessLogStatusResponseBody
	GetSlsLogstore() *string
	SetSlsProject(v string) *DescribeWebAccessLogStatusResponseBody
	GetSlsProject() *string
	SetSlsStatus(v bool) *DescribeWebAccessLogStatusResponseBody
	GetSlsStatus() *bool
}

type DescribeWebAccessLogStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CF33B4C3-196E-4015-AADD-5CAD00057B80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Logstore of the instance.
	//
	// example:
	//
	// ddoscoo-logstore
	SlsLogstore *string `json:"SlsLogstore,omitempty" xml:"SlsLogstore,omitempty"`
	// The Log Service project of the instance.
	//
	// example:
	//
	// ddoscoo-project-128965410602****-cn-hangzhou
	SlsProject *string `json:"SlsProject,omitempty" xml:"SlsProject,omitempty"`
	// Indicates whether the Log Analysis feature is enabled for the website. Valid values:
	//
	// 	- **true**: enabled
	//
	// 	- **false**: disabled
	//
	// example:
	//
	// true
	SlsStatus *bool `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
}

func (s DescribeWebAccessLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebAccessLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebAccessLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebAccessLogStatusResponseBody) GetSlsLogstore() *string {
	return s.SlsLogstore
}

func (s *DescribeWebAccessLogStatusResponseBody) GetSlsProject() *string {
	return s.SlsProject
}

func (s *DescribeWebAccessLogStatusResponseBody) GetSlsStatus() *bool {
	return s.SlsStatus
}

func (s *DescribeWebAccessLogStatusResponseBody) SetRequestId(v string) *DescribeWebAccessLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsLogstore(v string) *DescribeWebAccessLogStatusResponseBody {
	s.SlsLogstore = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsProject(v string) *DescribeWebAccessLogStatusResponseBody {
	s.SlsProject = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) SetSlsStatus(v bool) *DescribeWebAccessLogStatusResponseBody {
	s.SlsStatus = &v
	return s
}

func (s *DescribeWebAccessLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
