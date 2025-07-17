// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagToFlinkClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddTagToFlinkClusterResponseBody
	GetCode() *int32
	SetData(v string) *AddTagToFlinkClusterResponseBody
	GetData() *string
	SetRequestId(v string) *AddTagToFlinkClusterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTagToFlinkClusterResponseBody
	GetSuccess() *bool
}

type AddTagToFlinkClusterResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the business logic was executed. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2FC13182-B9AF-4E6B-BE51-72669B7C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTagToFlinkClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTagToFlinkClusterResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagToFlinkClusterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddTagToFlinkClusterResponseBody) GetData() *string {
	return s.Data
}

func (s *AddTagToFlinkClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTagToFlinkClusterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTagToFlinkClusterResponseBody) SetCode(v int32) *AddTagToFlinkClusterResponseBody {
	s.Code = &v
	return s
}

func (s *AddTagToFlinkClusterResponseBody) SetData(v string) *AddTagToFlinkClusterResponseBody {
	s.Data = &v
	return s
}

func (s *AddTagToFlinkClusterResponseBody) SetRequestId(v string) *AddTagToFlinkClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagToFlinkClusterResponseBody) SetSuccess(v bool) *AddTagToFlinkClusterResponseBody {
	s.Success = &v
	return s
}

func (s *AddTagToFlinkClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
