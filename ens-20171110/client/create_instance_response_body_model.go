// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateInstanceResponseBody
	GetCode() *int32
	SetInstanceIds(v *CreateInstanceResponseBodyInstanceIds) *CreateInstanceResponseBody
	GetInstanceIds() *CreateInstanceResponseBodyInstanceIds
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
}

type CreateInstanceResponseBody struct {
	// The return code. A value of 0 indicates that the request is successful.
	//
	// >  If you call this operation by using SDKs, the return value is of the integer type. If you call this operation by using common methods or HTTP requests, the return value is of the string type.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The IDs of instances.
	InstanceIds *CreateInstanceResponseBodyInstanceIds `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 4A431388-2D4B-46F4-A96B-D4E6BD0688C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateInstanceResponseBody) GetInstanceIds() *CreateInstanceResponseBodyInstanceIds {
	return s.InstanceIds
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) SetCode(v int32) *CreateInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateInstanceResponseBody) SetInstanceIds(v *CreateInstanceResponseBodyInstanceIds) *CreateInstanceResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	if s.InstanceIds != nil {
		if err := s.InstanceIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateInstanceResponseBodyInstanceIds struct {
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
}

func (s CreateInstanceResponseBodyInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyInstanceIds) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyInstanceIds) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyInstanceIds) SetInstanceId(v []*string) *CreateInstanceResponseBodyInstanceIds {
	s.InstanceId = v
	return s
}

func (s *CreateInstanceResponseBodyInstanceIds) Validate() error {
	return dara.Validate(s)
}
