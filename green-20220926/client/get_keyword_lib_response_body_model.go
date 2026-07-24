// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKeywordLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetKeywordLibResponseBody
	GetCode() *int32
	SetData(v *GetKeywordLibResponseBodyData) *GetKeywordLibResponseBody
	GetData() *GetKeywordLibResponseBodyData
	SetMsg(v string) *GetKeywordLibResponseBody
	GetMsg() *string
	SetRequestId(v string) *GetKeywordLibResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKeywordLibResponseBody
	GetSuccess() *bool
}

type GetKeywordLibResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *GetKeywordLibResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The description of the error code.
	//
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID assigned by the backend, which uniquely identifies the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKeywordLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordLibResponseBody) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetKeywordLibResponseBody) GetData() *GetKeywordLibResponseBodyData {
	return s.Data
}

func (s *GetKeywordLibResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *GetKeywordLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKeywordLibResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKeywordLibResponseBody) SetCode(v int32) *GetKeywordLibResponseBody {
	s.Code = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetData(v *GetKeywordLibResponseBodyData) *GetKeywordLibResponseBody {
	s.Data = v
	return s
}

func (s *GetKeywordLibResponseBody) SetMsg(v string) *GetKeywordLibResponseBody {
	s.Msg = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetRequestId(v string) *GetKeywordLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKeywordLibResponseBody) SetSuccess(v bool) *GetKeywordLibResponseBody {
	s.Success = &v
	return s
}

func (s *GetKeywordLibResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKeywordLibResponseBodyData struct {
	// The last modification time.
	//
	// example:
	//
	// 2024-01-29 10:26:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The number of keywords.
	//
	// example:
	//
	// 100
	KeywordCount *string `json:"KeywordCount,omitempty" xml:"KeywordCount,omitempty"`
	// The keyword library ID.
	//
	// example:
	//
	// customxx_xxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The library name.
	//
	// example:
	//
	// TestLibrary.
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 1825457112123838
	Uid *string `json:"Uid,omitempty" xml:"Uid,omitempty"`
}

func (s GetKeywordLibResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKeywordLibResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKeywordLibResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetKeywordLibResponseBodyData) GetKeywordCount() *string {
	return s.KeywordCount
}

func (s *GetKeywordLibResponseBodyData) GetLibId() *string {
	return s.LibId
}

func (s *GetKeywordLibResponseBodyData) GetLibName() *string {
	return s.LibName
}

func (s *GetKeywordLibResponseBodyData) GetUid() *string {
	return s.Uid
}

func (s *GetKeywordLibResponseBodyData) SetGmtModified(v string) *GetKeywordLibResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetKeywordCount(v string) *GetKeywordLibResponseBodyData {
	s.KeywordCount = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetLibId(v string) *GetKeywordLibResponseBodyData {
	s.LibId = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetLibName(v string) *GetKeywordLibResponseBodyData {
	s.LibName = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) SetUid(v string) *GetKeywordLibResponseBodyData {
	s.Uid = &v
	return s
}

func (s *GetKeywordLibResponseBodyData) Validate() error {
	return dara.Validate(s)
}
