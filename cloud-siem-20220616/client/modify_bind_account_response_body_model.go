// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBindAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyBindAccountResponseBodyData) *ModifyBindAccountResponseBody
	GetData() *ModifyBindAccountResponseBodyData
	SetRequestId(v string) *ModifyBindAccountResponseBody
	GetRequestId() *string
}

type ModifyBindAccountResponseBody struct {
	// The return value of the request.
	Data *ModifyBindAccountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBindAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBindAccountResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponseBody) GetData() *ModifyBindAccountResponseBodyData {
	return s.Data
}

func (s *ModifyBindAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBindAccountResponseBody) SetData(v *ModifyBindAccountResponseBodyData) *ModifyBindAccountResponseBody {
	s.Data = v
	return s
}

func (s *ModifyBindAccountResponseBody) SetRequestId(v string) *ModifyBindAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBindAccountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyBindAccountResponseBodyData struct {
	// The number of account bindings that are modified. A value of 1 indicates that the operation is successful. A value of 0 or less indicates that the operation failed.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s ModifyBindAccountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyBindAccountResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyBindAccountResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *ModifyBindAccountResponseBodyData) SetCount(v int32) *ModifyBindAccountResponseBodyData {
	s.Count = &v
	return s
}

func (s *ModifyBindAccountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
