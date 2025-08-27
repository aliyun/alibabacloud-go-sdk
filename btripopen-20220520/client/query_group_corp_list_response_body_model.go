// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGroupCorpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryGroupCorpListResponseBody
	GetRequestId() *string
	SetCode(v int32) *QueryGroupCorpListResponseBody
	GetCode() *int32
	SetMessage(v string) *QueryGroupCorpListResponseBody
	GetMessage() *string
	SetModule(v []*QueryGroupCorpListResponseBodyModule) *QueryGroupCorpListResponseBody
	GetModule() []*QueryGroupCorpListResponseBodyModule
	SetSuccess(v bool) *QueryGroupCorpListResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryGroupCorpListResponseBody
	GetTraceId() *string
}

type QueryGroupCorpListResponseBody struct {
	// example:
	//
	// A3F8DCAB-8301-5770-BD9F-71B0BF9E1A6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 0
	Code    *int32                                  `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                 `json:"message,omitempty" xml:"message,omitempty"`
	Module  []*QueryGroupCorpListResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"trace_id,omitempty" xml:"trace_id,omitempty"`
}

func (s QueryGroupCorpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupCorpListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryGroupCorpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryGroupCorpListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryGroupCorpListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryGroupCorpListResponseBody) GetModule() []*QueryGroupCorpListResponseBodyModule {
	return s.Module
}

func (s *QueryGroupCorpListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryGroupCorpListResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryGroupCorpListResponseBody) SetRequestId(v string) *QueryGroupCorpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryGroupCorpListResponseBody) SetCode(v int32) *QueryGroupCorpListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryGroupCorpListResponseBody) SetMessage(v string) *QueryGroupCorpListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryGroupCorpListResponseBody) SetModule(v []*QueryGroupCorpListResponseBodyModule) *QueryGroupCorpListResponseBody {
	s.Module = v
	return s
}

func (s *QueryGroupCorpListResponseBody) SetSuccess(v bool) *QueryGroupCorpListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryGroupCorpListResponseBody) SetTraceId(v string) *QueryGroupCorpListResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryGroupCorpListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryGroupCorpListResponseBodyModule struct {
	// example:
	//
	// corp1
	CorpId   *string `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CorpName *string `json:"corp_name,omitempty" xml:"corp_name,omitempty"`
}

func (s QueryGroupCorpListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryGroupCorpListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryGroupCorpListResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *QueryGroupCorpListResponseBodyModule) GetCorpName() *string {
	return s.CorpName
}

func (s *QueryGroupCorpListResponseBodyModule) SetCorpId(v string) *QueryGroupCorpListResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *QueryGroupCorpListResponseBodyModule) SetCorpName(v string) *QueryGroupCorpListResponseBodyModule {
	s.CorpName = &v
	return s
}

func (s *QueryGroupCorpListResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
