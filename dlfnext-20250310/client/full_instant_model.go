// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFullInstant interface {
	dara.Model
	String() string
	GoString() string
	SetSnapshotId(v int64) *FullInstant
	GetSnapshotId() *int64
	SetTagName(v string) *FullInstant
	GetTagName() *string
	SetType(v string) *FullInstant
	GetType() *string
}

type FullInstant struct {
	SnapshotId *int64  `json:"snapshotId,omitempty" xml:"snapshotId,omitempty"`
	TagName    *string `json:"tagName,omitempty" xml:"tagName,omitempty"`
	Type       *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FullInstant) String() string {
	return dara.Prettify(s)
}

func (s FullInstant) GoString() string {
	return s.String()
}

func (s *FullInstant) GetSnapshotId() *int64 {
	return s.SnapshotId
}

func (s *FullInstant) GetTagName() *string {
	return s.TagName
}

func (s *FullInstant) GetType() *string {
	return s.Type
}

func (s *FullInstant) SetSnapshotId(v int64) *FullInstant {
	s.SnapshotId = &v
	return s
}

func (s *FullInstant) SetTagName(v string) *FullInstant {
	s.TagName = &v
	return s
}

func (s *FullInstant) SetType(v string) *FullInstant {
	s.Type = &v
	return s
}

func (s *FullInstant) Validate() error {
	return dara.Validate(s)
}
