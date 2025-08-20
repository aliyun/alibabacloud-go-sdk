// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignalResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SignalResourceRequest
	GetClientToken() *string
	SetLogicalResourceId(v string) *SignalResourceRequest
	GetLogicalResourceId() *string
	SetRegionId(v string) *SignalResourceRequest
	GetRegionId() *string
	SetStackId(v string) *SignalResourceRequest
	GetStackId() *string
	SetStatus(v string) *SignalResourceRequest
	GetStatus() *string
	SetUniqueId(v string) *SignalResourceRequest
	GetUniqueId() *string
}

type SignalResourceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests.
	//
	// The token can be up to 64 characters in length and can contain letters, digits, hyphens (-) and underscores (_).
	//
	// For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/134212.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The logical ID of the resource as defined in the template.
	//
	// This parameter is required.
	//
	// example:
	//
	// WebServer
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
	// The status of the signal. Failure signals can cause stack creation or update to fail. If all signals are warnings, the stack cannot be created or updated. Valid values:
	//
	// 	- SUCCESS
	//
	// 	- FAILURE
	//
	// 	- WARNING
	//
	// This parameter is required.
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unique ID of the signal. The ID must be 1 to 64 characters in length. If multiple signals are sent to a single resource, each signal must have a unique ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27c7347b-352a-4377-accf-63d361c1****
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s SignalResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s SignalResourceRequest) GoString() string {
	return s.String()
}

func (s *SignalResourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SignalResourceRequest) GetLogicalResourceId() *string {
	return s.LogicalResourceId
}

func (s *SignalResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SignalResourceRequest) GetStackId() *string {
	return s.StackId
}

func (s *SignalResourceRequest) GetStatus() *string {
	return s.Status
}

func (s *SignalResourceRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *SignalResourceRequest) SetClientToken(v string) *SignalResourceRequest {
	s.ClientToken = &v
	return s
}

func (s *SignalResourceRequest) SetLogicalResourceId(v string) *SignalResourceRequest {
	s.LogicalResourceId = &v
	return s
}

func (s *SignalResourceRequest) SetRegionId(v string) *SignalResourceRequest {
	s.RegionId = &v
	return s
}

func (s *SignalResourceRequest) SetStackId(v string) *SignalResourceRequest {
	s.StackId = &v
	return s
}

func (s *SignalResourceRequest) SetStatus(v string) *SignalResourceRequest {
	s.Status = &v
	return s
}

func (s *SignalResourceRequest) SetUniqueId(v string) *SignalResourceRequest {
	s.UniqueId = &v
	return s
}

func (s *SignalResourceRequest) Validate() error {
	return dara.Validate(s)
}
