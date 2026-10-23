// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceImageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreParamValidation(v bool) *UpdateInstanceImageRequest
	GetIgnoreParamValidation() *bool
	SetImageId(v string) *UpdateInstanceImageRequest
	GetImageId() *string
	SetInstanceIdList(v []*string) *UpdateInstanceImageRequest
	GetInstanceIdList() []*string
	SetReset(v bool) *UpdateInstanceImageRequest
	GetReset() *bool
}

type UpdateInstanceImageRequest struct {
	// Specifies whether to skip parameter validation. A value of true indicates that validation is skipped. A value of false or an empty value indicates that validation is performed.
	IgnoreParamValidation *bool `json:"IgnoreParamValidation,omitempty" xml:"IgnoreParamValidation,omitempty"`
	// The image ID. You can call the [DescribeImageList](https://help.aliyun.com/document_detail/2807324.html) operation to query the image list of cloud phones.
	//
	// example:
	//
	// imgc-075cllfeuazh0****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The list of instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
	// Specifies whether to reset the instance. A reset clears all user data on the instance.
	//
	// example:
	//
	// false
	Reset *bool `json:"Reset,omitempty" xml:"Reset,omitempty"`
}

func (s UpdateInstanceImageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceImageRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceImageRequest) GetIgnoreParamValidation() *bool {
	return s.IgnoreParamValidation
}

func (s *UpdateInstanceImageRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateInstanceImageRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *UpdateInstanceImageRequest) GetReset() *bool {
	return s.Reset
}

func (s *UpdateInstanceImageRequest) SetIgnoreParamValidation(v bool) *UpdateInstanceImageRequest {
	s.IgnoreParamValidation = &v
	return s
}

func (s *UpdateInstanceImageRequest) SetImageId(v string) *UpdateInstanceImageRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateInstanceImageRequest) SetInstanceIdList(v []*string) *UpdateInstanceImageRequest {
	s.InstanceIdList = v
	return s
}

func (s *UpdateInstanceImageRequest) SetReset(v bool) *UpdateInstanceImageRequest {
	s.Reset = &v
	return s
}

func (s *UpdateInstanceImageRequest) Validate() error {
	return dara.Validate(s)
}
