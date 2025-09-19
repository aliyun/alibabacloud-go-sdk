// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagWithUuidRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTagName(v string) *DeleteTagWithUuidRequest
	GetTagName() *string
	SetUuidList(v string) *DeleteTagWithUuidRequest
	GetUuidList() *string
}

type DeleteTagWithUuidRequest struct {
	// The name of the tag.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
	// The UUIDs of servers.
	//
	// > If the UuidList parameter is specified, Security Center removes the tag only from the servers whose UUIDs are specified by UuidList. If the UuidList parameter is not specified, Security Center removes the tag from all servers.
	//
	// example:
	//
	// 111-xx,aa-bb
	UuidList *string `json:"UuidList,omitempty" xml:"UuidList,omitempty"`
}

func (s DeleteTagWithUuidRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagWithUuidRequest) GoString() string {
	return s.String()
}

func (s *DeleteTagWithUuidRequest) GetTagName() *string {
	return s.TagName
}

func (s *DeleteTagWithUuidRequest) GetUuidList() *string {
	return s.UuidList
}

func (s *DeleteTagWithUuidRequest) SetTagName(v string) *DeleteTagWithUuidRequest {
	s.TagName = &v
	return s
}

func (s *DeleteTagWithUuidRequest) SetUuidList(v string) *DeleteTagWithUuidRequest {
	s.UuidList = &v
	return s
}

func (s *DeleteTagWithUuidRequest) Validate() error {
	return dara.Validate(s)
}
