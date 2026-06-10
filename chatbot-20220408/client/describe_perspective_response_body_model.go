// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePerspectiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribePerspectiveResponseBody
	GetCreateTime() *string
	SetModifyTime(v string) *DescribePerspectiveResponseBody
	GetModifyTime() *string
	SetName(v string) *DescribePerspectiveResponseBody
	GetName() *string
	SetPerspectiveCode(v string) *DescribePerspectiveResponseBody
	GetPerspectiveCode() *string
	SetPerspectiveId(v string) *DescribePerspectiveResponseBody
	GetPerspectiveId() *string
	SetRequestId(v string) *DescribePerspectiveResponseBody
	GetRequestId() *string
	SetSelfDefine(v bool) *DescribePerspectiveResponseBody
	GetSelfDefine() *bool
	SetStatus(v int32) *DescribePerspectiveResponseBody
	GetStatus() *int32
}

type DescribePerspectiveResponseBody struct {
	// The creation time in UTC.
	//
	// example:
	//
	// 2021-07-27T07:05:37Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The modification time in UTC.
	//
	// example:
	//
	// 2021-07-26T07:05:37Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the perspective.
	//
	// example:
	//
	// 移动端视角
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The perspective code (used for the Q\\&A API).
	//
	// example:
	//
	// FZJBY3raWr
	PerspectiveCode *string `json:"PerspectiveCode,omitempty" xml:"PerspectiveCode,omitempty"`
	// The ID of the perspective.
	//
	// example:
	//
	// 3001
	PerspectiveId *string `json:"PerspectiveId,omitempty" xml:"PerspectiveId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F285D735-D580-18A8-B97F-B2E72B00F101
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the perspective is user-defined.
	//
	// example:
	//
	// true
	SelfDefine *bool `json:"SelfDefine,omitempty" xml:"SelfDefine,omitempty"`
	// The status of the perspective. Valid values: 3 (Enabled) and 1 (Disabled).
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribePerspectiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePerspectiveResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribePerspectiveResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribePerspectiveResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribePerspectiveResponseBody) GetPerspectiveCode() *string {
	return s.PerspectiveCode
}

func (s *DescribePerspectiveResponseBody) GetPerspectiveId() *string {
	return s.PerspectiveId
}

func (s *DescribePerspectiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePerspectiveResponseBody) GetSelfDefine() *bool {
	return s.SelfDefine
}

func (s *DescribePerspectiveResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribePerspectiveResponseBody) SetCreateTime(v string) *DescribePerspectiveResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetModifyTime(v string) *DescribePerspectiveResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetName(v string) *DescribePerspectiveResponseBody {
	s.Name = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveCode(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveCode = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetPerspectiveId(v string) *DescribePerspectiveResponseBody {
	s.PerspectiveId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetRequestId(v string) *DescribePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetSelfDefine(v bool) *DescribePerspectiveResponseBody {
	s.SelfDefine = &v
	return s
}

func (s *DescribePerspectiveResponseBody) SetStatus(v int32) *DescribePerspectiveResponseBody {
	s.Status = &v
	return s
}

func (s *DescribePerspectiveResponseBody) Validate() error {
	return dara.Validate(s)
}
