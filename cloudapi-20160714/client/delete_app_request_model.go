// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DeleteAppRequest
	GetAppId() *int64
	SetSecurityToken(v string) *DeleteAppRequest
	GetSecurityToken() *string
	SetTag(v []*DeleteAppRequestTag) *DeleteAppRequest
	GetTag() []*DeleteAppRequestTag
}

type DeleteAppRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110840611
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tags. Up to 20 tags are allowed.
	//
	// example:
	//
	// test2
	Tag []*DeleteAppRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DeleteAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DeleteAppRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteAppRequest) GetTag() []*DeleteAppRequestTag {
	return s.Tag
}

func (s *DeleteAppRequest) SetAppId(v int64) *DeleteAppRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppRequest) SetSecurityToken(v string) *DeleteAppRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteAppRequest) SetTag(v []*DeleteAppRequestTag) *DeleteAppRequest {
	s.Tag = v
	return s
}

func (s *DeleteAppRequest) Validate() error {
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

type DeleteAppRequestTag struct {
	// The key of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// appname
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// testapp
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeleteAppRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequestTag) GoString() string {
	return s.String()
}

func (s *DeleteAppRequestTag) GetKey() *string {
	return s.Key
}

func (s *DeleteAppRequestTag) GetValue() *string {
	return s.Value
}

func (s *DeleteAppRequestTag) SetKey(v string) *DeleteAppRequestTag {
	s.Key = &v
	return s
}

func (s *DeleteAppRequestTag) SetValue(v string) *DeleteAppRequestTag {
	s.Value = &v
	return s
}

func (s *DeleteAppRequestTag) Validate() error {
	return dara.Validate(s)
}
