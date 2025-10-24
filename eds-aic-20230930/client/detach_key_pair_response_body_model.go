// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetachKeyPairResponseBodyData) *DetachKeyPairResponseBody
	GetData() *DetachKeyPairResponseBodyData
	SetRequestId(v string) *DetachKeyPairResponseBody
	GetRequestId() *string
}

type DetachKeyPairResponseBody struct {
	// The object that is returned.
	Data *DetachKeyPairResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBody) GetData() *DetachKeyPairResponseBodyData {
	return s.Data
}

func (s *DetachKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachKeyPairResponseBody) SetData(v *DetachKeyPairResponseBodyData) *DetachKeyPairResponseBody {
	s.Data = v
	return s
}

func (s *DetachKeyPairResponseBody) SetRequestId(v string) *DetachKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachKeyPairResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetachKeyPairResponseBodyData struct {
	// The IDs of the cloud phone instances from which the ADB key pair is successfully detached.
	DetachedInstanceIds []*string `json:"DetachedInstanceIds,omitempty" xml:"DetachedInstanceIds,omitempty" type:"Repeated"`
	// The number of the cloud phone instances from which the ADB key pair failed to be detached.
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
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DetachKeyPairResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetachKeyPairResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetachKeyPairResponseBodyData) GetDetachedInstanceIds() []*string {
	return s.DetachedInstanceIds
}

func (s *DetachKeyPairResponseBodyData) GetFailCount() *int32 {
	return s.FailCount
}

func (s *DetachKeyPairResponseBodyData) GetKeyPairId() *string {
	return s.KeyPairId
}

func (s *DetachKeyPairResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DetachKeyPairResponseBodyData) SetDetachedInstanceIds(v []*string) *DetachKeyPairResponseBodyData {
	s.DetachedInstanceIds = v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetFailCount(v int32) *DetachKeyPairResponseBodyData {
	s.FailCount = &v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetKeyPairId(v string) *DetachKeyPairResponseBodyData {
	s.KeyPairId = &v
	return s
}

func (s *DetachKeyPairResponseBodyData) SetTotalCount(v int32) *DetachKeyPairResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DetachKeyPairResponseBodyData) Validate() error {
	return dara.Validate(s)
}
