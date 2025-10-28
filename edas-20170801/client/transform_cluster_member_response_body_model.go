// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformClusterMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *TransformClusterMemberResponseBody
	GetCode() *int32
	SetData(v string) *TransformClusterMemberResponseBody
	GetData() *string
	SetMessage(v string) *TransformClusterMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *TransformClusterMemberResponseBody
	GetRequestId() *string
}

type TransformClusterMemberResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// `Transform submit success!` is returned if the request is successful.
	//
	// example:
	//
	// Transform submit success!
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// b197-40ab-9155-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransformClusterMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransformClusterMemberResponseBody) GoString() string {
	return s.String()
}

func (s *TransformClusterMemberResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *TransformClusterMemberResponseBody) GetData() *string {
	return s.Data
}

func (s *TransformClusterMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TransformClusterMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransformClusterMemberResponseBody) SetCode(v int32) *TransformClusterMemberResponseBody {
	s.Code = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetData(v string) *TransformClusterMemberResponseBody {
	s.Data = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetMessage(v string) *TransformClusterMemberResponseBody {
	s.Message = &v
	return s
}

func (s *TransformClusterMemberResponseBody) SetRequestId(v string) *TransformClusterMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransformClusterMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
