// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeywordLibsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListKeywordLibsResponseBody
	GetCode() *int32
	SetData(v []*ListKeywordLibsResponseBodyData) *ListKeywordLibsResponseBody
	GetData() []*ListKeywordLibsResponseBodyData
	SetMsg(v string) *ListKeywordLibsResponseBody
	GetMsg() *string
	SetRequestId(v string) *ListKeywordLibsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKeywordLibsResponseBody
	GetSuccess() *bool
}

type ListKeywordLibsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*ListKeywordLibsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKeywordLibsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordLibsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListKeywordLibsResponseBody) GetData() []*ListKeywordLibsResponseBodyData {
	return s.Data
}

func (s *ListKeywordLibsResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ListKeywordLibsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeywordLibsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKeywordLibsResponseBody) SetCode(v int32) *ListKeywordLibsResponseBody {
	s.Code = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetData(v []*ListKeywordLibsResponseBodyData) *ListKeywordLibsResponseBody {
	s.Data = v
	return s
}

func (s *ListKeywordLibsResponseBody) SetMsg(v string) *ListKeywordLibsResponseBody {
	s.Msg = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetRequestId(v string) *ListKeywordLibsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeywordLibsResponseBody) SetSuccess(v bool) *ListKeywordLibsResponseBody {
	s.Success = &v
	return s
}

func (s *ListKeywordLibsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListKeywordLibsResponseBodyData struct {
	// example:
	//
	// 2022-11-30 16:30:29
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 10
	KeywordCount *string `json:"KeywordCount,omitempty" xml:"KeywordCount,omitempty"`
	// example:
	//
	// custom_xxxxx
	LibId        *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	LibName      *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	ServiceCodes *string `json:"ServiceCodes,omitempty" xml:"ServiceCodes,omitempty"`
	// UID。
	//
	// example:
	//
	// 19964*****086772
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s ListKeywordLibsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKeywordLibsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKeywordLibsResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListKeywordLibsResponseBodyData) GetKeywordCount() *string {
	return s.KeywordCount
}

func (s *ListKeywordLibsResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *ListKeywordLibsResponseBodyData) GetLibName() *string {
	return s.LibName
}

func (s *ListKeywordLibsResponseBodyData) GetServiceCodes() *string {
	return s.ServiceCodes
}

func (s *ListKeywordLibsResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *ListKeywordLibsResponseBodyData) SetGmtModified(v string) *ListKeywordLibsResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetKeywordCount(v string) *ListKeywordLibsResponseBodyData {
	s.KeywordCount = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetLibId(v string) *ListKeywordLibsResponseBodyData {
	s.LibId = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetLibName(v string) *ListKeywordLibsResponseBodyData {
	s.LibName = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetServiceCodes(v string) *ListKeywordLibsResponseBodyData {
	s.ServiceCodes = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) SetUid(v string) *ListKeywordLibsResponseBodyData {
	s.Uid = &v
	return s
}

func (s *ListKeywordLibsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
