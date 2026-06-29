// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrganizationMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddOrganizationMemberResponseBody
	GetCode() *string
	SetData(v *AddOrganizationMemberResponseBodyData) *AddOrganizationMemberResponseBody
	GetData() *AddOrganizationMemberResponseBodyData
	SetHttpStatusCode(v int32) *AddOrganizationMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddOrganizationMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrganizationMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddOrganizationMemberResponseBody
	GetSuccess() *bool
}

type AddOrganizationMemberResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data.
	Data *AddOrganizationMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// None
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 385C2BC3-52FC-564F-9312-97E5DFE1DFC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddOrganizationMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrganizationMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrganizationMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddOrganizationMemberResponseBody) GetData() *AddOrganizationMemberResponseBodyData {
	return s.Data
}

func (s *AddOrganizationMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddOrganizationMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrganizationMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrganizationMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddOrganizationMemberResponseBody) SetCode(v string) *AddOrganizationMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddOrganizationMemberResponseBody) SetData(v *AddOrganizationMemberResponseBodyData) *AddOrganizationMemberResponseBody {
	s.Data = v
	return s
}

func (s *AddOrganizationMemberResponseBody) SetHttpStatusCode(v int32) *AddOrganizationMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddOrganizationMemberResponseBody) SetMessage(v string) *AddOrganizationMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrganizationMemberResponseBody) SetRequestId(v string) *AddOrganizationMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrganizationMemberResponseBody) SetSuccess(v bool) *AddOrganizationMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddOrganizationMemberResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddOrganizationMemberResponseBodyData struct {
	// The account ID.
	//
	// example:
	//
	// acc_123456789
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// Indicates whether the seat was assigned successfully.
	//
	// example:
	//
	// true
	SeatAssigned *bool `json:"SeatAssigned,omitempty" xml:"SeatAssigned,omitempty"`
}

func (s AddOrganizationMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddOrganizationMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddOrganizationMemberResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *AddOrganizationMemberResponseBodyData) GetSeatAssigned() *bool {
	return s.SeatAssigned
}

func (s *AddOrganizationMemberResponseBodyData) SetAccountId(v string) *AddOrganizationMemberResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *AddOrganizationMemberResponseBodyData) SetSeatAssigned(v bool) *AddOrganizationMemberResponseBodyData {
	s.SeatAssigned = &v
	return s
}

func (s *AddOrganizationMemberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
