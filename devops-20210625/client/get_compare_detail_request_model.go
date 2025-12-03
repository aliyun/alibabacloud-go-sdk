// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCompareDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *GetCompareDetailRequest
	GetFrom() *string
	SetMaxDiffByte(v int32) *GetCompareDetailRequest
	GetMaxDiffByte() *int32
	SetMaxDiffFile(v int32) *GetCompareDetailRequest
	GetMaxDiffFile() *int32
	SetMergeBase(v bool) *GetCompareDetailRequest
	GetMergeBase() *bool
	SetOrganizationId(v string) *GetCompareDetailRequest
	GetOrganizationId() *string
	SetTo(v string) *GetCompareDetailRequest
	GetTo() *string
}

type GetCompareDetailRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c9fb781f3d66ef6ee60bdd5c414f5106454b1426
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	// example:
	//
	// 24 	- 1024 	- 1024
	MaxDiffByte *int32 `json:"maxDiffByte,omitempty" xml:"maxDiffByte,omitempty"`
	// example:
	//
	// 5000
	MaxDiffFile *int32 `json:"maxDiffFile,omitempty" xml:"maxDiffFile,omitempty"`
	// example:
	//
	// false
	MergeBase *bool `json:"mergeBase,omitempty" xml:"mergeBase,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5ebbc0228123212b59xxxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// b8f6f28520b1936aafe2e638373e19ccafa42b02
	To *string `json:"to,omitempty" xml:"to,omitempty"`
}

func (s GetCompareDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCompareDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCompareDetailRequest) GetFrom() *string {
	return s.From
}

func (s *GetCompareDetailRequest) GetMaxDiffByte() *int32 {
	return s.MaxDiffByte
}

func (s *GetCompareDetailRequest) GetMaxDiffFile() *int32 {
	return s.MaxDiffFile
}

func (s *GetCompareDetailRequest) GetMergeBase() *bool {
	return s.MergeBase
}

func (s *GetCompareDetailRequest) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *GetCompareDetailRequest) GetTo() *string {
	return s.To
}

func (s *GetCompareDetailRequest) SetFrom(v string) *GetCompareDetailRequest {
	s.From = &v
	return s
}

func (s *GetCompareDetailRequest) SetMaxDiffByte(v int32) *GetCompareDetailRequest {
	s.MaxDiffByte = &v
	return s
}

func (s *GetCompareDetailRequest) SetMaxDiffFile(v int32) *GetCompareDetailRequest {
	s.MaxDiffFile = &v
	return s
}

func (s *GetCompareDetailRequest) SetMergeBase(v bool) *GetCompareDetailRequest {
	s.MergeBase = &v
	return s
}

func (s *GetCompareDetailRequest) SetOrganizationId(v string) *GetCompareDetailRequest {
	s.OrganizationId = &v
	return s
}

func (s *GetCompareDetailRequest) SetTo(v string) *GetCompareDetailRequest {
	s.To = &v
	return s
}

func (s *GetCompareDetailRequest) Validate() error {
	return dara.Validate(s)
}
