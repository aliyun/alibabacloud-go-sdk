// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRCSReplyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *SendRCSReplyResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *SendRCSReplyResponseBody
	GetCode() *string
	SetData(v *SendRCSReplyResponseBodyData) *SendRCSReplyResponseBody
	GetData() *SendRCSReplyResponseBodyData
	SetMessage(v string) *SendRCSReplyResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendRCSReplyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SendRCSReplyResponseBody
	GetSuccess() *bool
}

type SendRCSReplyResponseBody struct {
	AccessDeniedDetail *string                       `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	Code               *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data               *SendRCSReplyResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message            *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId          *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success            *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendRCSReplyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendRCSReplyResponseBody) GoString() string {
	return s.String()
}

func (s *SendRCSReplyResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendRCSReplyResponseBody) GetCode() *string {
	return s.Code
}

func (s *SendRCSReplyResponseBody) GetData() *SendRCSReplyResponseBodyData {
	return s.Data
}

func (s *SendRCSReplyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendRCSReplyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRCSReplyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SendRCSReplyResponseBody) SetAccessDeniedDetail(v string) *SendRCSReplyResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendRCSReplyResponseBody) SetCode(v string) *SendRCSReplyResponseBody {
	s.Code = &v
	return s
}

func (s *SendRCSReplyResponseBody) SetData(v *SendRCSReplyResponseBodyData) *SendRCSReplyResponseBody {
	s.Data = v
	return s
}

func (s *SendRCSReplyResponseBody) SetMessage(v string) *SendRCSReplyResponseBody {
	s.Message = &v
	return s
}

func (s *SendRCSReplyResponseBody) SetRequestId(v string) *SendRCSReplyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendRCSReplyResponseBody) SetSuccess(v bool) *SendRCSReplyResponseBody {
	s.Success = &v
	return s
}

func (s *SendRCSReplyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SendRCSReplyResponseBodyData struct {
	// example:
	//
	// 示例值
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
	// 示例值示例值示例值
	E         *string                `json:"E,omitempty" xml:"E,omitempty"`
	ExtendMap map[string]interface{} `json:"ExtendMap,omitempty" xml:"ExtendMap,omitempty"`
	// example:
	//
	// 示例值
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
	// 示例值示例值示例值
	PartnerId *string `json:"PartnerId,omitempty" xml:"PartnerId,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SendRCSReplyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SendRCSReplyResponseBodyData) GoString() string {
	return s.String()
}

func (s *SendRCSReplyResponseBodyData) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *SendRCSReplyResponseBodyData) GetBdcust() *string {
	return s.Bdcust
}

func (s *SendRCSReplyResponseBodyData) GetCode() *string {
	return s.Code
}

func (s *SendRCSReplyResponseBodyData) GetDebug() map[string]interface{} {
	return s.Debug
}

func (s *SendRCSReplyResponseBodyData) GetE() *string {
	return s.E
}

func (s *SendRCSReplyResponseBodyData) GetExtendMap() map[string]interface{} {
	return s.ExtendMap
}

func (s *SendRCSReplyResponseBodyData) GetGateFailMsg() *string {
	return s.GateFailMsg
}

func (s *SendRCSReplyResponseBodyData) GetKeyString() *string {
	return s.KeyString
}

func (s *SendRCSReplyResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *SendRCSReplyResponseBodyData) GetModule() map[string]interface{} {
	return s.Module
}

func (s *SendRCSReplyResponseBodyData) GetPartnerId() *string {
	return s.PartnerId
}

func (s *SendRCSReplyResponseBodyData) GetRequestId() *string {
	return s.RequestId
}

func (s *SendRCSReplyResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *SendRCSReplyResponseBodyData) SetAccessDeniedDetail(v string) *SendRCSReplyResponseBodyData {
	s.AccessDeniedDetail = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetBdcust(v string) *SendRCSReplyResponseBodyData {
	s.Bdcust = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetCode(v string) *SendRCSReplyResponseBodyData {
	s.Code = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetDebug(v map[string]interface{}) *SendRCSReplyResponseBodyData {
	s.Debug = v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetE(v string) *SendRCSReplyResponseBodyData {
	s.E = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetExtendMap(v map[string]interface{}) *SendRCSReplyResponseBodyData {
	s.ExtendMap = v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetGateFailMsg(v string) *SendRCSReplyResponseBodyData {
	s.GateFailMsg = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetKeyString(v string) *SendRCSReplyResponseBodyData {
	s.KeyString = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetMessage(v string) *SendRCSReplyResponseBodyData {
	s.Message = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetModule(v map[string]interface{}) *SendRCSReplyResponseBodyData {
	s.Module = v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetPartnerId(v string) *SendRCSReplyResponseBodyData {
	s.PartnerId = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetRequestId(v string) *SendRCSReplyResponseBodyData {
	s.RequestId = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) SetSuccess(v bool) *SendRCSReplyResponseBodyData {
	s.Success = &v
	return s
}

func (s *SendRCSReplyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
