// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsInstanceInServiceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNeedResourceInfo(v bool) *OnsInstanceInServiceListRequest
	GetNeedResourceInfo() *bool
	SetTag(v []*OnsInstanceInServiceListRequestTag) *OnsInstanceInServiceListRequest
	GetTag() []*OnsInstanceInServiceListRequestTag
}

type OnsInstanceInServiceListRequest struct {
	// Specifies whether you want the resource information to be returned.
	//
	// example:
	//
	// true
	NeedResourceInfo *bool `json:"NeedResourceInfo,omitempty" xml:"NeedResourceInfo,omitempty"`
	// The tags that you want to attach to the instance. You can attach up to 20 tags to the instance.
	Tag []*OnsInstanceInServiceListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s OnsInstanceInServiceListRequest) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListRequest) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListRequest) GetNeedResourceInfo() *bool {
	return s.NeedResourceInfo
}

func (s *OnsInstanceInServiceListRequest) GetTag() []*OnsInstanceInServiceListRequestTag {
	return s.Tag
}

func (s *OnsInstanceInServiceListRequest) SetNeedResourceInfo(v bool) *OnsInstanceInServiceListRequest {
	s.NeedResourceInfo = &v
	return s
}

func (s *OnsInstanceInServiceListRequest) SetTag(v []*OnsInstanceInServiceListRequestTag) *OnsInstanceInServiceListRequest {
	s.Tag = v
	return s
}

func (s *OnsInstanceInServiceListRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type OnsInstanceInServiceListRequestTag struct {
	// The tag key. This parameter is not required. If you configure this parameter, you must also configure **Value**.***	- If you configure Key and Value in a request, this operation queries only the instances that use the specified tags. If you do not configure these parameters in a request, this operation queries all instances that you can access.
	//
	// 	- The value of this parameter cannot be an empty string.
	//
	// 	- The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// CartService
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. This parameter is not required. If you configure this parameter, you must also configure **Key**.***	- If you configure Key and Value in a request, this operation queries only the instances that use the specified tags. If you do not configure these parameters in a request, this operation queries all instances that you can access.
	//
	// 	- The value of this parameter can be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// SericeA
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s OnsInstanceInServiceListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s OnsInstanceInServiceListRequestTag) GoString() string {
	return s.String()
}

func (s *OnsInstanceInServiceListRequestTag) GetKey() *string {
	return s.Key
}

func (s *OnsInstanceInServiceListRequestTag) GetValue() *string {
	return s.Value
}

func (s *OnsInstanceInServiceListRequestTag) SetKey(v string) *OnsInstanceInServiceListRequestTag {
	s.Key = &v
	return s
}

func (s *OnsInstanceInServiceListRequestTag) SetValue(v string) *OnsInstanceInServiceListRequestTag {
	s.Value = &v
	return s
}

func (s *OnsInstanceInServiceListRequestTag) Validate() error {
	return dara.Validate(s)
}
