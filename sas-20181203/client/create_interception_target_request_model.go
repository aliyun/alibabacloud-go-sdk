// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInterceptionTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateInterceptionTargetRequest
	GetAppName() *string
	SetClusterId(v string) *CreateInterceptionTargetRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateInterceptionTargetRequest
	GetClusterName() *string
	SetImageList(v []*string) *CreateInterceptionTargetRequest
	GetImageList() []*string
	SetNamespace(v string) *CreateInterceptionTargetRequest
	GetNamespace() *string
	SetTagList(v []*string) *CreateInterceptionTargetRequest
	GetTagList() []*string
	SetTargetName(v string) *CreateInterceptionTargetRequest
	GetTargetName() *string
	SetTargetType(v string) *CreateInterceptionTargetRequest
	GetTargetType() *string
}

type CreateInterceptionTargetRequest struct {
	// The name of the application to which the network object belongs.
	//
	// example:
	//
	// frontend
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// cc20a1024011c44b6a8710d6f8b****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// sas-test-cnnf
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// An array that consists of the images of the network object.
	ImageList []*string `json:"ImageList,omitempty" xml:"ImageList,omitempty" type:"Repeated"`
	// The namespace to which the network object belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// An array that consists of the labels specified for the network object.
	TagList []*string `json:"TagList,omitempty" xml:"TagList,omitempty" type:"Repeated"`
	// The name of the object to be blocked.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The object type. Valid value:
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

func (s CreateInterceptionTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInterceptionTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateInterceptionTargetRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateInterceptionTargetRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInterceptionTargetRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateInterceptionTargetRequest) GetImageList() []*string {
	return s.ImageList
}

func (s *CreateInterceptionTargetRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateInterceptionTargetRequest) GetTagList() []*string {
	return s.TagList
}

func (s *CreateInterceptionTargetRequest) GetTargetName() *string {
	return s.TargetName
}

func (s *CreateInterceptionTargetRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateInterceptionTargetRequest) SetAppName(v string) *CreateInterceptionTargetRequest {
	s.AppName = &v
	return s
}

func (s *CreateInterceptionTargetRequest) SetClusterId(v string) *CreateInterceptionTargetRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInterceptionTargetRequest) SetClusterName(v string) *CreateInterceptionTargetRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateInterceptionTargetRequest) SetImageList(v []*string) *CreateInterceptionTargetRequest {
	s.ImageList = v
	return s
}

func (s *CreateInterceptionTargetRequest) SetNamespace(v string) *CreateInterceptionTargetRequest {
	s.Namespace = &v
	return s
}

func (s *CreateInterceptionTargetRequest) SetTagList(v []*string) *CreateInterceptionTargetRequest {
	s.TagList = v
	return s
}

func (s *CreateInterceptionTargetRequest) SetTargetName(v string) *CreateInterceptionTargetRequest {
	s.TargetName = &v
	return s
}

func (s *CreateInterceptionTargetRequest) SetTargetType(v string) *CreateInterceptionTargetRequest {
	s.TargetType = &v
	return s
}

func (s *CreateInterceptionTargetRequest) Validate() error {
	return dara.Validate(s)
}
