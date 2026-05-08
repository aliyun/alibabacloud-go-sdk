// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSelectResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunUid(v string) *SelectResourceResponseBody
	GetAliyunUid() *string
	SetRequestId(v string) *SelectResourceResponseBody
	GetRequestId() *string
	SetResourceInfoList(v []*SelectResourceResponseBodyResourceInfoList) *SelectResourceResponseBody
	GetResourceInfoList() []*SelectResourceResponseBodyResourceInfoList
}

type SelectResourceResponseBody struct {
	AliyunUid *string `json:"aliyunUid,omitempty" xml:"aliyunUid,omitempty"`
	// example:
	//
	// 0E8B1746-AE35-5C4B-A3A8-345B274AE32C
	RequestId        *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceInfoList []*SelectResourceResponseBodyResourceInfoList `json:"resourceInfoList,omitempty" xml:"resourceInfoList,omitempty" type:"Repeated"`
}

func (s SelectResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SelectResourceResponseBody) GoString() string {
	return s.String()
}

func (s *SelectResourceResponseBody) GetAliyunUid() *string {
	return s.AliyunUid
}

func (s *SelectResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SelectResourceResponseBody) GetResourceInfoList() []*SelectResourceResponseBodyResourceInfoList {
	return s.ResourceInfoList
}

func (s *SelectResourceResponseBody) SetAliyunUid(v string) *SelectResourceResponseBody {
	s.AliyunUid = &v
	return s
}

func (s *SelectResourceResponseBody) SetRequestId(v string) *SelectResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *SelectResourceResponseBody) SetResourceInfoList(v []*SelectResourceResponseBodyResourceInfoList) *SelectResourceResponseBody {
	s.ResourceInfoList = v
	return s
}

func (s *SelectResourceResponseBody) Validate() error {
	if s.ResourceInfoList != nil {
		for _, item := range s.ResourceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SelectResourceResponseBodyResourceInfoList struct {
	// example:
	//
	// 111
	ExpireTime *string `json:"expireTime,omitempty" xml:"expireTime,omitempty"`
	// example:
	//
	// 11
	LastExpire *int32 `json:"lastExpire,omitempty" xml:"lastExpire,omitempty"`
	// example:
	//
	// 1249
	RemainCount *int32 `json:"remainCount,omitempty" xml:"remainCount,omitempty"`
	// example:
	//
	// 2
	ResourceType *int32 `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// second
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
}

func (s SelectResourceResponseBodyResourceInfoList) String() string {
	return dara.Prettify(s)
}

func (s SelectResourceResponseBodyResourceInfoList) GoString() string {
	return s.String()
}

func (s *SelectResourceResponseBodyResourceInfoList) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *SelectResourceResponseBodyResourceInfoList) GetLastExpire() *int32 {
	return s.LastExpire
}

func (s *SelectResourceResponseBodyResourceInfoList) GetRemainCount() *int32 {
	return s.RemainCount
}

func (s *SelectResourceResponseBodyResourceInfoList) GetResourceType() *int32 {
	return s.ResourceType
}

func (s *SelectResourceResponseBodyResourceInfoList) GetUnit() *string {
	return s.Unit
}

func (s *SelectResourceResponseBodyResourceInfoList) SetExpireTime(v string) *SelectResourceResponseBodyResourceInfoList {
	s.ExpireTime = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetLastExpire(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.LastExpire = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetRemainCount(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.RemainCount = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetResourceType(v int32) *SelectResourceResponseBodyResourceInfoList {
	s.ResourceType = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) SetUnit(v string) *SelectResourceResponseBodyResourceInfoList {
	s.Unit = &v
	return s
}

func (s *SelectResourceResponseBodyResourceInfoList) Validate() error {
	return dara.Validate(s)
}
