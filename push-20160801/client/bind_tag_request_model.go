// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *BindTagRequest
	GetAppKey() *int64
	SetClientKey(v string) *BindTagRequest
	GetClientKey() *string
	SetKeyType(v string) *BindTagRequest
	GetKeyType() *string
	SetTagName(v string) *BindTagRequest
	GetTagName() *string
}

type BindTagRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b16557673****
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test_tag,test_tag2
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s BindTagRequest) String() string {
	return dara.Prettify(s)
}

func (s BindTagRequest) GoString() string {
	return s.String()
}

func (s *BindTagRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *BindTagRequest) GetClientKey() *string {
	return s.ClientKey
}

func (s *BindTagRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *BindTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *BindTagRequest) SetAppKey(v int64) *BindTagRequest {
	s.AppKey = &v
	return s
}

func (s *BindTagRequest) SetClientKey(v string) *BindTagRequest {
	s.ClientKey = &v
	return s
}

func (s *BindTagRequest) SetKeyType(v string) *BindTagRequest {
	s.KeyType = &v
	return s
}

func (s *BindTagRequest) SetTagName(v string) *BindTagRequest {
	s.TagName = &v
	return s
}

func (s *BindTagRequest) Validate() error {
	return dara.Validate(s)
}
