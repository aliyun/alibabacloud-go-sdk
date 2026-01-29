// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyAccountRelationResponseBody
	GetCode() *string
	SetData(v *ModifyAccountRelationResponseBodyData) *ModifyAccountRelationResponseBody
	GetData() *ModifyAccountRelationResponseBodyData
	SetMessage(v string) *ModifyAccountRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyAccountRelationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyAccountRelationResponseBody
	GetSuccess() *bool
}

type ModifyAccountRelationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *ModifyAccountRelationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Message returned
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// request_id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyAccountRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountRelationResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAccountRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyAccountRelationResponseBody) GetData() *ModifyAccountRelationResponseBodyData {
	return s.Data
}

func (s *ModifyAccountRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyAccountRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAccountRelationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyAccountRelationResponseBody) SetCode(v string) *ModifyAccountRelationResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyAccountRelationResponseBody) SetData(v *ModifyAccountRelationResponseBodyData) *ModifyAccountRelationResponseBody {
	s.Data = v
	return s
}

func (s *ModifyAccountRelationResponseBody) SetMessage(v string) *ModifyAccountRelationResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyAccountRelationResponseBody) SetRequestId(v string) *ModifyAccountRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAccountRelationResponseBody) SetSuccess(v bool) *ModifyAccountRelationResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyAccountRelationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAccountRelationResponseBodyData struct {
	// HostId
	//
	// example:
	//
	// HostId
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s ModifyAccountRelationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountRelationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyAccountRelationResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *ModifyAccountRelationResponseBodyData) SetHostId(v string) *ModifyAccountRelationResponseBodyData {
	s.HostId = &v
	return s
}

func (s *ModifyAccountRelationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
