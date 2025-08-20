// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackDriftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetectStackDriftRequest
	GetClientToken() *string
	SetLogicalResourceId(v []*string) *DetectStackDriftRequest
	GetLogicalResourceId() []*string
	SetRegionId(v string) *DetectStackDriftRequest
	GetRegionId() *string
	SetStackId(v string) *DetectStackDriftRequest
	GetStackId() *string
}

type DetectStackDriftRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The value can be up to 64 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical ID of resource.
	//
	// example:
	//
	// ScalingRule
	LogicalResourceId []*string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty" type:"Repeated"`
	// The region ID of the stack. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/131035.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the stack.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4a6c9851-3b0f-4f5f-b4ca-a14bf691****
	StackId *string `json:"StackId,omitempty" xml:"StackId,omitempty"`
}

func (s DetectStackDriftRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectStackDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackDriftRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetectStackDriftRequest) GetLogicalResourceId() []*string {
	return s.LogicalResourceId
}

func (s *DetectStackDriftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetectStackDriftRequest) GetStackId() *string {
	return s.StackId
}

func (s *DetectStackDriftRequest) SetClientToken(v string) *DetectStackDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackDriftRequest) SetLogicalResourceId(v []*string) *DetectStackDriftRequest {
	s.LogicalResourceId = v
	return s
}

func (s *DetectStackDriftRequest) SetRegionId(v string) *DetectStackDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackDriftRequest) SetStackId(v string) *DetectStackDriftRequest {
	s.StackId = &v
	return s
}

func (s *DetectStackDriftRequest) Validate() error {
	return dara.Validate(s)
}
