// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogContentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *ListLogContentsRequest
	GetFileName() *string
	SetLength(v int32) *ListLogContentsRequest
	GetLength() *int32
	SetOffset(v int32) *ListLogContentsRequest
	GetOffset() *int32
	SetRegionId(v string) *ListLogContentsRequest
	GetRegionId() *string
}

type ListLogContentsRequest struct {
	// Full path of the file.
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// Length of the log.
	//
	// example:
	//
	// 9999
	Length *int32 `json:"length,omitempty" xml:"length,omitempty"`
	// Start line for query.
	//
	// example:
	//
	// 0
	Offset *int32 `json:"offset,omitempty" xml:"offset,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s ListLogContentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogContentsRequest) GoString() string {
	return s.String()
}

func (s *ListLogContentsRequest) GetFileName() *string {
	return s.FileName
}

func (s *ListLogContentsRequest) GetLength() *int32 {
	return s.Length
}

func (s *ListLogContentsRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListLogContentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLogContentsRequest) SetFileName(v string) *ListLogContentsRequest {
	s.FileName = &v
	return s
}

func (s *ListLogContentsRequest) SetLength(v int32) *ListLogContentsRequest {
	s.Length = &v
	return s
}

func (s *ListLogContentsRequest) SetOffset(v int32) *ListLogContentsRequest {
	s.Offset = &v
	return s
}

func (s *ListLogContentsRequest) SetRegionId(v string) *ListLogContentsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLogContentsRequest) Validate() error {
	return dara.Validate(s)
}
