// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRCSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendRCSResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendRCSResponseBody
	GetCode() *string
	SetData(v *SendRCSResponseBodyData) *SendRCSResponseBody
	GetData() *SendRCSResponseBodyData
	SetMessage(v string) *SendRCSResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendRCSResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendRCSResponseBody
	GetSuccess() *bool
}

type SendRCSResponseBody struct {
	AccessDeniedDetail *string                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *SendRCSResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendRCSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendRCSResponseBody) GoString() string {
	return s.String()
}

func (s *SendRCSResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendRCSResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendRCSResponseBody) GetData() *SendRCSResponseBodyData {
	return s.Data
}

func (s *SendRCSResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendRCSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRCSResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendRCSResponseBody) SetAccessDeniedDetail(v string) *SendRCSResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendRCSResponseBody) SetCode(v string) *SendRCSResponseBody {
	s.Code = &v
	return s
}

func (s *SendRCSResponseBody) SetData(v *SendRCSResponseBodyData) *SendRCSResponseBody {
	s.Data = v
	return s
}

func (s *SendRCSResponseBody) SetMessage(v string) *SendRCSResponseBody {
	s.Message = &v
	return s
}

func (s *SendRCSResponseBody) SetRequestId(v string) *SendRCSResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendRCSResponseBody) SetSuccess(v bool) *SendRCSResponseBody {
	s.Success = &v
	return s
}

func (s *SendRCSResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendRCSResponseBodyData struct {
	// example:
	//
	// 示例值示例值
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Bdcust *string `json:"Bdcust,omitempty" xml:"Bdcust,omitempty"`
	// example:
	//
	// 示例值示例值
	Code  *string                `json:"Code,omitempty" xml:"Code,omitempty"`
	Debug map[string]interface{} `json:"Debug,omitempty" xml:"Debug,omitempty"`
	// example:
	//
	// 示例值示例值
	E         *string                `json:"E,omitempty" xml:"E,omitempty"`
	ExtendMap map[string]interface{} `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	// example:
	//
	// 示例值示例值
	GateFailMsg *string `json:"GateFailMsg,omitempty" xml:"GateFailMsg,omitempty"`
	// example:
	//
	// 示例值示例值
	KeyString *string `json:"KeyString,omitempty" xml:"KeyString,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string                `json:"Message,omitempty" xml:"Message,omitempty"`
	Module  map[string]interface{} `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 示例值示例值
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendRCSResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendRCSResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendRCSResponseBodyData) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendRCSResponseBodyData) GetBdcust() *string {
	return s.Bdcust
}

func (s *SendRCSResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *SendRCSResponseBodyData) GetDebug() map[string]interface{} {
	return s.Debug
}

func (s *SendRCSResponseBodyData) GetE() *string {
	return s.E
}

func (s *SendRCSResponseBodyData) GetExtendMap() map[string]interface{} {
	return s.ExtendMap
}

func (s *SendRCSResponseBodyData) GetGateFailMsg() *string {
	return s.GateFailMsg
}

func (s *SendRCSResponseBodyData) GetKeyString() *string {
	return s.KeyString
}

func (s *SendRCSResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SendRCSResponseBodyData) GetModule() map[string]interface{} {
	return s.Module
}

func (s *SendRCSResponseBodyData) GetPartnerId() *string {
	return s.PartnerId
}

func (s *SendRCSResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRCSResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SendRCSResponseBodyData) SetAccessDeniedDetail(v string) *SendRCSResponseBodyData {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendRCSResponseBodyData) SetBdcust(v string) *SendRCSResponseBodyData {
	s.Bdcust = &v
	return s
}

func (s *SendRCSResponseBodyData) SetCode(v string) *SendRCSResponseBodyData {
	s.Code = &v
	return s
}

func (s *SendRCSResponseBodyData) SetDebug(v map[string]interface{}) *SendRCSResponseBodyData {
	s.Debug = v
	return s
}

func (s *SendRCSResponseBodyData) SetE(v string) *SendRCSResponseBodyData {
	s.E = &v
	return s
}

func (s *SendRCSResponseBodyData) SetExtendMap(v map[string]interface{}) *SendRCSResponseBodyData {
	s.ExtendMap = v
	return s
}

func (s *SendRCSResponseBodyData) SetGateFailMsg(v string) *SendRCSResponseBodyData {
	s.GateFailMsg = &v
	return s
}

func (s *SendRCSResponseBodyData) SetKeyString(v string) *SendRCSResponseBodyData {
	s.KeyString = &v
	return s
}

func (s *SendRCSResponseBodyData) SetMessage(v string) *SendRCSResponseBodyData {
	s.Message = &v
	return s
}

func (s *SendRCSResponseBodyData) SetModule(v map[string]interface{}) *SendRCSResponseBodyData {
	s.Module = v
	return s
}

func (s *SendRCSResponseBodyData) SetPartnerId(v string) *SendRCSResponseBodyData {
	s.PartnerId = &v
	return s
}

func (s *SendRCSResponseBodyData) SetRequestId(v string) *SendRCSResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *SendRCSResponseBodyData) SetSuccess(v bool) *SendRCSResponseBodyData {
	s.Success = &v
	return s
}

func (s *SendRCSResponseBodyData) Validate() error {
	return dara.Validate(s)
}
