// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceMapInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAscendingSequence(v bool) *GetSourceMapInfoRequest
	GetAscendingSequence() *bool
	SetEdition(v string) *GetSourceMapInfoRequest
	GetEdition() *string
	SetID(v string) *GetSourceMapInfoRequest
	GetID() *string
	SetKeyword(v string) *GetSourceMapInfoRequest
	GetKeyword() *string
	SetOrderField(v string) *GetSourceMapInfoRequest
	GetOrderField() *string
	SetRegionId(v string) *GetSourceMapInfoRequest
	GetRegionId() *string
}

type GetSourceMapInfoRequest struct {
	// The order in which the files are sorted. Valid values:
	//
	// 	- true: ascending order
	//
	// 	- false: descending order
	//
	// example:
	//
	// true
	AscendingSequence *bool `json:"AscendingSequence,omitempty" xml:"AscendingSequence,omitempty"`
	// The version of the SourceMap file.
	//
	// example:
	//
	// 0.0.0
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The ID of the SourceMap file.
	//
	// This parameter is required.
	ID *string `json:"ID,omitempty" xml:"ID,omitempty"`
	// The keyword in the file name. The files are searched by keyword.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The criterion by which the files are sorted. Valid values:
	//
	// 	- version: The files are sorted by version.
	//
	// 	- uploadTime: The files are sorted by upload time.
	//
	// example:
	//
	// version
	OrderField *string `json:"OrderField,omitempty" xml:"OrderField,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSourceMapInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSourceMapInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSourceMapInfoRequest) GetAscendingSequence() *bool {
	return s.AscendingSequence
}

func (s *GetSourceMapInfoRequest) GetEdition() *string {
	return s.Edition
}

func (s *GetSourceMapInfoRequest) GetID() *string {
	return s.ID
}

func (s *GetSourceMapInfoRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *GetSourceMapInfoRequest) GetOrderField() *string {
	return s.OrderField
}

func (s *GetSourceMapInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSourceMapInfoRequest) SetAscendingSequence(v bool) *GetSourceMapInfoRequest {
	s.AscendingSequence = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetEdition(v string) *GetSourceMapInfoRequest {
	s.Edition = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetID(v string) *GetSourceMapInfoRequest {
	s.ID = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetKeyword(v string) *GetSourceMapInfoRequest {
	s.Keyword = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetOrderField(v string) *GetSourceMapInfoRequest {
	s.OrderField = &v
	return s
}

func (s *GetSourceMapInfoRequest) SetRegionId(v string) *GetSourceMapInfoRequest {
	s.RegionId = &v
	return s
}

func (s *GetSourceMapInfoRequest) Validate() error {
	return dara.Validate(s)
}
