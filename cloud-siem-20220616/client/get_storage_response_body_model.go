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
	// The details of the storage settings.
	Data *GetStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStorageResponseBodyData struct {
	// Indicates whether you can change the storage region. You can change the storage region only once. The default value is false. Valid values:
	//
	// - true: You can change the storage region.
	//
	// - false: You cannot change the storage region.
	//
	// example:
	//
	// false
	CanOperate *bool `json:"CanOperate,omitempty" xml:"CanOperate,omitempty"`
	// Indicates whether you have permission to change the storage region. The default value is false. Valid values:
	//
	// - true: You have permission.
	//
	// - false: You do not have permission.
	//
	// example:
	//
	// false
	DisplayRegion *bool `json:"DisplayRegion,omitempty" xml:"DisplayRegion,omitempty"`
	// The storage region.
	//
	// If the Data Management hub is in the **cn-hangzhou*	- region, the default value of **Region*	- is \\`cn-shanghai\\`. If the Data Management hub is in the **ap-southeast-1*	- region, the default value of **Region*	- is \\`ap-southeast-1\\`.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The storage duration in days. The default value is 180. The value must be an integer from 30 to 3000.
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
