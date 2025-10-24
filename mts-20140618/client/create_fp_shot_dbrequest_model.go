// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFpShotDBRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *CreateFpShotDBRequest
	GetConfig() *string
	SetDescription(v string) *CreateFpShotDBRequest
	GetDescription() *string
	SetModelId(v int32) *CreateFpShotDBRequest
	GetModelId() *int32
	SetName(v string) *CreateFpShotDBRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateFpShotDBRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateFpShotDBRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *CreateFpShotDBRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateFpShotDBRequest
	GetResourceOwnerId() *int64
}

type CreateFpShotDBRequest struct {
	// The configurations of the media fingerprint library. By default, this parameter is left empty. You can customize the configurations based on your business requirements. The value must be a string in the JSON format.
	//
	// example:
	//
	// null
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The description of the media fingerprint library.
	//
	// example:
	//
	// The library is a text fingerprint library.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The model ID of the media fingerprint library. To create a text fingerprint library, set the parameter to **11**. To create a video fingerprint library, set the parameter to **12**. To create an audio fingerprint library, set the parameter to **13**. To create an image fingerprint library, set the parameter to **14**.
	//
	// example:
	//
	// 11
	ModelId *int32 `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	// The name of the media fingerprint library.
	//
	// This parameter is required.
	//
	// example:
	//
	// example name
	Name                 *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s CreateFpShotDBRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFpShotDBRequest) GoString() string {
	return s.String()
}

func (s *CreateFpShotDBRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateFpShotDBRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFpShotDBRequest) GetModelId() *int32 {
	return s.ModelId
}

func (s *CreateFpShotDBRequest) GetName() *string {
	return s.Name
}

func (s *CreateFpShotDBRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateFpShotDBRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateFpShotDBRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateFpShotDBRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateFpShotDBRequest) SetConfig(v string) *CreateFpShotDBRequest {
	s.Config = &v
	return s
}

func (s *CreateFpShotDBRequest) SetDescription(v string) *CreateFpShotDBRequest {
	s.Description = &v
	return s
}

func (s *CreateFpShotDBRequest) SetModelId(v int32) *CreateFpShotDBRequest {
	s.ModelId = &v
	return s
}

func (s *CreateFpShotDBRequest) SetName(v string) *CreateFpShotDBRequest {
	s.Name = &v
	return s
}

func (s *CreateFpShotDBRequest) SetOwnerAccount(v string) *CreateFpShotDBRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateFpShotDBRequest) SetOwnerId(v int64) *CreateFpShotDBRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateFpShotDBRequest) SetResourceOwnerAccount(v string) *CreateFpShotDBRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateFpShotDBRequest) SetResourceOwnerId(v int64) *CreateFpShotDBRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateFpShotDBRequest) Validate() error {
	return dara.Validate(s)
}
