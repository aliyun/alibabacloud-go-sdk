// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveUsersResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveUsersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveUsersResponseBody
	GetMessage() *string
	SetParams(v []*string) *RemoveUsersResponseBody
	GetParams() []*string
	SetRequestId(v string) *RemoveUsersResponseBody
	GetRequestId() *string
}

type RemoveUsersResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// EEEE671A-3E24-4A04-81E6-6C4F5B39DF75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUsersResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveUsersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveUsersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveUsersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *RemoveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUsersResponseBody) SetCode(v string) *RemoveUsersResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUsersResponseBody) SetHttpStatusCode(v int32) *RemoveUsersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveUsersResponseBody) SetMessage(v string) *RemoveUsersResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUsersResponseBody) SetParams(v []*string) *RemoveUsersResponseBody {
	s.Params = v
	return s
}

func (s *RemoveUsersResponseBody) SetRequestId(v string) *RemoveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
