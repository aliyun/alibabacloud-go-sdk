// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomImageShareAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v []*int64) *AddCustomImageShareAccountRequest
	GetAccount() []*int64
	SetClientToken(v string) *AddCustomImageShareAccountRequest
	GetClientToken() *string
	SetImageId(v string) *AddCustomImageShareAccountRequest
	GetImageId() *string
	SetRegionId(v string) *AddCustomImageShareAccountRequest
	GetRegionId() *string
}

type AddCustomImageShareAccountRequest struct {
	// The account name.
	//
	// This parameter is required.
	Account []*int64 `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can generate the value on your client, but you must ensure that the value is unique among different requests. The **ClientToken*	- parameter can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-saacssasc****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the latest Alibaba Cloud region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AddCustomImageShareAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AddCustomImageShareAccountRequest) GoString() string {
	return s.String()
}

func (s *AddCustomImageShareAccountRequest) GetAccount() []*int64 {
	return s.Account
}

func (s *AddCustomImageShareAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddCustomImageShareAccountRequest) GetImageId() *string {
	return s.ImageId
}

func (s *AddCustomImageShareAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddCustomImageShareAccountRequest) SetAccount(v []*int64) *AddCustomImageShareAccountRequest {
	s.Account = v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetClientToken(v string) *AddCustomImageShareAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetImageId(v string) *AddCustomImageShareAccountRequest {
	s.ImageId = &v
	return s
}

func (s *AddCustomImageShareAccountRequest) SetRegionId(v string) *AddCustomImageShareAccountRequest {
	s.RegionId = &v
	return s
}

func (s *AddCustomImageShareAccountRequest) Validate() error {
	return dara.Validate(s)
}
