// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckNFCBindPopResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckNFCBindPopResponseBody
	GetAccessDeniedDetail() *string
	SetData(v *CheckNFCBindPopResponseBodyData) *CheckNFCBindPopResponseBody
	GetData() *CheckNFCBindPopResponseBodyData
	SetErrCode(v string) *CheckNFCBindPopResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CheckNFCBindPopResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *CheckNFCBindPopResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *CheckNFCBindPopResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckNFCBindPopResponseBody
	GetSuccess() *bool
}

type CheckNFCBindPopResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// data
	Data *CheckNFCBindPopResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// error
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckNFCBindPopResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckNFCBindPopResponseBody) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckNFCBindPopResponseBody) GetData() *CheckNFCBindPopResponseBodyData {
	return s.Data
}

func (s *CheckNFCBindPopResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CheckNFCBindPopResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CheckNFCBindPopResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CheckNFCBindPopResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckNFCBindPopResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckNFCBindPopResponseBody) SetAccessDeniedDetail(v string) *CheckNFCBindPopResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetData(v *CheckNFCBindPopResponseBodyData) *CheckNFCBindPopResponseBody {
	s.Data = v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetErrCode(v string) *CheckNFCBindPopResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetErrMessage(v string) *CheckNFCBindPopResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetHttpStatusCode(v int32) *CheckNFCBindPopResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetRequestId(v string) *CheckNFCBindPopResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) SetSuccess(v bool) *CheckNFCBindPopResponseBody {
	s.Success = &v
	return s
}

func (s *CheckNFCBindPopResponseBody) Validate() error {
	return dara.Validate(s)
}

type CheckNFCBindPopResponseBodyData struct {
	// example:
	//
	// 0
	IsGlobal *int32 `json:"IsGlobal,omitempty" xml:"IsGlobal,omitempty"`
	// example:
	//
	// true
	IsSign *bool `json:"IsSign,omitempty" xml:"IsSign,omitempty"`
}

func (s CheckNFCBindPopResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckNFCBindPopResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckNFCBindPopResponseBodyData) GetIsGlobal() *int32 {
	return s.IsGlobal
}

func (s *CheckNFCBindPopResponseBodyData) GetIsSign() *bool {
	return s.IsSign
}

func (s *CheckNFCBindPopResponseBodyData) SetIsGlobal(v int32) *CheckNFCBindPopResponseBodyData {
	s.IsGlobal = &v
	return s
}

func (s *CheckNFCBindPopResponseBodyData) SetIsSign(v bool) *CheckNFCBindPopResponseBodyData {
	s.IsSign = &v
	return s
}

func (s *CheckNFCBindPopResponseBodyData) Validate() error {
	return dara.Validate(s)
}
