// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSaseUserIds(v []*string) *UnbindSaseUserTagRequest
	GetSaseUserIds() []*string
	SetTagIds(v []*string) *UnbindSaseUserTagRequest
	GetTagIds() []*string
}

type UnbindSaseUserTagRequest struct {
	// The collection of user IDs.
	SaseUserIds []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// The collection of user label IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s UnbindSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *UnbindSaseUserTagRequest) GetSaseUserIds() []*string {
	return s.SaseUserIds
}

func (s *UnbindSaseUserTagRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *UnbindSaseUserTagRequest) SetSaseUserIds(v []*string) *UnbindSaseUserTagRequest {
	s.SaseUserIds = v
	return s
}

func (s *UnbindSaseUserTagRequest) SetTagIds(v []*string) *UnbindSaseUserTagRequest {
	s.TagIds = v
	return s
}

func (s *UnbindSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
