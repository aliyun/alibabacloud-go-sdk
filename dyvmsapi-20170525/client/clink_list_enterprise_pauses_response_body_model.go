// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListEnterprisePausesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListEnterprisePausesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListEnterprisePausesResponseBody
	GetCode() *string
	SetData(v *ClinkListEnterprisePausesResponseBodyData) *ClinkListEnterprisePausesResponseBody
	GetData() *ClinkListEnterprisePausesResponseBodyData
	SetMessage(v string) *ClinkListEnterprisePausesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListEnterprisePausesResponseBody
	GetRequestId() *string
}

type ClinkListEnterprisePausesResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListEnterprisePausesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListEnterprisePausesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListEnterprisePausesResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListEnterprisePausesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListEnterprisePausesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListEnterprisePausesResponseBody) GetData() *ClinkListEnterprisePausesResponseBodyData {
	return s.Data
}

func (s *ClinkListEnterprisePausesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListEnterprisePausesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListEnterprisePausesResponseBody) SetAccessDeniedDetail(v string) *ClinkListEnterprisePausesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBody) SetCode(v string) *ClinkListEnterprisePausesResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBody) SetData(v *ClinkListEnterprisePausesResponseBodyData) *ClinkListEnterprisePausesResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListEnterprisePausesResponseBody) SetMessage(v string) *ClinkListEnterprisePausesResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBody) SetRequestId(v string) *ClinkListEnterprisePausesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListEnterprisePausesResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 企业置忙状态列表
	EnterprisePauses []*ClinkListEnterprisePausesResponseBodyDataEnterprisePauses `json:"EnterprisePauses,omitempty" xml:"EnterprisePauses,omitempty" type:"Repeated"`
}

func (s ClinkListEnterprisePausesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListEnterprisePausesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListEnterprisePausesResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListEnterprisePausesResponseBodyData) GetEnterprisePauses() []*ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	return s.EnterprisePauses
}

func (s *ClinkListEnterprisePausesResponseBodyData) SetClinkRequestId(v string) *ClinkListEnterprisePausesResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyData) SetEnterprisePauses(v []*ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) *ClinkListEnterprisePausesResponseBodyData {
	s.EnterprisePauses = v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyData) Validate() error {
	if s.EnterprisePauses != nil {
		for _, item := range s.EnterprisePauses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListEnterprisePausesResponseBodyDataEnterprisePauses struct {
	// 置忙状态Id
	//
	// example:
	//
	// 14
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 默认状态，0：不是；1：是
	//
	// example:
	//
	// 0
	IsDefault *int64 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 休息状态，0：不是；1：是
	//
	// example:
	//
	// 0
	IsRest *string `json:"IsRest,omitempty" xml:"IsRest,omitempty"`
	// 置忙状态描述
	//
	// example:
	//
	// 示例值
	PauseStatus *string `json:"PauseStatus,omitempty" xml:"PauseStatus,omitempty"`
	// 优先级，数字越小优先级越高
	//
	// example:
	//
	// 8
	Sort *int64 `json:"Sort,omitempty" xml:"Sort,omitempty"`
}

func (s ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) String() string {
	return dara.Prettify(s)
}

func (s ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GoString() string {
	return s.String()
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GetId() *int64 {
	return s.Id
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GetIsDefault() *int64 {
	return s.IsDefault
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GetIsRest() *string {
	return s.IsRest
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GetPauseStatus() *string {
	return s.PauseStatus
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) GetSort() *int64 {
	return s.Sort
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) SetId(v int64) *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	s.Id = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) SetIsDefault(v int64) *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	s.IsDefault = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) SetIsRest(v string) *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	s.IsRest = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) SetPauseStatus(v string) *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	s.PauseStatus = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) SetSort(v int64) *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses {
	s.Sort = &v
	return s
}

func (s *ClinkListEnterprisePausesResponseBodyDataEnterprisePauses) Validate() error {
	return dara.Validate(s)
}
