// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAllMemberResponseBody
	GetCode() *string
	SetData(v []*ListAllMemberResponseBodyData) *ListAllMemberResponseBody
	GetData() []*ListAllMemberResponseBodyData
	SetHttpStatusCode(v int32) *ListAllMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListAllMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAllMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAllMemberResponseBody
	GetSuccess() *bool
}

type ListAllMemberResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*ListAllMemberResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAllMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAllMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListAllMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAllMemberResponseBody) GetData() []*ListAllMemberResponseBodyData {
	return s.Data
}

func (s *ListAllMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListAllMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAllMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAllMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAllMemberResponseBody) SetCode(v string) *ListAllMemberResponseBody {
	s.Code = &v
	return s
}

func (s *ListAllMemberResponseBody) SetData(v []*ListAllMemberResponseBodyData) *ListAllMemberResponseBody {
	s.Data = v
	return s
}

func (s *ListAllMemberResponseBody) SetHttpStatusCode(v int32) *ListAllMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListAllMemberResponseBody) SetMessage(v string) *ListAllMemberResponseBody {
	s.Message = &v
	return s
}

func (s *ListAllMemberResponseBody) SetRequestId(v string) *ListAllMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAllMemberResponseBody) SetSuccess(v bool) *ListAllMemberResponseBody {
	s.Success = &v
	return s
}

func (s *ListAllMemberResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAllMemberResponseBodyData struct {
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	Name     *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListAllMemberResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAllMemberResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAllMemberResponseBodyData) GetMemberId() *string {
	return s.MemberId
}

func (s *ListAllMemberResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAllMemberResponseBodyData) SetMemberId(v string) *ListAllMemberResponseBodyData {
	s.MemberId = &v
	return s
}

func (s *ListAllMemberResponseBodyData) SetName(v string) *ListAllMemberResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAllMemberResponseBodyData) Validate() error {
	return dara.Validate(s)
}
