// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAppGroupResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteAppGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAppGroupResponseBody
	GetSuccess() *bool
}

type DeleteAppGroupResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The additional information that is returned.
	//
	// example:
	//
	// app is not existed, groupId=xxxx, namesapce=xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indices whether the request was successful. Valid values:
	//
	// true: The request was successful.
	//
	// false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppGroupResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAppGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppGroupResponseBody) SetCode(v int32) *DeleteAppGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetMessage(v string) *DeleteAppGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetRequestId(v string) *DeleteAppGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppGroupResponseBody) SetSuccess(v bool) *DeleteAppGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAppGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
