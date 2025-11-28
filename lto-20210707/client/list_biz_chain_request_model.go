// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizChainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizChainId(v string) *ListBizChainRequest
	GetBizChainId() *string
	SetName(v string) *ListBizChainRequest
	GetName() *string
	SetNum(v int32) *ListBizChainRequest
	GetNum() *int32
	SetRegionId(v string) *ListBizChainRequest
	GetRegionId() *string
	SetSize(v int32) *ListBizChainRequest
	GetSize() *int32
}

type ListBizChainRequest struct {
	BizChainId *string `json:"BizChainId,omitempty" xml:"BizChainId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Num      *int32  `json:"Num,omitempty" xml:"Num,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListBizChainRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBizChainRequest) GoString() string {
	return s.String()
}

func (s *ListBizChainRequest) GetBizChainId() *string {
	return s.BizChainId
}

func (s *ListBizChainRequest) GetName() *string {
	return s.Name
}

func (s *ListBizChainRequest) GetNum() *int32 {
	return s.Num
}

func (s *ListBizChainRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListBizChainRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListBizChainRequest) SetBizChainId(v string) *ListBizChainRequest {
	s.BizChainId = &v
	return s
}

func (s *ListBizChainRequest) SetName(v string) *ListBizChainRequest {
	s.Name = &v
	return s
}

func (s *ListBizChainRequest) SetNum(v int32) *ListBizChainRequest {
	s.Num = &v
	return s
}

func (s *ListBizChainRequest) SetRegionId(v string) *ListBizChainRequest {
	s.RegionId = &v
	return s
}

func (s *ListBizChainRequest) SetSize(v int32) *ListBizChainRequest {
	s.Size = &v
	return s
}

func (s *ListBizChainRequest) Validate() error {
	return dara.Validate(s)
}
