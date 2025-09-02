// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgWhiteListQueryListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataType(v string) *DsgWhiteListQueryListRequest
	GetDataType() *string
	SetPageNumber(v int32) *DsgWhiteListQueryListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DsgWhiteListQueryListRequest
	GetPageSize() *int32
	SetSceneId(v int64) *DsgWhiteListQueryListRequest
	GetSceneId() *int64
}

type DsgWhiteListQueryListRequest struct {
	// The keyword of the sensitive field type.
	//
	// example:
	//
	// phone
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the level-2 data masking scenario. You can call the [DsgSceneQuerySceneListByName](https://help.aliyun.com/document_detail/2786322.html) operation to query the list of IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	SceneId *int64 `json:"SceneId,omitempty" xml:"SceneId,omitempty"`
}

func (s DsgWhiteListQueryListRequest) String() string {
	return dara.Prettify(s)
}

func (s DsgWhiteListQueryListRequest) GoString() string {
	return s.String()
}

func (s *DsgWhiteListQueryListRequest) GetDataType() *string {
	return s.DataType
}

func (s *DsgWhiteListQueryListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DsgWhiteListQueryListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DsgWhiteListQueryListRequest) GetSceneId() *int64 {
	return s.SceneId
}

func (s *DsgWhiteListQueryListRequest) SetDataType(v string) *DsgWhiteListQueryListRequest {
	s.DataType = &v
	return s
}

func (s *DsgWhiteListQueryListRequest) SetPageNumber(v int32) *DsgWhiteListQueryListRequest {
	s.PageNumber = &v
	return s
}

func (s *DsgWhiteListQueryListRequest) SetPageSize(v int32) *DsgWhiteListQueryListRequest {
	s.PageSize = &v
	return s
}

func (s *DsgWhiteListQueryListRequest) SetSceneId(v int64) *DsgWhiteListQueryListRequest {
	s.SceneId = &v
	return s
}

func (s *DsgWhiteListQueryListRequest) Validate() error {
	return dara.Validate(s)
}
