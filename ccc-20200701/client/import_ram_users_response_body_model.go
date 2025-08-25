// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRamUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportRamUsersResponseBody
	GetCode() *string
	SetData(v string) *ImportRamUsersResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ImportRamUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportRamUsersResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportRamUsersResponseBody
	GetRequestId() *string
}

type ImportRamUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1ca2b084-6f0a-454b-9851-29768a9a5832
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 03C67DAD-EB26-41D8-949D-9B0C470FB716
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImportRamUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportRamUsersResponseBody) GoString() string {
	return s.String()
}

func (s *ImportRamUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportRamUsersResponseBody) GetData() *string {
	return s.Data
}

func (s *ImportRamUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportRamUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportRamUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportRamUsersResponseBody) SetCode(v string) *ImportRamUsersResponseBody {
	s.Code = &v
	return s
}

func (s *ImportRamUsersResponseBody) SetData(v string) *ImportRamUsersResponseBody {
	s.Data = &v
	return s
}

func (s *ImportRamUsersResponseBody) SetHttpStatusCode(v int32) *ImportRamUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportRamUsersResponseBody) SetMessage(v string) *ImportRamUsersResponseBody {
	s.Message = &v
	return s
}

func (s *ImportRamUsersResponseBody) SetRequestId(v string) *ImportRamUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportRamUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
