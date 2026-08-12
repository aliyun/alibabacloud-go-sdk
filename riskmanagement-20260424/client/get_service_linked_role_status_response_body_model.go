// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceLinkedRoleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetServiceLinkedRoleStatusResponseBody
	GetCode() *string
	SetData(v *GetServiceLinkedRoleStatusResponseBodyData) *GetServiceLinkedRoleStatusResponseBody
	GetData() *GetServiceLinkedRoleStatusResponseBodyData
	SetMessage(v string) *GetServiceLinkedRoleStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetServiceLinkedRoleStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetServiceLinkedRoleStatusResponseBody
	GetSuccess() *bool
}

type GetServiceLinkedRoleStatusResponseBody struct {
	// The status code. Valid values:
	//
	// - **200**: Succeeded.
	//
	// - **Other (400, 500)**: Failed.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *GetServiceLinkedRoleStatusResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FBDD713-00A5-5C98-B661-3FD31A349B6E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetServiceLinkedRoleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetServiceLinkedRoleStatusResponseBody) GetData() *GetServiceLinkedRoleStatusResponseBodyData {
	return s.Data
}

func (s *GetServiceLinkedRoleStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetServiceLinkedRoleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceLinkedRoleStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetServiceLinkedRoleStatusResponseBody) SetCode(v string) *GetServiceLinkedRoleStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBody) SetData(v *GetServiceLinkedRoleStatusResponseBodyData) *GetServiceLinkedRoleStatusResponseBody {
	s.Data = v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBody) SetMessage(v string) *GetServiceLinkedRoleStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBody) SetRequestId(v string) *GetServiceLinkedRoleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBody) SetSuccess(v bool) *GetServiceLinkedRoleStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceLinkedRoleStatusResponseBodyData struct {
	// The authorization status. Valid values:
	//
	// - **true**: authorized
	//
	// - **false**: not authorized
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetServiceLinkedRoleStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetServiceLinkedRoleStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetServiceLinkedRoleStatusResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *GetServiceLinkedRoleStatusResponseBodyData) SetStatus(v bool) *GetServiceLinkedRoleStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetServiceLinkedRoleStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
