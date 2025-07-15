// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConfigItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigItems(v string) *UpdateConfigItemsRequest
	GetConfigItems() *string
	SetInstanceId(v string) *UpdateConfigItemsRequest
	GetInstanceId() *string
	SetObjectId(v string) *UpdateConfigItemsRequest
	GetObjectId() *string
	SetObjectType(v string) *UpdateConfigItemsRequest
	GetObjectType() *string
}

type UpdateConfigItemsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"ShowCalledID","value":"-1"},{"name":"ShowCalleeID","value":"1"},{"name":"AllowHangup","value":"0"},{"name":"AutoAnswerCall","value":"-1"},{"name":"AllowAudioDownload","value":"1"},{"name":"AllowChooseSignedSkillGroup","value":"1"}]
	ConfigItems *string `json:"ConfigItems,omitempty" xml:"ConfigItems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// INSTANCE
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
}

func (s UpdateConfigItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConfigItemsRequest) GoString() string {
	return s.String()
}

func (s *UpdateConfigItemsRequest) GetConfigItems() *string {
	return s.ConfigItems
}

func (s *UpdateConfigItemsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateConfigItemsRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *UpdateConfigItemsRequest) GetObjectType() *string {
	return s.ObjectType
}

func (s *UpdateConfigItemsRequest) SetConfigItems(v string) *UpdateConfigItemsRequest {
	s.ConfigItems = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetInstanceId(v string) *UpdateConfigItemsRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetObjectId(v string) *UpdateConfigItemsRequest {
	s.ObjectId = &v
	return s
}

func (s *UpdateConfigItemsRequest) SetObjectType(v string) *UpdateConfigItemsRequest {
	s.ObjectType = &v
	return s
}

func (s *UpdateConfigItemsRequest) Validate() error {
	return dara.Validate(s)
}
