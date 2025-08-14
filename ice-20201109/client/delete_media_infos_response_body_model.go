// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMediaInfosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForbiddenList(v []*string) *DeleteMediaInfosResponseBody
	GetForbiddenList() []*string
	SetIgnoredList(v []*string) *DeleteMediaInfosResponseBody
	GetIgnoredList() []*string
	SetRequestId(v string) *DeleteMediaInfosResponseBody
	GetRequestId() *string
}

type DeleteMediaInfosResponseBody struct {
	// The IDs or URLs of media assets that cannot be deleted. Generally, media assets cannot be deleted if you do not have the required permissions.
	ForbiddenList []*string `json:"ForbiddenList,omitempty" xml:"ForbiddenList,omitempty" type:"Repeated"`
	// The IDs or URLs of ignored media assets. An error occurred while obtaining such media assets.
	IgnoredList []*string `json:"IgnoredList,omitempty" xml:"IgnoredList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 0622C702-41BE-467E-AF2E-883D4517962E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMediaInfosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMediaInfosResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMediaInfosResponseBody) GetForbiddenList() []*string {
	return s.ForbiddenList
}

func (s *DeleteMediaInfosResponseBody) GetIgnoredList() []*string {
	return s.IgnoredList
}

func (s *DeleteMediaInfosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMediaInfosResponseBody) SetForbiddenList(v []*string) *DeleteMediaInfosResponseBody {
	s.ForbiddenList = v
	return s
}

func (s *DeleteMediaInfosResponseBody) SetIgnoredList(v []*string) *DeleteMediaInfosResponseBody {
	s.IgnoredList = v
	return s
}

func (s *DeleteMediaInfosResponseBody) SetRequestId(v string) *DeleteMediaInfosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMediaInfosResponseBody) Validate() error {
	return dara.Validate(s)
}
