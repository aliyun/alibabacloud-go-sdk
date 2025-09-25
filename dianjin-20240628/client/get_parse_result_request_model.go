// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetParseResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *GetParseResultRequest
	GetDocId() *string
	SetLibraryId(v string) *GetParseResultRequest
	GetLibraryId() *string
	SetUseUrlResult(v bool) *GetParseResultRequest
	GetUseUrlResult() *bool
}

type GetParseResultRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 873648346573245
	DocId *string `json:"docId,omitempty" xml:"docId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sjdgdsfg
	LibraryId    *string `json:"libraryId,omitempty" xml:"libraryId,omitempty"`
	UseUrlResult *bool   `json:"useUrlResult,omitempty" xml:"useUrlResult,omitempty"`
}

func (s GetParseResultRequest) String() string {
	return dara.Prettify(s)
}

func (s GetParseResultRequest) GoString() string {
	return s.String()
}

func (s *GetParseResultRequest) GetDocId() *string {
	return s.DocId
}

func (s *GetParseResultRequest) GetLibraryId() *string {
	return s.LibraryId
}

func (s *GetParseResultRequest) GetUseUrlResult() *bool {
	return s.UseUrlResult
}

func (s *GetParseResultRequest) SetDocId(v string) *GetParseResultRequest {
	s.DocId = &v
	return s
}

func (s *GetParseResultRequest) SetLibraryId(v string) *GetParseResultRequest {
	s.LibraryId = &v
	return s
}

func (s *GetParseResultRequest) SetUseUrlResult(v bool) *GetParseResultRequest {
	s.UseUrlResult = &v
	return s
}

func (s *GetParseResultRequest) Validate() error {
	return dara.Validate(s)
}
