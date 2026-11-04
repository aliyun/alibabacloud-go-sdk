// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePictureSearchAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *CreatePictureSearchAppRequest
	GetDesc() *string
	SetIotInstanceId(v string) *CreatePictureSearchAppRequest
	GetIotInstanceId() *string
	SetName(v string) *CreatePictureSearchAppRequest
	GetName() *string
}

type CreatePictureSearchAppRequest struct {
	// The description of the image search application instance.
	//
	// The description can contain Chinese characters, uppercase and lowercase English letters, digits, underscores (_), and hyphens (-). It can be up to 1,024 characters in length, and each Chinese character counts as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageSearchDemo
	Desc *string `json:"Desc,omitempty" xml:"Desc,omitempty"`
	// The instance ID.
	//
	// You can view the **ID*	- of the current instance on the **Instance Overview*	- page in the IoT Platform console. If the instance has an ID, you must specify this ID. Otherwise, the call fails.
	//
	// > If the public instance does not have an ID, see [How to obtain an instance ID](https://help.aliyun.com/document_detail/267533.html).
	//
	// example:
	//
	// iot_instc_pu****_c*-v64********
	IotInstanceId *string `json:"IotInstanceId,omitempty" xml:"IotInstanceId,omitempty"`
	// The name of the image search application instance.
	//
	// The name can contain Chinese characters, uppercase and lowercase English letters, digits, underscores (_), and hyphens (-). It can be up to 64 characters in length, and each Chinese character counts as two characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// ImageSearch
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreatePictureSearchAppRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePictureSearchAppRequest) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchAppRequest) GetDesc() *string {
	return s.Desc
}

func (s *CreatePictureSearchAppRequest) GetIotInstanceId() *string {
	return s.IotInstanceId
}

func (s *CreatePictureSearchAppRequest) GetName() *string {
	return s.Name
}

func (s *CreatePictureSearchAppRequest) SetDesc(v string) *CreatePictureSearchAppRequest {
	s.Desc = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetIotInstanceId(v string) *CreatePictureSearchAppRequest {
	s.IotInstanceId = &v
	return s
}

func (s *CreatePictureSearchAppRequest) SetName(v string) *CreatePictureSearchAppRequest {
	s.Name = &v
	return s
}

func (s *CreatePictureSearchAppRequest) Validate() error {
	return dara.Validate(s)
}
