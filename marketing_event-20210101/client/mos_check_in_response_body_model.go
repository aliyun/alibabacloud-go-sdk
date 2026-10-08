// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMosCheckInResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *MosCheckInResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *MosCheckInResponseBody
	GetCode() *string
	SetData(v interface{}) *MosCheckInResponseBody
	GetData() interface{}
	SetMessage(v string) *MosCheckInResponseBody
	GetMessage() *string
	SetRequestId(v string) *MosCheckInResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MosCheckInResponseBody
	GetSuccess() *bool
}

type MosCheckInResponseBody struct {
	// example:
	//
	// deny
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// XXX
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// data
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// XXX
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1skladklasmda
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MosCheckInResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MosCheckInResponseBody) GoString() string {
	return s.String()
}

func (s *MosCheckInResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *MosCheckInResponseBody) GetCode() *string {
	return s.Code
}

func (s *MosCheckInResponseBody) GetData() interface{} {
	return s.Data
}

func (s *MosCheckInResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MosCheckInResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MosCheckInResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MosCheckInResponseBody) SetAccessDeniedDetail(v string) *MosCheckInResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *MosCheckInResponseBody) SetCode(v string) *MosCheckInResponseBody {
	s.Code = &v
	return s
}

func (s *MosCheckInResponseBody) SetData(v interface{}) *MosCheckInResponseBody {
	s.Data = v
	return s
}

func (s *MosCheckInResponseBody) SetMessage(v string) *MosCheckInResponseBody {
	s.Message = &v
	return s
}

func (s *MosCheckInResponseBody) SetRequestId(v string) *MosCheckInResponseBody {
	s.RequestId = &v
	return s
}

func (s *MosCheckInResponseBody) SetSuccess(v bool) *MosCheckInResponseBody {
	s.Success = &v
	return s
}

func (s *MosCheckInResponseBody) Validate() error {
	return dara.Validate(s)
}
