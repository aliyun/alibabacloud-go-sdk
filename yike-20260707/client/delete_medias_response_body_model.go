// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenList(v []*string) *DeleteMediasResponseBody
	GetForbiddenList() []*string
	SetIgnoredList(v []*string) *DeleteMediasResponseBody
	GetIgnoredList() []*string
	SetRequestId(v string) *DeleteMediasResponseBody
	GetRequestId() *string
}

type DeleteMediasResponseBody struct {
	ForbiddenList []*string `json:"ForbiddenList,omitempty" xml:"ForbiddenList,omitempty" type:"Repeated"`
	IgnoredList   []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediasResponseBody) GetForbiddenList() []*string {
	return s.ForbiddenList
}

func (s *DeleteMediasResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *DeleteMediasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediasResponseBody) SetForbiddenList(v []*string) *DeleteMediasResponseBody {
	s.ForbiddenList = v
	return s
}

func (s *DeleteMediasResponseBody) SetIgnoredList(v []*string) *DeleteMediasResponseBody {
	s.IgnoredList = v
	return s
}

func (s *DeleteMediasResponseBody) SetRequestId(v string) *DeleteMediasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediasResponseBody) Validate() error {
	return dara.Validate(s)
}
