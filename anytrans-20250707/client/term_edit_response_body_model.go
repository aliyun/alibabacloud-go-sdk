// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTermEditResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TermEditResponseBody
	GetCode() *string
	SetData(v *TermEditResponseBodyData) *TermEditResponseBody
	GetData() *TermEditResponseBodyData
	SetHttpStatusCode(v string) *TermEditResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *TermEditResponseBody
	GetMessage() *string
	SetRequestId(v string) *TermEditResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TermEditResponseBody
	GetSuccess() *bool
}

type TermEditResponseBody struct {
	// example:
	//
	// "success"
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *TermEditResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// 1DCD50EC-D218-1844-9CD8-E97CAB9D31BE
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s TermEditResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TermEditResponseBody) GoString() string {
	return s.String()
}

func (s *TermEditResponseBody) GetCode() *string {
	return s.Code
}

func (s *TermEditResponseBody) GetData() *TermEditResponseBodyData {
	return s.Data
}

func (s *TermEditResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *TermEditResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TermEditResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TermEditResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TermEditResponseBody) SetCode(v string) *TermEditResponseBody {
	s.Code = &v
	return s
}

func (s *TermEditResponseBody) SetData(v *TermEditResponseBodyData) *TermEditResponseBody {
	s.Data = v
	return s
}

func (s *TermEditResponseBody) SetHttpStatusCode(v string) *TermEditResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *TermEditResponseBody) SetMessage(v string) *TermEditResponseBody {
	s.Message = &v
	return s
}

func (s *TermEditResponseBody) SetRequestId(v string) *TermEditResponseBody {
	s.RequestId = &v
	return s
}

func (s *TermEditResponseBody) SetSuccess(v bool) *TermEditResponseBody {
	s.Success = &v
	return s
}

func (s *TermEditResponseBody) Validate() error {
	return dara.Validate(s)
}

type TermEditResponseBodyData struct {
	// example:
	//
	// 0
	FailCount *int64                           `json:"failCount,omitempty" xml:"failCount,omitempty"`
	Terms     []*TermEditResponseBodyDataTerms `json:"terms,omitempty" xml:"terms,omitempty" type:"Repeated"`
}

func (s TermEditResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TermEditResponseBodyData) GoString() string {
	return s.String()
}

func (s *TermEditResponseBodyData) GetFailCount() *int64 {
	return s.FailCount
}

func (s *TermEditResponseBodyData) GetTerms() []*TermEditResponseBodyDataTerms {
	return s.Terms
}

func (s *TermEditResponseBodyData) SetFailCount(v int64) *TermEditResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *TermEditResponseBodyData) SetTerms(v []*TermEditResponseBodyDataTerms) *TermEditResponseBodyData {
	s.Terms = v
	return s
}

func (s *TermEditResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type TermEditResponseBodyDataTerms struct {
	Src *string `json:"src,omitempty" xml:"src,omitempty"`
	// example:
	//
	// 92669964
	TermId *string `json:"termId,omitempty" xml:"termId,omitempty"`
	// example:
	//
	// LLM
	Tgt *string `json:"tgt,omitempty" xml:"tgt,omitempty"`
}

func (s TermEditResponseBodyDataTerms) String() string {
	return dara.Prettify(s)
}

func (s TermEditResponseBodyDataTerms) GoString() string {
	return s.String()
}

func (s *TermEditResponseBodyDataTerms) GetSrc() *string {
	return s.Src
}

func (s *TermEditResponseBodyDataTerms) GetTermId() *string {
	return s.TermId
}

func (s *TermEditResponseBodyDataTerms) GetTgt() *string {
	return s.Tgt
}

func (s *TermEditResponseBodyDataTerms) SetSrc(v string) *TermEditResponseBodyDataTerms {
	s.Src = &v
	return s
}

func (s *TermEditResponseBodyDataTerms) SetTermId(v string) *TermEditResponseBodyDataTerms {
	s.TermId = &v
	return s
}

func (s *TermEditResponseBodyDataTerms) SetTgt(v string) *TermEditResponseBodyDataTerms {
	s.Tgt = &v
	return s
}

func (s *TermEditResponseBodyDataTerms) Validate() error {
	return dara.Validate(s)
}
