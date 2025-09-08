// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddDataSourceResponseBodyData) *AddDataSourceResponseBody
	GetData() *AddDataSourceResponseBodyData
	SetRequestId(v string) *AddDataSourceResponseBody
	GetRequestId() *string
}

type AddDataSourceResponseBody struct {
	// The data returned.
	Data *AddDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBody) GetData() *AddDataSourceResponseBodyData {
	return s.Data
}

func (s *AddDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataSourceResponseBody) SetData(v *AddDataSourceResponseBodyData) *AddDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *AddDataSourceResponseBody) SetRequestId(v string) *AddDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddDataSourceResponseBodyData struct {
	// The number of data sources that are added. The value 1 indicates that data source is added, and a value less than or equal to 0 indicates that the data source failed to be added.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the data source. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
}

func (s AddDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataSourceResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *AddDataSourceResponseBodyData) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *AddDataSourceResponseBodyData) SetCount(v int32) *AddDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *AddDataSourceResponseBodyData) SetDataSourceInstanceId(v string) *AddDataSourceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *AddDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
