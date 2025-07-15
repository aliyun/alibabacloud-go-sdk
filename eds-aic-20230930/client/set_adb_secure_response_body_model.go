// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAdbSecureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SetAdbSecureResponseBodyData) *SetAdbSecureResponseBody
	GetData() *SetAdbSecureResponseBodyData
	SetRequestId(v string) *SetAdbSecureResponseBody
	GetRequestId() *string
}

type SetAdbSecureResponseBody struct {
	// The returned object.
	Data *SetAdbSecureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAdbSecureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAdbSecureResponseBody) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponseBody) GetData() *SetAdbSecureResponseBodyData {
	return s.Data
}

func (s *SetAdbSecureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAdbSecureResponseBody) SetData(v *SetAdbSecureResponseBodyData) *SetAdbSecureResponseBody {
	s.Data = v
	return s
}

func (s *SetAdbSecureResponseBody) SetRequestId(v string) *SetAdbSecureResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAdbSecureResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetAdbSecureResponseBodyData struct {
	// The number of the cloud phone instances for which the ADB authentication feature failed to be enabled.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The IDs of the cloud phone instances for which the ADB authentication feature is enabled.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The total number of the cloud phone instances.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s SetAdbSecureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetAdbSecureResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetAdbSecureResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *SetAdbSecureResponseBodyData) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *SetAdbSecureResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *SetAdbSecureResponseBodyData) SetFailCount(v int32) *SetAdbSecureResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *SetAdbSecureResponseBodyData) SetInstanceIds(v []*string) *SetAdbSecureResponseBodyData {
	s.InstanceIds = v
	return s
}

func (s *SetAdbSecureResponseBodyData) SetTotalCount(v int32) *SetAdbSecureResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *SetAdbSecureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
