// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLakeCacheSizeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyLakeCacheSizeResponseBody
	GetCode() *int32
	SetData(v *ModifyLakeCacheSizeResponseBodyData) *ModifyLakeCacheSizeResponseBody
	GetData() *ModifyLakeCacheSizeResponseBodyData
	SetRequestId(v string) *ModifyLakeCacheSizeResponseBody
	GetRequestId() *string
}

type ModifyLakeCacheSizeResponseBody struct {
	// The status code. The value 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifyLakeCacheSizeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 494486CE-6F49-574E-B304-29127EA12E36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyLakeCacheSizeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakeCacheSizeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyLakeCacheSizeResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyLakeCacheSizeResponseBody) GetData() *ModifyLakeCacheSizeResponseBodyData {
	return s.Data
}

func (s *ModifyLakeCacheSizeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyLakeCacheSizeResponseBody) SetCode(v int32) *ModifyLakeCacheSizeResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyLakeCacheSizeResponseBody) SetData(v *ModifyLakeCacheSizeResponseBodyData) *ModifyLakeCacheSizeResponseBody {
	s.Data = v
	return s
}

func (s *ModifyLakeCacheSizeResponseBody) SetRequestId(v string) *ModifyLakeCacheSizeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyLakeCacheSizeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyLakeCacheSizeResponseBodyData struct {
	// The size of the lake cache space. Unit: GB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The size of the data that occupies the lake cache space. Unit: GB.
	//
	// example:
	//
	// 100
	DataSize *int64 `json:"DataSize,omitempty" xml:"DataSize,omitempty"`
	// The clusters that share the lake cache space.
	Instances []*string `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
}

func (s ModifyLakeCacheSizeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyLakeCacheSizeResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyLakeCacheSizeResponseBodyData) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ModifyLakeCacheSizeResponseBodyData) GetDataSize() *int64 {
	return s.DataSize
}

func (s *ModifyLakeCacheSizeResponseBodyData) GetInstances() []*string {
	return s.Instances
}

func (s *ModifyLakeCacheSizeResponseBodyData) SetCapacity(v int64) *ModifyLakeCacheSizeResponseBodyData {
	s.Capacity = &v
	return s
}

func (s *ModifyLakeCacheSizeResponseBodyData) SetDataSize(v int64) *ModifyLakeCacheSizeResponseBodyData {
	s.DataSize = &v
	return s
}

func (s *ModifyLakeCacheSizeResponseBodyData) SetInstances(v []*string) *ModifyLakeCacheSizeResponseBodyData {
	s.Instances = v
	return s
}

func (s *ModifyLakeCacheSizeResponseBodyData) Validate() error {
	return dara.Validate(s)
}
