// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectStackResourceDriftRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DetectStackResourceDriftRequest
	GetClientToken() *string
	SetLogicalResourceId(v string) *DetectStackResourceDriftRequest
	GetLogicalResourceId() *string
	SetRegionId(v string) *DetectStackResourceDriftRequest
	GetRegionId() *string
	SetStackId(v string) *DetectStackResourceDriftRequest
	GetStackId() *string
}

type DetectStackResourceDriftRequest struct {
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
	// The name of the resource.
	//
	// This parameter is required.
	//
	// example:
	//
	// ScalingRuleName
	LogicalResourceId *string `json:"LogicalResourceId,omitempty" xml:"LogicalResourceId,omitempty"`
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

func (s DetectStackResourceDriftRequest) String() string {
	return dara.Prettify(s)
}

func (s DetectStackResourceDriftRequest) GoString() string {
	return s.String()
}

func (s *DetectStackResourceDriftRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DetectStackResourceDriftRequest) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *DetectStackResourceDriftRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DetectStackResourceDriftRequest) GetStackId() *string {
	return s.StackId
}

func (s *DetectStackResourceDriftRequest) SetClientToken(v string) *DetectStackResourceDriftRequest {
	s.ClientToken = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetLogicalResourceId(v string) *DetectStackResourceDriftRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetRegionId(v string) *DetectStackResourceDriftRequest {
	s.RegionId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) SetStackId(v string) *DetectStackResourceDriftRequest {
	s.StackId = &v
	return s
}

func (s *DetectStackResourceDriftRequest) Validate() error {
	return dara.Validate(s)
}
