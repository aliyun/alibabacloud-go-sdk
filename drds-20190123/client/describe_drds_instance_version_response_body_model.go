// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstanceVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDrdsInstanceVersionResponseBodyData) *DescribeDrdsInstanceVersionResponseBody
	GetData() *DescribeDrdsInstanceVersionResponseBodyData
	SetRequestId(v string) *DescribeDrdsInstanceVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDrdsInstanceVersionResponseBody
	GetSuccess() *bool
}

type DescribeDrdsInstanceVersionResponseBody struct {
	// The details about the instance version.
	Data *DescribeDrdsInstanceVersionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2F7F8080-9132-4279-85D0-B7E5C4305162
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBody) GetData() *DescribeDrdsInstanceVersionResponseBodyData {
	return s.Data
}

func (s *DescribeDrdsInstanceVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDrdsInstanceVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetData(v *DescribeDrdsInstanceVersionResponseBodyData) *DescribeDrdsInstanceVersionResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetRequestId(v string) *DescribeDrdsInstanceVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) SetSuccess(v bool) *DescribeDrdsInstanceVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDrdsInstanceVersionResponseBodyData struct {
	// The current version of the instance.
	//
	// example:
	//
	// 5.3.12-15682777
	InstanceVersion *string `json:"InstanceVersion,omitempty" xml:"InstanceVersion,omitempty"`
	// The latest version of the instance.
	//
	// example:
	//
	// 5.4.12-16315258
	NewestVersion *string `json:"NewestVersion,omitempty" xml:"NewestVersion,omitempty"`
}

func (s DescribeDrdsInstanceVersionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstanceVersionResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) GetInstanceVersion() *string {
	return s.InstanceVersion
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) GetNewestVersion() *string {
	return s.NewestVersion
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetInstanceVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.InstanceVersion = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) SetNewestVersion(v string) *DescribeDrdsInstanceVersionResponseBodyData {
	s.NewestVersion = &v
	return s
}

func (s *DescribeDrdsInstanceVersionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
