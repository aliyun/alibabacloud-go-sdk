// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAttackTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAttackTargetResponseBodyData) *CreateAttackTargetResponseBody
	GetData() *CreateAttackTargetResponseBodyData
	SetRequestId(v string) *CreateAttackTargetResponseBody
	GetRequestId() *string
}

type CreateAttackTargetResponseBody struct {
	// The operation result. Upon successful creation, the TargetId of the new scan target is returned.
	Data *CreateAttackTargetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, used for troubleshooting and log tracing.
	//
	// example:
	//
	// 1EBD0C05-6C1F-4C95-9C63-B7AB7B5A9C8E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAttackTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackTargetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAttackTargetResponseBody) GetData() *CreateAttackTargetResponseBodyData {
	return s.Data
}

func (s *CreateAttackTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAttackTargetResponseBody) SetData(v *CreateAttackTargetResponseBodyData) *CreateAttackTargetResponseBody {
	s.Data = v
	return s
}

func (s *CreateAttackTargetResponseBody) SetRequestId(v string) *CreateAttackTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAttackTargetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAttackTargetResponseBodyData struct {
	// The unique identifier of the created scan target. You can use this value as the TargetId parameter in subsequent calls such as TestConnectivity and scan task creation.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s CreateAttackTargetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAttackTargetResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAttackTargetResponseBodyData) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateAttackTargetResponseBodyData) SetTargetId(v string) *CreateAttackTargetResponseBodyData {
	s.TargetId = &v
	return s
}

func (s *CreateAttackTargetResponseBodyData) Validate() error {
	return dara.Validate(s)
}
