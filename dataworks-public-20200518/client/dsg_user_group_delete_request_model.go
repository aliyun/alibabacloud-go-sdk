// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgUserGroupDeleteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v []*int64) *DsgUserGroupDeleteRequest
	GetIds() []*int64
}

type DsgUserGroupDeleteRequest struct {
	// The information about the user group.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s DsgUserGroupDeleteRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgUserGroupDeleteRequest) GoString() string {
	return s.String()
}

func (s *DsgUserGroupDeleteRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *DsgUserGroupDeleteRequest) SetIds(v []*int64) *DsgUserGroupDeleteRequest {
	s.Ids = v
	return s
}

func (s *DsgUserGroupDeleteRequest) Validate() error {
	return dara.Validate(s)
}
