// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogRevision(v int64) *DeleteSkillResponseBody
	GetCatalogRevision() *int64
	SetDeleted(v bool) *DeleteSkillResponseBody
	GetDeleted() *bool
	SetRequestId(v string) *DeleteSkillResponseBody
	GetRequestId() *string
	SetResult(v string) *DeleteSkillResponseBody
	GetResult() *string
	SetSkillId(v string) *DeleteSkillResponseBody
	GetSkillId() *string
}

type DeleteSkillResponseBody struct {
	// The Skill catalog revision number.
	//
	// example:
	//
	// 1
	CatalogRevision *int64 `json:"CatalogRevision,omitempty" xml:"CatalogRevision,omitempty"`
	// Indicates whether the Skill is deleted.
	Deleted *bool `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The unique identifier of the request.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned result.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// The ID of the deleted Skill.
	//
	// example:
	//
	// skill-example
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s DeleteSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillResponseBody) GetCatalogRevision() *int64 {
	return s.CatalogRevision
}

func (s *DeleteSkillResponseBody) GetDeleted() *bool {
	return s.Deleted
}

func (s *DeleteSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *DeleteSkillResponseBody) SetCatalogRevision(v int64) *DeleteSkillResponseBody {
	s.CatalogRevision = &v
	return s
}

func (s *DeleteSkillResponseBody) SetDeleted(v bool) *DeleteSkillResponseBody {
	s.Deleted = &v
	return s
}

func (s *DeleteSkillResponseBody) SetRequestId(v string) *DeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillResponseBody) SetResult(v string) *DeleteSkillResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSkillResponseBody) SetSkillId(v string) *DeleteSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *DeleteSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
