// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDataSourceLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AddDataSourceLogResponseBodyData) *AddDataSourceLogResponseBody
	GetData() *AddDataSourceLogResponseBodyData
	SetRequestId(v string) *AddDataSourceLogResponseBody
	GetRequestId() *string
}

type AddDataSourceLogResponseBody struct {
	// The data returned.
	Data *AddDataSourceLogResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddDataSourceLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceLogResponseBody) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponseBody) GetData() *AddDataSourceLogResponseBodyData {
	return s.Data
}

func (s *AddDataSourceLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddDataSourceLogResponseBody) SetData(v *AddDataSourceLogResponseBodyData) *AddDataSourceLogResponseBody {
	s.Data = v
	return s
}

func (s *AddDataSourceLogResponseBody) SetRequestId(v string) *AddDataSourceLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddDataSourceLogResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDataSourceLogResponseBodyData struct {
	// The number of logs that are added. The value 1 indicates that the log is added, and a value less than or equal to 0 indicates that the log failed to be added.
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

func (s AddDataSourceLogResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddDataSourceLogResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddDataSourceLogResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *AddDataSourceLogResponseBodyData) GetLogInstanceId() *string {
	return s.LogInstanceId
}

func (s *AddDataSourceLogResponseBodyData) SetCount(v int32) *AddDataSourceLogResponseBodyData {
	s.Count = &v
	return s
}

func (s *AddDataSourceLogResponseBodyData) SetLogInstanceId(v string) *AddDataSourceLogResponseBodyData {
	s.LogInstanceId = &v
	return s
}

func (s *AddDataSourceLogResponseBodyData) Validate() error {
	return dara.Validate(s)
}
