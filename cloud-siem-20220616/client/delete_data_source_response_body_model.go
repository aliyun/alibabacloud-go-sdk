// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataSourceResponseBodyData) *DeleteDataSourceResponseBody
	GetData() *DeleteDataSourceResponseBodyData
	SetRequestId(v string) *DeleteDataSourceResponseBody
	GetRequestId() *string
}

type DeleteDataSourceResponseBody struct {
	// The data returned.
	Data *DeleteDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) GetData() *DeleteDataSourceResponseBodyData {
	return s.Data
}

func (s *DeleteDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceResponseBody) SetData(v *DeleteDataSourceResponseBodyData) *DeleteDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteDataSourceResponseBodyData struct {
	// The number of data sources that are removed. The value 1 indicates that data source is removed, and a value less than or equal to 0 indicates that the data source failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DeleteDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DeleteDataSourceResponseBodyData) SetCount(v int32) *DeleteDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
