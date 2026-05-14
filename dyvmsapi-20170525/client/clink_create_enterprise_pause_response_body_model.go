// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateEnterprisePauseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkCreateEnterprisePauseResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkCreateEnterprisePauseResponseBody
	GetCode() *string
	SetData(v *ClinkCreateEnterprisePauseResponseBodyData) *ClinkCreateEnterprisePauseResponseBody
	GetData() *ClinkCreateEnterprisePauseResponseBodyData
	SetMessage(v string) *ClinkCreateEnterprisePauseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkCreateEnterprisePauseResponseBody
	GetRequestId() *string
}

type ClinkCreateEnterprisePauseResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkCreateEnterprisePauseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkCreateEnterprisePauseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateEnterprisePauseResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkCreateEnterprisePauseResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkCreateEnterprisePauseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkCreateEnterprisePauseResponseBody) GetData() *ClinkCreateEnterprisePauseResponseBodyData {
	return s.Data
}

func (s *ClinkCreateEnterprisePauseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkCreateEnterprisePauseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkCreateEnterprisePauseResponseBody) SetAccessDeniedDetail(v string) *ClinkCreateEnterprisePauseResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBody) SetCode(v string) *ClinkCreateEnterprisePauseResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBody) SetData(v *ClinkCreateEnterprisePauseResponseBodyData) *ClinkCreateEnterprisePauseResponseBody {
	s.Data = v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBody) SetMessage(v string) *ClinkCreateEnterprisePauseResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBody) SetRequestId(v string) *ClinkCreateEnterprisePauseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateEnterprisePauseResponseBodyData struct {
	// 请求标识
	//
	// example:
	//
	// 示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 企业置忙状态
	EnterprisePause *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause `json:"EnterprisePause,omitempty" xml:"EnterprisePause,omitempty" type:"Struct"`
}

func (s ClinkCreateEnterprisePauseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateEnterprisePauseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkCreateEnterprisePauseResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkCreateEnterprisePauseResponseBodyData) GetEnterprisePause() *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause {
	return s.EnterprisePause
}

func (s *ClinkCreateEnterprisePauseResponseBodyData) SetClinkRequestId(v string) *ClinkCreateEnterprisePauseResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBodyData) SetEnterprisePause(v *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) *ClinkCreateEnterprisePauseResponseBodyData {
	s.EnterprisePause = v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBodyData) Validate() error {
	if s.EnterprisePause != nil {
		if err := s.EnterprisePause.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause struct {
	// 默认状态，0：不是；1：是
	//
	// example:
	//
	// 1
	IsDefault *int64 `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// 休息状态，0：不是；1：是
	//
	// example:
	//
	// 1
	IsRest *string `json:"IsRest,omitempty" xml:"IsRest,omitempty"`
	// 置忙状态描述（不超4个字符）
	//
	// example:
	//
	// xxx
	PauseStatus *string `json:"PauseStatus,omitempty" xml:"PauseStatus,omitempty"`
}

func (s ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) GoString() string {
	return s.String()
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) GetIsDefault() *int64 {
	return s.IsDefault
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) GetIsRest() *string {
	return s.IsRest
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) GetPauseStatus() *string {
	return s.PauseStatus
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) SetIsDefault(v int64) *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause {
	s.IsDefault = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) SetIsRest(v string) *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause {
	s.IsRest = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) SetPauseStatus(v string) *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause {
	s.PauseStatus = &v
	return s
}

func (s *ClinkCreateEnterprisePauseResponseBodyDataEnterprisePause) Validate() error {
	return dara.Validate(s)
}
