// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScriptHotWordListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotWord(v string) *QueryScriptHotWordListRequest
	GetHotWord() *string
	SetPageNumber(v int32) *QueryScriptHotWordListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *QueryScriptHotWordListRequest
	GetPageSize() *int32
	SetScriptRecordId(v string) *QueryScriptHotWordListRequest
	GetScriptRecordId() *string
}

type QueryScriptHotWordListRequest struct {
	HotWord        *string `json:"hotWord,omitempty" xml:"hotWord,omitempty"`
	PageNumber     *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize       *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ScriptRecordId *string `json:"scriptRecordId,omitempty" xml:"scriptRecordId,omitempty"`
}

func (s QueryScriptHotWordListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryScriptHotWordListRequest) GoString() string {
	return s.String()
}

func (s *QueryScriptHotWordListRequest) GetHotWord() *string {
	return s.HotWord
}

func (s *QueryScriptHotWordListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *QueryScriptHotWordListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryScriptHotWordListRequest) GetScriptRecordId() *string {
	return s.ScriptRecordId
}

func (s *QueryScriptHotWordListRequest) SetHotWord(v string) *QueryScriptHotWordListRequest {
	s.HotWord = &v
	return s
}

func (s *QueryScriptHotWordListRequest) SetPageNumber(v int32) *QueryScriptHotWordListRequest {
	s.PageNumber = &v
	return s
}

func (s *QueryScriptHotWordListRequest) SetPageSize(v int32) *QueryScriptHotWordListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryScriptHotWordListRequest) SetScriptRecordId(v string) *QueryScriptHotWordListRequest {
	s.ScriptRecordId = &v
	return s
}

func (s *QueryScriptHotWordListRequest) Validate() error {
	return dara.Validate(s)
}
