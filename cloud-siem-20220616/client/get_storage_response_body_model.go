// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetStorageResponseBodyData) *GetStorageResponseBody
	GetData() *GetStorageResponseBodyData
	SetRequestId(v string) *GetStorageResponseBody
	GetRequestId() *string
}

type GetStorageResponseBody struct {
	// The information about the storage.
	Data *GetStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 97A31C3A-3F9F-5866-8979-5159E3DC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBody) GetData() *GetStorageResponseBodyData {
	return s.Data
}

func (s *GetStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageResponseBody) SetData(v *GetStorageResponseBodyData) *GetStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetStorageResponseBody) SetRequestId(v string) *GetStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageResponseBodyData struct {
	// Indicates whether the storage region can be changed for once. Default value: false Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	CanOperate *bool `json:"CanOperate,omitempty" xml:"CanOperate,omitempty"`
	// Indicates whether the storage region can be changed. Default value: false Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	DisplayRegion *bool `json:"DisplayRegion,omitempty" xml:"DisplayRegion,omitempty"`
	// The region where the data is stored.
	//
	// If the data management center is **cn-hangzhou**, the default value of **Region*	- is cn-shanghai, which specifies the China (Shanghai) region. If the data management center is **ap-southeast-1**, the default value of **Region*	- is ap-southeast-1, which specifies the Singapore region.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The storage period of logs. Unit: day. Default value: 180. Valid values: 30 to 3000.
	//
	// example:
	//
	// 180
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s GetStorageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetStorageResponseBodyData) GetCanOperate() *bool {
	return s.CanOperate
}

func (s *GetStorageResponseBodyData) GetDisplayRegion() *bool {
	return s.DisplayRegion
}

func (s *GetStorageResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetStorageResponseBodyData) GetTtl() *int32 {
	return s.Ttl
}

func (s *GetStorageResponseBodyData) SetCanOperate(v bool) *GetStorageResponseBodyData {
	s.CanOperate = &v
	return s
}

func (s *GetStorageResponseBodyData) SetDisplayRegion(v bool) *GetStorageResponseBodyData {
	s.DisplayRegion = &v
	return s
}

func (s *GetStorageResponseBodyData) SetRegion(v string) *GetStorageResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetStorageResponseBodyData) SetTtl(v int32) *GetStorageResponseBodyData {
	s.Ttl = &v
	return s
}

func (s *GetStorageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
