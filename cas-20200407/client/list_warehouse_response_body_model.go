// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListWarehouseResponseBodyData) *ListWarehouseResponseBody
	GetData() []*ListWarehouseResponseBodyData
	SetMaxResults(v int32) *ListWarehouseResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListWarehouseResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListWarehouseResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListWarehouseResponseBody
	GetTotalCount() *int64
}

type ListWarehouseResponseBody struct {
	// A list of warehouse objects.
	Data []*ListWarehouseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The maximum number of entries returned on each page. The default value is 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token to retrieve the next page of results. If this parameter is not returned, all results have been retrieved.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries in the result set.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *ListWarehouseResponseBody) GetData() []*ListWarehouseResponseBodyData {
	return s.Data
}

func (s *ListWarehouseResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListWarehouseResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWarehouseResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListWarehouseResponseBody) SetData(v []*ListWarehouseResponseBodyData) *ListWarehouseResponseBody {
	s.Data = v
	return s
}

func (s *ListWarehouseResponseBody) SetMaxResults(v int32) *ListWarehouseResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWarehouseResponseBody) SetNextToken(v string) *ListWarehouseResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWarehouseResponseBody) SetRequestId(v string) *ListWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWarehouseResponseBody) SetTotalCount(v int64) *ListWarehouseResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWarehouseResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListWarehouseResponseBodyData struct {
	// The warehouse instance ID.
	//
	// example:
	//
	// cas-wh-Q7ID6V
	WarehouseInstanceId *string `json:"WarehouseInstanceId,omitempty" xml:"WarehouseInstanceId,omitempty"`
	// The warehouse name.
	//
	// example:
	//
	// default_warehouse
	WarehouseName *string `json:"WarehouseName,omitempty" xml:"WarehouseName,omitempty"`
	// The warehouse type.
	//
	// example:
	//
	// pcaCaCert
	WarehouseType *string `json:"WarehouseType,omitempty" xml:"WarehouseType,omitempty"`
}

func (s ListWarehouseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListWarehouseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListWarehouseResponseBodyData) GetWarehouseInstanceId() *string {
	return s.WarehouseInstanceId
}

func (s *ListWarehouseResponseBodyData) GetWarehouseName() *string {
	return s.WarehouseName
}

func (s *ListWarehouseResponseBodyData) GetWarehouseType() *string {
	return s.WarehouseType
}

func (s *ListWarehouseResponseBodyData) SetWarehouseInstanceId(v string) *ListWarehouseResponseBodyData {
	s.WarehouseInstanceId = &v
	return s
}

func (s *ListWarehouseResponseBodyData) SetWarehouseName(v string) *ListWarehouseResponseBodyData {
	s.WarehouseName = &v
	return s
}

func (s *ListWarehouseResponseBodyData) SetWarehouseType(v string) *ListWarehouseResponseBodyData {
	s.WarehouseType = &v
	return s
}

func (s *ListWarehouseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
