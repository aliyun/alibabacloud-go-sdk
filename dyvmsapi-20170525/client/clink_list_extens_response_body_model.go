// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListExtensResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListExtensResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListExtensResponseBody
	GetCode() *string
	SetData(v *ClinkListExtensResponseBodyData) *ClinkListExtensResponseBody
	GetData() *ClinkListExtensResponseBodyData
	SetMessage(v string) *ClinkListExtensResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListExtensResponseBody
	GetRequestId() *string
}

type ClinkListExtensResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListExtensResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListExtensResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListExtensResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListExtensResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListExtensResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListExtensResponseBody) GetData() *ClinkListExtensResponseBodyData {
	return s.Data
}

func (s *ClinkListExtensResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListExtensResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListExtensResponseBody) SetAccessDeniedDetail(v string) *ClinkListExtensResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListExtensResponseBody) SetCode(v string) *ClinkListExtensResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListExtensResponseBody) SetData(v *ClinkListExtensResponseBodyData) *ClinkListExtensResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListExtensResponseBody) SetMessage(v string) *ClinkListExtensResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListExtensResponseBody) SetRequestId(v string) *ClinkListExtensResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListExtensResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListExtensResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 话机对象列表
	Extens []*ClinkListExtensResponseBodyDataExtens `json:"Extens,omitempty" xml:"Extens,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 38
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListExtensResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListExtensResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListExtensResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListExtensResponseBodyData) GetExtens() []*ClinkListExtensResponseBodyDataExtens {
	return s.Extens
}

func (s *ClinkListExtensResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListExtensResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListExtensResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListExtensResponseBodyData) SetClinkRequestId(v string) *ClinkListExtensResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListExtensResponseBodyData) SetExtens(v []*ClinkListExtensResponseBodyDataExtens) *ClinkListExtensResponseBodyData {
	s.Extens = v
	return s
}

func (s *ClinkListExtensResponseBodyData) SetPageNumber(v int64) *ClinkListExtensResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListExtensResponseBodyData) SetPageSize(v int64) *ClinkListExtensResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListExtensResponseBodyData) SetTotalCount(v int64) *ClinkListExtensResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListExtensResponseBodyData) Validate() error {
	if s.Extens != nil {
		for _, item := range s.Extens {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListExtensResponseBodyDataExtens struct {
	// 语音编码。1：g729(已弃用)；2：g729,alaw,ulaw；3：alaw,ulaw,g729；4：myopus,alaw,ulaw,g729；5：alaw,ulaw；6：myopus,alaw,ulaw
	//
	// example:
	//
	// 2
	Allow *int64 `json:"Allow,omitempty" xml:"Allow,omitempty"`
	// 话机区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 话机号码
	//
	// example:
	//
	// xxx
	ExtenNumber *string `json:"ExtenNumber,omitempty" xml:"ExtenNumber,omitempty"`
	// 是否允许主叫外呼，1：允许；0：不允许。话机类型为IP话机时，可以开启主叫外呼功能，直接通过IP话机进行外呼。若要使用主叫外呼功能，需要开启企业级开关。
	//
	// example:
	//
	// 1
	IsDirect *int64 `json:"IsDirect,omitempty" xml:"IsDirect,omitempty"`
	// 网络防抖开关，0：关闭；1：开启
	//
	// example:
	//
	// 0
	JittBuffer *int64 `json:"JittBuffer,omitempty" xml:"JittBuffer,omitempty"`
	// 话机类型。1: IP话机， 2: 软电话
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkListExtensResponseBodyDataExtens) String() string {
	return dara.Prettify(s)
}

func (s ClinkListExtensResponseBodyDataExtens) GoString() string {
	return s.String()
}

func (s *ClinkListExtensResponseBodyDataExtens) GetAllow() *int64 {
	return s.Allow
}

func (s *ClinkListExtensResponseBodyDataExtens) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkListExtensResponseBodyDataExtens) GetExtenNumber() *string {
	return s.ExtenNumber
}

func (s *ClinkListExtensResponseBodyDataExtens) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *ClinkListExtensResponseBodyDataExtens) GetJittBuffer() *int64 {
	return s.JittBuffer
}

func (s *ClinkListExtensResponseBodyDataExtens) GetType() *int64 {
	return s.Type
}

func (s *ClinkListExtensResponseBodyDataExtens) SetAllow(v int64) *ClinkListExtensResponseBodyDataExtens {
	s.Allow = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) SetAreaCode(v string) *ClinkListExtensResponseBodyDataExtens {
	s.AreaCode = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) SetExtenNumber(v string) *ClinkListExtensResponseBodyDataExtens {
	s.ExtenNumber = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) SetIsDirect(v int64) *ClinkListExtensResponseBodyDataExtens {
	s.IsDirect = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) SetJittBuffer(v int64) *ClinkListExtensResponseBodyDataExtens {
	s.JittBuffer = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) SetType(v int64) *ClinkListExtensResponseBodyDataExtens {
	s.Type = &v
	return s
}

func (s *ClinkListExtensResponseBodyDataExtens) Validate() error {
	return dara.Validate(s)
}
