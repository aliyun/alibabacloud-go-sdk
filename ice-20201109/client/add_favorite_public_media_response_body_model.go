// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFavoritePublicMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredList(v []*string) *AddFavoritePublicMediaResponseBody
	GetIgnoredList() []*string
	SetRequestId(v string) *AddFavoritePublicMediaResponseBody
	GetRequestId() *string
}

type AddFavoritePublicMediaResponseBody struct {
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddFavoritePublicMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFavoritePublicMediaResponseBody) GoString() string {
	return s.String()
}

func (s *AddFavoritePublicMediaResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *AddFavoritePublicMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFavoritePublicMediaResponseBody) SetIgnoredList(v []*string) *AddFavoritePublicMediaResponseBody {
	s.IgnoredList = v
	return s
}

func (s *AddFavoritePublicMediaResponseBody) SetRequestId(v string) *AddFavoritePublicMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFavoritePublicMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
