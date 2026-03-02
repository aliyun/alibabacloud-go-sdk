// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcSubscribe interface {
	dara.Model
	String() string
	GoString() string
	SetDeveloperId(v string) *PbcSubscribe
	GetDeveloperId() *string
	SetId(v int64) *PbcSubscribe
	GetId() *int64
	SetPbcName(v string) *PbcSubscribe
	GetPbcName() *string
	SetPurpose(v string) *PbcSubscribe
	GetPurpose() *string
	SetSubPbcName(v string) *PbcSubscribe
	GetSubPbcName() *string
}

type PbcSubscribe struct {
	DeveloperId *string `json:"developerId,omitempty" xml:"developerId,omitempty"`
	Id          *int64  `json:"id,omitempty" xml:"id,omitempty"`
	PbcName     *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	Purpose     *string `json:"purpose,omitempty" xml:"purpose,omitempty"`
	SubPbcName  *string `json:"subPbcName,omitempty" xml:"subPbcName,omitempty"`
}

func (s PbcSubscribe) String() string {
	return dara.Prettify(s)
}

func (s PbcSubscribe) GoString() string {
	return s.String()
}

func (s *PbcSubscribe) GetDeveloperId() *string {
	return s.DeveloperId
}

func (s *PbcSubscribe) GetId() *int64 {
	return s.Id
}

func (s *PbcSubscribe) GetPbcName() *string {
	return s.PbcName
}

func (s *PbcSubscribe) GetPurpose() *string {
	return s.Purpose
}

func (s *PbcSubscribe) GetSubPbcName() *string {
	return s.SubPbcName
}

func (s *PbcSubscribe) SetDeveloperId(v string) *PbcSubscribe {
	s.DeveloperId = &v
	return s
}

func (s *PbcSubscribe) SetId(v int64) *PbcSubscribe {
	s.Id = &v
	return s
}

func (s *PbcSubscribe) SetPbcName(v string) *PbcSubscribe {
	s.PbcName = &v
	return s
}

func (s *PbcSubscribe) SetPurpose(v string) *PbcSubscribe {
	s.Purpose = &v
	return s
}

func (s *PbcSubscribe) SetSubPbcName(v string) *PbcSubscribe {
	s.SubPbcName = &v
	return s
}

func (s *PbcSubscribe) Validate() error {
	return dara.Validate(s)
}
