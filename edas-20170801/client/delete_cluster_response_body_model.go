// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteClusterResponseBody
	GetCode() *int32
	SetData(v bool) *DeleteClusterResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
}

type DeleteClusterResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the cluster is deleted. Valid values:
	//
	// 	- true: The cluster is deleted.
	//
	// 	- false: The cluster is not deleted.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// 856D4438-****-4EA9-****-894628C0434E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteClusterResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) SetCode(v int32) *DeleteClusterResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteClusterResponseBody) SetData(v bool) *DeleteClusterResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteClusterResponseBody) SetMessage(v string) *DeleteClusterResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
