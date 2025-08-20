// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStackResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetStackResourceRequest
	GetClientToken() *string
	SetLogicalResourceId(v string) *GetStackResourceRequest
	GetLogicalResourceId() *string
	SetRegionId(v string) *GetStackResourceRequest
	GetRegionId() *string
	SetResourceAttributes(v []*string) *GetStackResourceRequest
	GetResourceAttributes() []*string
	SetShowResourceAttributes(v bool) *GetStackResourceRequest
	GetShowResourceAttributes() *bool
	SetStackId(v string) *GetStackResourceRequest
	GetStackId() *string
}

type GetStackResourceRequest struct {
	// Specifies whether to query the resource properties. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The name of resource property N that you want to query.
	//
	// >  Maximum value of N: 20.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebServer
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
	// The logical ID of the resource defined in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the resource. Valid values:
	//
	// 	- CREATE_COMPLETE
	//
	// 	- CREATE_FAILED
	//
	// 	- CREATE_IN_PROGRESS
	//
	// 	- UPDATE_IN_PROGRESS
	//
	// 	- UPDATE_FAILED
	//
	// 	- UPDATE_COMPLETE
	//
	// 	- DELETE_IN_PROGRESS
	//
	// 	- DELETE_FAILED
	//
	// 	- CHECK_IN_PROGRESS
	//
	// 	- CHECK_FAILED
	//
	// 	- CHECK_COMPLETE
	//
	// 	- IMPORT_IN_PROGRESS
	//
	// 	- IMPORT_FAILED
	//
	// 	- IMPORT_COMPLETE
	ResourceAttributes []*string `json:"ResourceAttributes,omitempty" xml:"ResourceAttributes,omitempty" type:"Repeated"`
	// The name of resource property N that you want to query.
	//
	// example:
	//
	// true
	ShowResourceAttributes *bool `json:"ShowResourceAttributes,omitempty" xml:"ShowResourceAttributes,omitempty"`
	// The ID of the region to which the stack belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s GetStackResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStackResourceRequest) GoString() string {
	return s.String()
}

func (s *GetStackResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetStackResourceRequest) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *GetStackResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetStackResourceRequest) GetResourceAttributes() []*string {
	return s.ResourceAttributes
}

func (s *GetStackResourceRequest) GetShowResourceAttributes() *bool {
	return s.ShowResourceAttributes
}

func (s *GetStackResourceRequest) GetStackId() *string {
	return s.StackId
}

func (s *GetStackResourceRequest) SetClientToken(v string) *GetStackResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *GetStackResourceRequest) SetLogicalResourceId(v string) *GetStackResourceRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *GetStackResourceRequest) SetRegionId(v string) *GetStackResourceRequest {
	s.RegionId = &v
	return s
}

func (s *GetStackResourceRequest) SetResourceAttributes(v []*string) *GetStackResourceRequest {
	s.ResourceAttributes = v
	return s
}

func (s *GetStackResourceRequest) SetShowResourceAttributes(v bool) *GetStackResourceRequest {
	s.ShowResourceAttributes = &v
	return s
}

func (s *GetStackResourceRequest) SetStackId(v string) *GetStackResourceRequest {
	s.StackId = &v
	return s
}

func (s *GetStackResourceRequest) Validate() error {
	return dara.Validate(s)
}
