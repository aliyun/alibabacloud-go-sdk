// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkCreateExtenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkCreateExtenResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkCreateExtenResponseBody
	GetCode() *string
	SetData(v *ClinkCreateExtenResponseBodyData) *ClinkCreateExtenResponseBody
	GetData() *ClinkCreateExtenResponseBodyData
	SetMessage(v string) *ClinkCreateExtenResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkCreateExtenResponseBody
	GetRequestId() *string
}

type ClinkCreateExtenResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkCreateExtenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkCreateExtenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateExtenResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkCreateExtenResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkCreateExtenResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkCreateExtenResponseBody) GetData() *ClinkCreateExtenResponseBodyData {
	return s.Data
}

func (s *ClinkCreateExtenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkCreateExtenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkCreateExtenResponseBody) SetAccessDeniedDetail(v string) *ClinkCreateExtenResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkCreateExtenResponseBody) SetCode(v string) *ClinkCreateExtenResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkCreateExtenResponseBody) SetData(v *ClinkCreateExtenResponseBodyData) *ClinkCreateExtenResponseBody {
	s.Data = v
	return s
}

func (s *ClinkCreateExtenResponseBody) SetMessage(v string) *ClinkCreateExtenResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkCreateExtenResponseBody) SetRequestId(v string) *ClinkCreateExtenResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkCreateExtenResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateExtenResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// 示例值
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// 分机
	Exten *ClinkCreateExtenResponseBodyDataExten `json:"Exten,omitempty" xml:"Exten,omitempty" type:"Struct"`
}

func (s ClinkCreateExtenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateExtenResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkCreateExtenResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkCreateExtenResponseBodyData) GetExten() *ClinkCreateExtenResponseBodyDataExten {
	return s.Exten
}

func (s *ClinkCreateExtenResponseBodyData) SetClinkRequestId(v string) *ClinkCreateExtenResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyData) SetExten(v *ClinkCreateExtenResponseBodyDataExten) *ClinkCreateExtenResponseBodyData {
	s.Exten = v
	return s
}

func (s *ClinkCreateExtenResponseBodyData) Validate() error {
	if s.Exten != nil {
		if err := s.Exten.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkCreateExtenResponseBodyDataExten struct {
	// 语音编码。 1：g729(已弃用) 2：g729,alaw,ulaw 3：alaw,ulaw,g729 4：myopus,alaw,ulaw,g729 5：alaw,ulaw 6：myopus,alaw,ulaw
	//
	// example:
	//
	// 6
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
	// 11354
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
	// 话机密码，密码规则参见企业配置
	//
	// example:
	//
	// xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// 话机类型。1: IP话机， 2: 软电话
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ClinkCreateExtenResponseBodyDataExten) String() string {
	return dara.Prettify(s)
}

func (s ClinkCreateExtenResponseBodyDataExten) GoString() string {
	return s.String()
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetAllow() *int64 {
	return s.Allow
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetExtenNumber() *string {
	return s.ExtenNumber
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetIsDirect() *int64 {
	return s.IsDirect
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetJittBuffer() *int64 {
	return s.JittBuffer
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetPassword() *string {
	return s.Password
}

func (s *ClinkCreateExtenResponseBodyDataExten) GetType() *int64 {
	return s.Type
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetAllow(v int64) *ClinkCreateExtenResponseBodyDataExten {
	s.Allow = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetAreaCode(v string) *ClinkCreateExtenResponseBodyDataExten {
	s.AreaCode = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetExtenNumber(v string) *ClinkCreateExtenResponseBodyDataExten {
	s.ExtenNumber = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetIsDirect(v int64) *ClinkCreateExtenResponseBodyDataExten {
	s.IsDirect = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetJittBuffer(v int64) *ClinkCreateExtenResponseBodyDataExten {
	s.JittBuffer = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetPassword(v string) *ClinkCreateExtenResponseBodyDataExten {
	s.Password = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) SetType(v int64) *ClinkCreateExtenResponseBodyDataExten {
	s.Type = &v
	return s
}

func (s *ClinkCreateExtenResponseBodyDataExten) Validate() error {
	return dara.Validate(s)
}
