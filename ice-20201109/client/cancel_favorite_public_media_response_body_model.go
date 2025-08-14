// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFavoritePublicMediaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoredList(v []*string) *CancelFavoritePublicMediaResponseBody
	GetIgnoredList() []*string
	SetRequestId(v string) *CancelFavoritePublicMediaResponseBody
	GetRequestId() *string
}

type CancelFavoritePublicMediaResponseBody struct {
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelFavoritePublicMediaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelFavoritePublicMediaResponseBody) GoString() string {
	return s.String()
}

func (s *CancelFavoritePublicMediaResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *CancelFavoritePublicMediaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelFavoritePublicMediaResponseBody) SetIgnoredList(v []*string) *CancelFavoritePublicMediaResponseBody {
	s.IgnoredList = v
	return s
}

func (s *CancelFavoritePublicMediaResponseBody) SetRequestId(v string) *CancelFavoritePublicMediaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelFavoritePublicMediaResponseBody) Validate() error {
	return dara.Validate(s)
}
