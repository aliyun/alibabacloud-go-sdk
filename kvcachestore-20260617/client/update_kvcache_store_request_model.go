// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKVCacheStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapacity(v int64) *UpdateKVCacheStoreRequest
	GetCapacity() *int64
	SetClientToken(v string) *UpdateKVCacheStoreRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateKVCacheStoreRequest
	GetDescription() *string
	SetKvcsId(v string) *UpdateKVCacheStoreRequest
	GetKvcsId() *string
	SetName(v string) *UpdateKVCacheStoreRequest
	GetName() *string
	SetRegionId(v string) *UpdateKVCacheStoreRequest
	GetRegionId() *string
	SetTag(v []*UpdateKVCacheStoreRequestTag) *UpdateKVCacheStoreRequest
	GetTag() []*UpdateKVCacheStoreRequestTag
}

type UpdateKVCacheStoreRequest struct {
	// The new storage capacity in GiB. The value must be a multiple of 300 TiB and greater than the current capacity.
	//
	// example:
	//
	// 4096
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// YOUR_CLIENT_TOKEN
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new KVCacheStore description. The description must be 2 to 256 characters in length and can contain English and Chinese characters. The description cannot start with http:// or https://. Default value: empty.
	//
	// example:
	//
	// project name pass the check
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The KVCacheStore instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// kvcs-your-id
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The new KVCacheStore name. The name must be 2 to 128 characters in length and can contain characters that are categorized as letter in Unicode (including English and Chinese characters) and digits. The name can contain colons (:), underscores (_), periods (.), and hyphens (-). If this parameter is not specified, the default value is the KVCacheStore ID.
	//
	// example:
	//
	// 1633730290118313-HD-m3u8
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The region ID, such as cn-hangzhou.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of resource tag key-value pairs. A maximum of 20 tags are supported.
	Tag []*UpdateKVCacheStoreRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s UpdateKVCacheStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKVCacheStoreRequest) GoString() string {
	return s.String()
}

func (s *UpdateKVCacheStoreRequest) GetCapacity() *int64 {
	return s.Capacity
}

func (s *UpdateKVCacheStoreRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateKVCacheStoreRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKVCacheStoreRequest) GetKvcsId() *string {
	return s.KvcsId
}

func (s *UpdateKVCacheStoreRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKVCacheStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateKVCacheStoreRequest) GetTag() []*UpdateKVCacheStoreRequestTag {
	return s.Tag
}

func (s *UpdateKVCacheStoreRequest) SetCapacity(v int64) *UpdateKVCacheStoreRequest {
	s.Capacity = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetClientToken(v string) *UpdateKVCacheStoreRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetDescription(v string) *UpdateKVCacheStoreRequest {
	s.Description = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetKvcsId(v string) *UpdateKVCacheStoreRequest {
	s.KvcsId = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetName(v string) *UpdateKVCacheStoreRequest {
	s.Name = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetRegionId(v string) *UpdateKVCacheStoreRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateKVCacheStoreRequest) SetTag(v []*UpdateKVCacheStoreRequestTag) *UpdateKVCacheStoreRequest {
	s.Tag = v
	return s
}

func (s *UpdateKVCacheStoreRequest) Validate() error {
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

type UpdateKVCacheStoreRequestTag struct {
	// The tag key of the resource.
	//
	// example:
	//
	// projectId
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// projectName
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s UpdateKVCacheStoreRequestTag) String() string {
	return dara.Prettify(s)
}

func (s UpdateKVCacheStoreRequestTag) GoString() string {
	return s.String()
}

func (s *UpdateKVCacheStoreRequestTag) GetTagKey() *string {
	return s.TagKey
}

func (s *UpdateKVCacheStoreRequestTag) GetTagValue() *string {
	return s.TagValue
}

func (s *UpdateKVCacheStoreRequestTag) SetTagKey(v string) *UpdateKVCacheStoreRequestTag {
	s.TagKey = &v
	return s
}

func (s *UpdateKVCacheStoreRequestTag) SetTagValue(v string) *UpdateKVCacheStoreRequestTag {
	s.TagValue = &v
	return s
}

func (s *UpdateKVCacheStoreRequestTag) Validate() error {
	return dara.Validate(s)
}
