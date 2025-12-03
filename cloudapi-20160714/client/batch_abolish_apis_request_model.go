// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAbolishApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApi(v []*BatchAbolishApisRequestApi) *BatchAbolishApisRequest
	GetApi() []*BatchAbolishApisRequestApi
	SetSecurityToken(v string) *BatchAbolishApisRequest
	GetSecurityToken() *string
}

type BatchAbolishApisRequest struct {
	// The APIs that you want to operate.
	//
	// This parameter is required.
	Api           []*BatchAbolishApisRequestApi `json:"Api,omitempty" xml:"Api,omitempty" type:"Repeated"`
	SecurityToken *string                       `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchAbolishApisRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchAbolishApisRequest) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisRequest) GetApi() []*BatchAbolishApisRequestApi {
	return s.Api
}

func (s *BatchAbolishApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchAbolishApisRequest) SetApi(v []*BatchAbolishApisRequestApi) *BatchAbolishApisRequest {
	s.Api = v
	return s
}

func (s *BatchAbolishApisRequest) SetSecurityToken(v string) *BatchAbolishApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchAbolishApisRequest) Validate() error {
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

type BatchAbolishApisRequestApi struct {
	// The ID of the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 994f72dcdaf04af0b38022c65fdbd1ac
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// The ID of the API group.
	//
	// This parameter is required.
	//
	// example:
	//
	// ced5ab777f7b440398ea70e4470124de
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// 979fd16250644d5b82173534f465ac77
	StageId *string `json:"StageId,omitempty" xml:"StageId,omitempty"`
	// The name of the environment.
	//
	// example:
	//
	// RELEASE
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s BatchAbolishApisRequestApi) String() string {
	return dara.Prettify(s)
}

func (s BatchAbolishApisRequestApi) GoString() string {
	return s.String()
}

func (s *BatchAbolishApisRequestApi) GetApiUid() *string {
	return s.ApiUid
}

func (s *BatchAbolishApisRequestApi) GetGroupId() *string {
	return s.GroupId
}

func (s *BatchAbolishApisRequestApi) GetStageId() *string {
	return s.StageId
}

func (s *BatchAbolishApisRequestApi) GetStageName() *string {
	return s.StageName
}

func (s *BatchAbolishApisRequestApi) SetApiUid(v string) *BatchAbolishApisRequestApi {
	s.ApiUid = &v
	return s
}

func (s *BatchAbolishApisRequestApi) SetGroupId(v string) *BatchAbolishApisRequestApi {
	s.GroupId = &v
	return s
}

func (s *BatchAbolishApisRequestApi) SetStageId(v string) *BatchAbolishApisRequestApi {
	s.StageId = &v
	return s
}

func (s *BatchAbolishApisRequestApi) SetStageName(v string) *BatchAbolishApisRequestApi {
	s.StageName = &v
	return s
}

func (s *BatchAbolishApisRequestApi) Validate() error {
	return dara.Validate(s)
}
