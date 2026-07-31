// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterCreateModelGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterCreateModelGroupResponseBodyData) *ModelRouterCreateModelGroupResponseBody
	GetData() *ModelRouterCreateModelGroupResponseBodyData
	SetErrCode(v string) *ModelRouterCreateModelGroupResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterCreateModelGroupResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterCreateModelGroupResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterCreateModelGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterCreateModelGroupResponseBody
	GetSuccess() *bool
}

type ModelRouterCreateModelGroupResponseBody struct {
	// The response struct.
	Data *ModelRouterCreateModelGroupResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message code.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Unknown error
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterCreateModelGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelGroupResponseBody) GetData() *ModelRouterCreateModelGroupResponseBodyData {
	return s.Data
}

func (s *ModelRouterCreateModelGroupResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterCreateModelGroupResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterCreateModelGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterCreateModelGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterCreateModelGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterCreateModelGroupResponseBody) SetData(v *ModelRouterCreateModelGroupResponseBodyData) *ModelRouterCreateModelGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) SetErrCode(v string) *ModelRouterCreateModelGroupResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) SetErrMessage(v string) *ModelRouterCreateModelGroupResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) SetHttpStatusCode(v int32) *ModelRouterCreateModelGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) SetRequestId(v string) *ModelRouterCreateModelGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) SetSuccess(v bool) *ModelRouterCreateModelGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterCreateModelGroupResponseBodyData struct {
	// The unique identifier of the group. The identifier has the mg_ prefix and does not change after creation.
	//
	// example:
	//
	// mg_a1b2c3d4e5f6g7h8i9j0
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
}

func (s ModelRouterCreateModelGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterCreateModelGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterCreateModelGroupResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *ModelRouterCreateModelGroupResponseBodyData) SetGroupId(v string) *ModelRouterCreateModelGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ModelRouterCreateModelGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
