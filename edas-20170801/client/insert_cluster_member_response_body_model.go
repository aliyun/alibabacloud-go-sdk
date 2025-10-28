// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertClusterMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertClusterMemberResponseBody
	GetCode() *int32
	SetData(v string) *InsertClusterMemberResponseBody
	GetData() *string
	SetMessage(v string) *InsertClusterMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertClusterMemberResponseBody
	GetRequestId() *string
}

type InsertClusterMemberResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned for the request.
	//
	// example:
	//
	// Transform submit success!
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 01B49A88-B06F-423B-A5EF-E5C0A89******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertClusterMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *InsertClusterMemberResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertClusterMemberResponseBody) GetData() *string {
	return s.Data
}

func (s *InsertClusterMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertClusterMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertClusterMemberResponseBody) SetCode(v int32) *InsertClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetData(v string) *InsertClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetMessage(v string) *InsertClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *InsertClusterMemberResponseBody) SetRequestId(v string) *InsertClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertClusterMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
