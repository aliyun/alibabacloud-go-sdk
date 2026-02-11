// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataSourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyDataSourceLogResponseBodyData) *ModifyDataSourceLogResponseBody
	GetData() *ModifyDataSourceLogResponseBodyData
	SetRequestId(v string) *ModifyDataSourceLogResponseBody
	GetRequestId() *string
}

type ModifyDataSourceLogResponseBody struct {
	// The data returned.
	Data *ModifyDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDataSourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponseBody) GetData() *ModifyDataSourceLogResponseBodyData {
	return s.Data
}

func (s *ModifyDataSourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDataSourceLogResponseBody) SetData(v *ModifyDataSourceLogResponseBodyData) *ModifyDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *ModifyDataSourceLogResponseBody) SetRequestId(v string) *ModifyDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDataSourceLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDataSourceLogResponseBodyData struct {
	// The number of logs that are modified. The value 1 indicates that the modification is successful, and a value less than or equal to 0 indicates that the modification failed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the log. The ID is an MD5 hash value that is calculated by the threat analysis feature based on specific parameters.
	//
	// example:
	//
	// 220ba97c9d1fdb0b9c7e8c7ca328d7ea
	LogInstanceId *string `json:"LogInstanceId,omitempty" xml:"LogInstanceId,omitempty"`
}

func (s ModifyDataSourceLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyDataSourceLogResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ModifyDataSourceLogResponseBodyData) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *ModifyDataSourceLogResponseBodyData) SetCount(v int32) *ModifyDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifyDataSourceLogResponseBodyData) SetLogInstanceId(v string) *ModifyDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

func (s *ModifyDataSourceLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
