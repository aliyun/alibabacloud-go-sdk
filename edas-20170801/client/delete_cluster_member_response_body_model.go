// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteClusterMemberResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteClusterMemberResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteClusterMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClusterMemberResponseBody
	GetRequestId() *string
}

type DeleteClusterMemberResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-7ca7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterMemberResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteClusterMemberResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteClusterMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClusterMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterMemberResponseBody) SetCode(v int32) *DeleteClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetData(v bool) *DeleteClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetMessage(v string) *DeleteClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) SetRequestId(v string) *DeleteClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
