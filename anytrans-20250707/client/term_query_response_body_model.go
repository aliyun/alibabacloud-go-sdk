// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TermQueryResponseBody
	GetCode() *string
	SetData(v *TermQueryResponseBodyData) *TermQueryResponseBody
	GetData() *TermQueryResponseBodyData
	SetHttpStatusCode(v string) *TermQueryResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *TermQueryResponseBody
	GetMessage() *string
	SetRequestId(v string) *TermQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TermQueryResponseBody
	GetSuccess() *bool
}

type TermQueryResponseBody struct {
	// example:
	//
	// "success"
	Code *string                    `json:"code,omitempty" xml:"code,omitempty"`
	Data *TermQueryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// "success"
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// C3C70C8F-E026-17D8-854E-7F8EF2F5C909
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TermQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TermQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TermQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TermQueryResponseBody) GetData() *TermQueryResponseBodyData {
	return s.Data
}

func (s *TermQueryResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *TermQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TermQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TermQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TermQueryResponseBody) SetCode(v string) *TermQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TermQueryResponseBody) SetData(v *TermQueryResponseBodyData) *TermQueryResponseBody {
	s.Data = v
	return s
}

func (s *TermQueryResponseBody) SetHttpStatusCode(v string) *TermQueryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TermQueryResponseBody) SetMessage(v string) *TermQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TermQueryResponseBody) SetRequestId(v string) *TermQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TermQueryResponseBody) SetSuccess(v bool) *TermQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TermQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TermQueryResponseBodyData struct {
	// example:
	//
	// 0
	FailCount *int64                            `json:"failCount,omitempty" xml:"failCount,omitempty"`
	Terms     []*TermQueryResponseBodyDataTerms `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
}

func (s TermQueryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TermQueryResponseBodyData) GoString() string {
	return s.String()
}

func (s *TermQueryResponseBodyData) GetFailCount() *int64 {
	return s.FailCount
}

func (s *TermQueryResponseBodyData) GetTerms() []*TermQueryResponseBodyDataTerms {
	return s.Terms
}

func (s *TermQueryResponseBodyData) SetFailCount(v int64) *TermQueryResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *TermQueryResponseBodyData) SetTerms(v []*TermQueryResponseBodyDataTerms) *TermQueryResponseBodyData {
	s.Terms = v
	return s
}

func (s *TermQueryResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type TermQueryResponseBodyDataTerms struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// 92669963
	TermId *string `json:"termId,omitempty" xml:"termId,omitempty"`
	// example:
	//
	// LLM
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s TermQueryResponseBodyDataTerms) String() string {
	return dara.Prettify(s)
}

func (s TermQueryResponseBodyDataTerms) GoString() string {
	return s.String()
}

func (s *TermQueryResponseBodyDataTerms) GetSrc() *string {
	return s.Src
}

func (s *TermQueryResponseBodyDataTerms) GetTermId() *string {
	return s.TermId
}

func (s *TermQueryResponseBodyDataTerms) GetTgt() *string {
	return s.Tgt
}

func (s *TermQueryResponseBodyDataTerms) SetSrc(v string) *TermQueryResponseBodyDataTerms {
	s.Src = &v
	return s
}

func (s *TermQueryResponseBodyDataTerms) SetTermId(v string) *TermQueryResponseBodyDataTerms {
	s.TermId = &v
	return s
}

func (s *TermQueryResponseBodyDataTerms) SetTgt(v string) *TermQueryResponseBodyDataTerms {
	s.Tgt = &v
	return s
}

func (s *TermQueryResponseBodyDataTerms) Validate() error {
	return dara.Validate(s)
}
