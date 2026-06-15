// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *UnbindTagRequest
	GetAppKey() *int64
	SetClientKey(v string) *UnbindTagRequest
	GetClientKey() *string
	SetKeyType(v string) *UnbindTagRequest
	GetKeyType() *string
	SetTagName(v string) *UnbindTagRequest
	GetTagName() *string
}

type UnbindTagRequest struct {
	// The AppKey of your application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The ID of the device.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b16557673****
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// The type of the ClientKey. Valid value:
	//
	// - **DEVICE**: The key is a device ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// Specifies the tag to unbind. To unbind multiple tags, separate them with commas. A maximum of 10 tags, each up to 128 characters long, can be unbound per request. The system supports a total of 10,000 tags.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_tag1,test_tag2
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s UnbindTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindTagRequest) GoString() string {
	return s.String()
}

func (s *UnbindTagRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *UnbindTagRequest) GetClientKey() *string {
	return s.ClientKey
}

func (s *UnbindTagRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *UnbindTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *UnbindTagRequest) SetAppKey(v int64) *UnbindTagRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindTagRequest) SetClientKey(v string) *UnbindTagRequest {
	s.ClientKey = &v
	return s
}

func (s *UnbindTagRequest) SetKeyType(v string) *UnbindTagRequest {
	s.KeyType = &v
	return s
}

func (s *UnbindTagRequest) SetTagName(v string) *UnbindTagRequest {
	s.TagName = &v
	return s
}

func (s *UnbindTagRequest) Validate() error {
	return dara.Validate(s)
}
