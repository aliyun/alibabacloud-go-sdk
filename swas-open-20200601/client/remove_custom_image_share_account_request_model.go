// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCustomImageShareAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v []*int64) *RemoveCustomImageShareAccountRequest
	GetAccount() []*int64
	SetClientToken(v string) *RemoveCustomImageShareAccountRequest
	GetClientToken() *string
	SetImageId(v string) *RemoveCustomImageShareAccountRequest
	GetImageId() *string
	SetRegionId(v string) *RemoveCustomImageShareAccountRequest
	GetRegionId() *string
}

type RemoveCustomImageShareAccountRequest struct {
	// The IDs of the Alibaba Cloud accounts with which you want to unshare the image.
	//
	// This parameter is required.
	Account []*int64 `json:"Account,omitempty" xml:"Account,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the shared custom image.
	//
	// This parameter is required.
	//
	// example:
	//
	// m-2zehv38jjmwva1ee****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RemoveCustomImageShareAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveCustomImageShareAccountRequest) GoString() string {
	return s.String()
}

func (s *RemoveCustomImageShareAccountRequest) GetAccount() []*int64 {
	return s.Account
}

func (s *RemoveCustomImageShareAccountRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveCustomImageShareAccountRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RemoveCustomImageShareAccountRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveCustomImageShareAccountRequest) SetAccount(v []*int64) *RemoveCustomImageShareAccountRequest {
	s.Account = v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetClientToken(v string) *RemoveCustomImageShareAccountRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetImageId(v string) *RemoveCustomImageShareAccountRequest {
	s.ImageId = &v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) SetRegionId(v string) *RemoveCustomImageShareAccountRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveCustomImageShareAccountRequest) Validate() error {
	return dara.Validate(s)
}
