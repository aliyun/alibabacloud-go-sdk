// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvatarResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQueryResourceInfoList(v []*QueryAvatarResourceResponseBodyQueryResourceInfoList) *QueryAvatarResourceResponseBody
	GetQueryResourceInfoList() []*QueryAvatarResourceResponseBodyQueryResourceInfoList
	SetRequestId(v string) *QueryAvatarResourceResponseBody
	GetRequestId() *string
}

type QueryAvatarResourceResponseBody struct {
	QueryResourceInfoList []*QueryAvatarResourceResponseBodyQueryResourceInfoList `json:"queryResourceInfoList,omitempty" xml:"queryResourceInfoList,omitempty" type:"Repeated"`
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s QueryAvatarResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarResourceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponseBody) GetQueryResourceInfoList() []*QueryAvatarResourceResponseBodyQueryResourceInfoList {
	return s.QueryResourceInfoList
}

func (s *QueryAvatarResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvatarResourceResponseBody) SetQueryResourceInfoList(v []*QueryAvatarResourceResponseBodyQueryResourceInfoList) *QueryAvatarResourceResponseBody {
	s.QueryResourceInfoList = v
	return s
}

func (s *QueryAvatarResourceResponseBody) SetRequestId(v string) *QueryAvatarResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvatarResourceResponseBody) Validate() error {
	if s.QueryResourceInfoList != nil {
		for _, item := range s.QueryResourceInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryAvatarResourceResponseBodyQueryResourceInfoList struct {
	// example:
	//
	// 21275
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// STANDARD
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 1719904342237
	ValidPeriodTime *string `json:"validPeriodTime,omitempty" xml:"validPeriodTime,omitempty"`
}

func (s QueryAvatarResourceResponseBodyQueryResourceInfoList) String() string {
	return dara.Prettify(s)
}

func (s QueryAvatarResourceResponseBodyQueryResourceInfoList) GoString() string {
	return s.String()
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) GetResourceId() *string {
	return s.ResourceId
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) GetType() *string {
	return s.Type
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) GetValidPeriodTime() *string {
	return s.ValidPeriodTime
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetResourceId(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.ResourceId = &v
	return s
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetType(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.Type = &v
	return s
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) SetValidPeriodTime(v string) *QueryAvatarResourceResponseBodyQueryResourceInfoList {
	s.ValidPeriodTime = &v
	return s
}

func (s *QueryAvatarResourceResponseBodyQueryResourceInfoList) Validate() error {
	return dara.Validate(s)
}
