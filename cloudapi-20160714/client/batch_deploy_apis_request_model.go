// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDeployApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v []*BatchDeployApisRequestApi) *BatchDeployApisRequest
	GetApi() []*BatchDeployApisRequestApi
	SetDescription(v string) *BatchDeployApisRequest
	GetDescription() *string
	SetSecurityToken(v string) *BatchDeployApisRequest
	GetSecurityToken() *string
	SetStageName(v string) *BatchDeployApisRequest
	GetStageName() *string
}

type BatchDeployApisRequest struct {
	// The APIs that you want to publish.
	Api []*BatchDeployApisRequestApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
	// The description.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Description   *string `json:"Description,omitempty" xml:"Description,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **TEST**
	//
	// 	- PRE: the pre-release environment
	//
	// This parameter is required.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s BatchDeployApisRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployApisRequest) GoString() string {
	return s.String()
}

func (s *BatchDeployApisRequest) GetApi() []*BatchDeployApisRequestApi {
	return s.Api
}

func (s *BatchDeployApisRequest) GetDescription() *string {
	return s.Description
}

func (s *BatchDeployApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchDeployApisRequest) GetStageName() *string {
	return s.StageName
}

func (s *BatchDeployApisRequest) SetApi(v []*BatchDeployApisRequestApi) *BatchDeployApisRequest {
	s.Api = v
	return s
}

func (s *BatchDeployApisRequest) SetDescription(v string) *BatchDeployApisRequest {
	s.Description = &v
	return s
}

func (s *BatchDeployApisRequest) SetSecurityToken(v string) *BatchDeployApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchDeployApisRequest) SetStageName(v string) *BatchDeployApisRequest {
	s.StageName = &v
	return s
}

func (s *BatchDeployApisRequest) Validate() error {
	if s.Api != nil {
		for _, item := range s.Api {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchDeployApisRequestApi struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2b35dd68345b472f8051647306a16415
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// The API group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// b4f5c342b8bc4ef88ccda0332402e0fa
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s BatchDeployApisRequestApi) String() string {
	return dara.Prettify(s)
}

func (s BatchDeployApisRequestApi) GoString() string {
	return s.String()
}

func (s *BatchDeployApisRequestApi) GetApiUid() *string {
	return s.ApiUid
}

func (s *BatchDeployApisRequestApi) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchDeployApisRequestApi) SetApiUid(v string) *BatchDeployApisRequestApi {
	s.ApiUid = &v
	return s
}

func (s *BatchDeployApisRequestApi) SetGroupId(v string) *BatchDeployApisRequestApi {
	s.GroupId = &v
	return s
}

func (s *BatchDeployApisRequestApi) Validate() error {
	return dara.Validate(s)
}
