// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePlayInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenList(v []*string) *DeletePlayInfoResponseBody
	GetForbiddenList() []*string
	SetIgnoredList(v []*string) *DeletePlayInfoResponseBody
	GetIgnoredList() []*string
	SetRequestId(v string) *DeletePlayInfoResponseBody
	GetRequestId() *string
}

type DeletePlayInfoResponseBody struct {
	// The URLs of the media streams that cannot be deleted. Generally, media streams cannot be deleted if you do not have the required permissions.
	ForbiddenList []*string `json:"ForbiddenList,omitempty" xml:"ForbiddenList,omitempty" type:"Repeated"`
	// The URLs of ignored media streams. An error occurred while obtaining such media assets because the IDs or URLs of the media assets do not exist.
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePlayInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePlayInfoResponseBody) GetForbiddenList() []*string {
	return s.ForbiddenList
}

func (s *DeletePlayInfoResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *DeletePlayInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePlayInfoResponseBody) SetForbiddenList(v []*string) *DeletePlayInfoResponseBody {
	s.ForbiddenList = v
	return s
}

func (s *DeletePlayInfoResponseBody) SetIgnoredList(v []*string) *DeletePlayInfoResponseBody {
	s.IgnoredList = v
	return s
}

func (s *DeletePlayInfoResponseBody) SetRequestId(v string) *DeletePlayInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePlayInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
