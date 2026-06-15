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
	// The AppKey of your application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// The ID of the target device. You can specify a maximum of 1,000 device IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b16557673****
	ClientKey *string `json:"ClientKey,omitempty" xml:"ClientKey,omitempty"`
	// The type of the `ClientKey`. Valid value:
	//
	// - **DEVICE**: Indicates a device target.
	//
	// This parameter is required.
	//
	// example:
	//
	// DEVICE
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The tags to bind. Separate multiple tags with commas (,). You can bind up to 10 tags per request.
	//
	// A tag name can be up to 128 characters long (each Chinese character counts as 1 character). Each application can have up to 10,000 tags. A single device can be bound to multiple tags.
	//
	// 	Notice:
	//
	// Do not bind a single tag to more than 100,000 devices. This practice can increase push processing time and increase response time.
	//
	// - Use the full push feature to send notifications to all devices.
	//
	// - Split the device set into multiple fine-grained tags and call the push API in batches.
	//
	//
	//
	// > - If you attempt to bind the same tag multiple times, the system automatically removes the duplicates.
	//
	// >
	//
	// > - When a user uninstalls the application from a device, the tags associated with that device are automatically unbound. This unbinding process may be slightly delayed.
	//
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
