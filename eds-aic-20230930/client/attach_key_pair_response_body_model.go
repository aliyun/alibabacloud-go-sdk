// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AttachKeyPairResponseBodyData) *AttachKeyPairResponseBody
	GetData() *AttachKeyPairResponseBodyData
	SetRequestId(v string) *AttachKeyPairResponseBody
	GetRequestId() *string
}

type AttachKeyPairResponseBody struct {
	// The object that is returned.
	Data *AttachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBody) GetData() *AttachKeyPairResponseBodyData {
	return s.Data
}

func (s *AttachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachKeyPairResponseBody) SetData(v *AttachKeyPairResponseBodyData) *AttachKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *AttachKeyPairResponseBody) SetRequestId(v string) *AttachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}

type AttachKeyPairResponseBodyData struct {
	// The IDs of the cloud phone instances to which the ADB key pair is successfully attached.
	AttachedInstanceIds []*string `json:"AttachedInstanceIds,omitempty" xml:"AttachedInstanceIds,omitempty" type:"Repeated"`
	// The number of the cloud phone instances to which the ADB key pair failed to be attached.
	//
	// example:
	//
	// 0
	FailCount *int32 `json:"FailCount,omitempty" xml:"FailCount,omitempty"`
	// The ID of the ADB key pair.
	//
	// example:
	//
	// kp-6v2q33ae4tw3a****
	KeyPairId *string `json:"KeyPairId,omitempty" xml:"KeyPairId,omitempty"`
	// The total number of the cloud phone instances.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s AttachKeyPairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AttachKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *AttachKeyPairResponseBodyData) GetAttachedInstanceIds() []*string {
	return s.AttachedInstanceIds
}

func (s *AttachKeyPairResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *AttachKeyPairResponseBodyData) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *AttachKeyPairResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AttachKeyPairResponseBodyData) SetAttachedInstanceIds(v []*string) *AttachKeyPairResponseBodyData {
	s.AttachedInstanceIds = v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetFailCount(v int32) *AttachKeyPairResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetKeyPairId(v string) *AttachKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *AttachKeyPairResponseBodyData) SetTotalCount(v int32) *AttachKeyPairResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *AttachKeyPairResponseBodyData) Validate() error {
	return dara.Validate(s)
}
