// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDataSourceResponseBodyData) *ModifyDataSourceResponseBody
	GetData() *ModifyDataSourceResponseBodyData
	SetRequestId(v string) *ModifyDataSourceResponseBody
	GetRequestId() *string
}

type ModifyDataSourceResponseBody struct {
	// The data returned.
	Data *ModifyDataSourceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBody) GetData() *ModifyDataSourceResponseBodyData {
	return s.Data
}

func (s *ModifyDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataSourceResponseBody) SetData(v *ModifyDataSourceResponseBodyData) *ModifyDataSourceResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDataSourceResponseBody) SetRequestId(v string) *ModifyDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDataSourceResponseBodyData struct {
	// The number of modified data sources. A value of 1 indicates success. A value of 0 or less indicates failure.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the data source. The ID is an MD5 hash that is calculated based on the values of other parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	DataSourceInstanceId *string `json:"DataSourceInstanceId,omitempty" xml:"DataSourceInstanceId,omitempty"`
}

func (s ModifyDataSourceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ModifyDataSourceResponseBodyData) GetDataSourceInstanceId() *string {
	return s.DataSourceInstanceId
}

func (s *ModifyDataSourceResponseBodyData) SetCount(v int32) *ModifyDataSourceResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifyDataSourceResponseBodyData) SetDataSourceInstanceId(v string) *ModifyDataSourceResponseBodyData {
	s.DataSourceInstanceId = &v
	return s
}

func (s *ModifyDataSourceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
