// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPerspectivesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPerspectives(v []*QueryPerspectivesResponseBodyPerspectives) *QueryPerspectivesResponseBody
	GetPerspectives() []*QueryPerspectivesResponseBodyPerspectives
	SetRequestId(v string) *QueryPerspectivesResponseBody
	GetRequestId() *string
}

type QueryPerspectivesResponseBody struct {
	Perspectives []*QueryPerspectivesResponseBodyPerspectives `json:"Perspectives,omitempty" xml:"Perspectives,omitempty" type:"Repeated"`
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPerspectivesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPerspectivesResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBody) GetPerspectives() []*QueryPerspectivesResponseBodyPerspectives {
	return s.Perspectives
}

func (s *QueryPerspectivesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPerspectivesResponseBody) SetPerspectives(v []*QueryPerspectivesResponseBodyPerspectives) *QueryPerspectivesResponseBody {
	s.Perspectives = v
	return s
}

func (s *QueryPerspectivesResponseBody) SetRequestId(v string) *QueryPerspectivesResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPerspectivesResponseBody) Validate() error {
	if s.Perspectives != nil {
		for _, item := range s.Perspectives {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPerspectivesResponseBodyPerspectives struct {
	// example:
	//
	// 2022-04-12T06:30:17Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2022-04-29T03:38:54Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// FZJBY3raWr
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// example:
	//
	// true
	SelfDefine *bool `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryPerspectivesResponseBodyPerspectives) String() string {
	return dara.Prettify(s)
}

func (s QueryPerspectivesResponseBodyPerspectives) GoString() string {
	return s.String()
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetCreateTime() *string {
	return s.CreateTime
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetName() *string {
	return s.Name
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetPerspectiveCode() *string {
	return s.PerspectiveCode
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetSelfDefine() *bool {
	return s.SelfDefine
}

func (s *QueryPerspectivesResponseBodyPerspectives) GetStatus() *int32 {
	return s.Status
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetCreateTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.CreateTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetModifyTime(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.ModifyTime = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetName(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.Name = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveCode(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveCode = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetPerspectiveId(v string) *QueryPerspectivesResponseBodyPerspectives {
	s.PerspectiveId = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetSelfDefine(v bool) *QueryPerspectivesResponseBodyPerspectives {
	s.SelfDefine = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) SetStatus(v int32) *QueryPerspectivesResponseBodyPerspectives {
	s.Status = &v
	return s
}

func (s *QueryPerspectivesResponseBodyPerspectives) Validate() error {
	return dara.Validate(s)
}
