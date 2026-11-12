// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyImageShareStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyImageShareStatusRequest
	GetClientToken() *string
	SetImageId(v string) *ModifyImageShareStatusRequest
	GetImageId() *string
	SetOperation(v string) *ModifyImageShareStatusRequest
	GetOperation() *string
	SetRegionId(v string) *ModifyImageShareStatusRequest
	GetRegionId() *string
}

type ModifyImageShareStatusRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-saacssasc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Valid values:
	//
	// - Share
	//
	// - UnShare
	//
	// This parameter is required.
	//
	// example:
	//
	// Share
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The region ID of the custom image. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyImageShareStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyImageShareStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyImageShareStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyImageShareStatusRequest) GetImageId() *string {
	return s.ImageId
}

func (s *ModifyImageShareStatusRequest) GetOperation() *string {
	return s.Operation
}

func (s *ModifyImageShareStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyImageShareStatusRequest) SetClientToken(v string) *ModifyImageShareStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetImageId(v string) *ModifyImageShareStatusRequest {
	s.ImageId = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetOperation(v string) *ModifyImageShareStatusRequest {
	s.Operation = &v
	return s
}

func (s *ModifyImageShareStatusRequest) SetRegionId(v string) *ModifyImageShareStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyImageShareStatusRequest) Validate() error {
	return dara.Validate(s)
}
