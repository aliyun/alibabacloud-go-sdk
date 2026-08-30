// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindSaseUserTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSaseUserIds(v []*string) *BindSaseUserTagRequest
	GetSaseUserIds() []*string
	SetTagIds(v []*string) *BindSaseUserTagRequest
	GetTagIds() []*string
}

type BindSaseUserTagRequest struct {
	// The collection of user IDs.
	SaseUserIds []*string `json:"SaseUserIds,omitempty" xml:"SaseUserIds,omitempty" type:"Repeated"`
	// The collection of user label IDs.
	TagIds []*string `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
}

func (s BindSaseUserTagRequest) String() string {
	return dara.Prettify(s)
}

func (s BindSaseUserTagRequest) GoString() string {
	return s.String()
}

func (s *BindSaseUserTagRequest) GetSaseUserIds() []*string {
	return s.SaseUserIds
}

func (s *BindSaseUserTagRequest) GetTagIds() []*string {
	return s.TagIds
}

func (s *BindSaseUserTagRequest) SetSaseUserIds(v []*string) *BindSaseUserTagRequest {
	s.SaseUserIds = v
	return s
}

func (s *BindSaseUserTagRequest) SetTagIds(v []*string) *BindSaseUserTagRequest {
	s.TagIds = v
	return s
}

func (s *BindSaseUserTagRequest) Validate() error {
	return dara.Validate(s)
}
