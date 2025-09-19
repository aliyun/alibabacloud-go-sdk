// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInterceptionTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ModifyInterceptionTargetRequest
	GetAppName() *string
	SetImageList(v []*string) *ModifyInterceptionTargetRequest
	GetImageList() []*string
	SetNamespace(v string) *ModifyInterceptionTargetRequest
	GetNamespace() *string
	SetTagList(v []*string) *ModifyInterceptionTargetRequest
	GetTagList() []*string
	SetTargetId(v int64) *ModifyInterceptionTargetRequest
	GetTargetId() *int64
	SetTargetName(v string) *ModifyInterceptionTargetRequest
	GetTargetName() *string
	SetTargetType(v string) *ModifyInterceptionTargetRequest
	GetTargetType() *string
}

type ModifyInterceptionTargetRequest struct {
	// The name of the application.
	//
	// > You can call the [DescribeContainerTags](~~DescribeContainerTags~~) operation to obtain the value of this parameter.
	//
	// example:
	//
	// yasintt-daemonst
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// An array that consists of images.
	//
	// > You can call the [DescribeContainerTags](~~DescribeContainerTags~~) operation to obtain the value of this parameter.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace.
	//
	// > You can call the [DescribeContainerTags](~~DescribeContainerTags~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// demo4
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// An array that consists of tags.
	//
	// > You can call the [DescribeContainerTags](~~DescribeContainerTags~~) operation to obtain the value of this parameter.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The ID of the network object.
	//
	// > You can call the [ListInterceptionTargetPage](~~ListInterceptionTargetPage~~) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 400913
	TargetId *int64 `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test001
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The object type. Valid values:
	//
	// 	- **IMAGE**
	//
	// This parameter is required.
	//
	// example:
	//
	// IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s ModifyInterceptionTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInterceptionTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyInterceptionTargetRequest) GetAppName() *string {
	return s.AppName
}

func (s *ModifyInterceptionTargetRequest) GetImageList() []*string {
	return s.ImageList
}

func (s *ModifyInterceptionTargetRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyInterceptionTargetRequest) GetTagList() []*string {
	return s.TagList
}

func (s *ModifyInterceptionTargetRequest) GetTargetId() *int64 {
	return s.TargetId
}

func (s *ModifyInterceptionTargetRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *ModifyInterceptionTargetRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *ModifyInterceptionTargetRequest) SetAppName(v string) *ModifyInterceptionTargetRequest {
	s.AppName = &v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetImageList(v []*string) *ModifyInterceptionTargetRequest {
	s.ImageList = v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetNamespace(v string) *ModifyInterceptionTargetRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetTagList(v []*string) *ModifyInterceptionTargetRequest {
	s.TagList = v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetTargetId(v int64) *ModifyInterceptionTargetRequest {
	s.TargetId = &v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetTargetName(v string) *ModifyInterceptionTargetRequest {
	s.TargetName = &v
	return s
}

func (s *ModifyInterceptionTargetRequest) SetTargetType(v string) *ModifyInterceptionTargetRequest {
	s.TargetType = &v
	return s
}

func (s *ModifyInterceptionTargetRequest) Validate() error {
	return dara.Validate(s)
}
