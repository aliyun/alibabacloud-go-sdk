// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnswerLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *ModifyAnswerLibRequest
	GetLibId() *string
	SetLibName(v string) *ModifyAnswerLibRequest
	GetLibName() *string
	SetRegionId(v string) *ModifyAnswerLibRequest
	GetRegionId() *string
}

type ModifyAnswerLibRequest struct {
	// example:
	//
	// custom_xxxx
	LibId   *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyAnswerLibRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibRequest) GetLibId() *string {
	return s.LibId
}

func (s *ModifyAnswerLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *ModifyAnswerLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAnswerLibRequest) SetLibId(v string) *ModifyAnswerLibRequest {
	s.LibId = &v
	return s
}

func (s *ModifyAnswerLibRequest) SetLibName(v string) *ModifyAnswerLibRequest {
	s.LibName = &v
	return s
}

func (s *ModifyAnswerLibRequest) SetRegionId(v string) *ModifyAnswerLibRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAnswerLibRequest) Validate() error {
	return dara.Validate(s)
}
