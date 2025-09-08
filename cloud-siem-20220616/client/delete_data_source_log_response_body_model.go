// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataSourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteDataSourceLogResponseBodyData) *DeleteDataSourceLogResponseBody
	GetData() *DeleteDataSourceLogResponseBodyData
	SetRequestId(v string) *DeleteDataSourceLogResponseBody
	GetRequestId() *string
}

type DeleteDataSourceLogResponseBody struct {
	// The data returned.
	Data *DeleteDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponseBody) GetData() *DeleteDataSourceLogResponseBodyData {
	return s.Data
}

func (s *DeleteDataSourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataSourceLogResponseBody) SetData(v *DeleteDataSourceLogResponseBodyData) *DeleteDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *DeleteDataSourceLogResponseBody) SetRequestId(v string) *DeleteDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataSourceLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DeleteDataSourceLogResponseBodyData struct {
	// The number of logs that are removed. The value 1 indicates that the log is removed, and a value less than or equal to 0 indicates that the log failed to be removed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// ef33097c9d1fdb0b9c7e8c7ca320pkl1
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
}

func (s DeleteDataSourceLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceLogResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *DeleteDataSourceLogResponseBodyData) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *DeleteDataSourceLogResponseBodyData) SetCount(v int32) *DeleteDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *DeleteDataSourceLogResponseBodyData) SetLogInstanceId(v string) *DeleteDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

func (s *DeleteDataSourceLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
