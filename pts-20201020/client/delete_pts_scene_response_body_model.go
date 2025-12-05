// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeletePtsSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeletePtsSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeletePtsSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeletePtsSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeletePtsSceneResponseBody
	GetSuccess() *bool
}

type DeletePtsSceneResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 4001
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A3ED870E-C3BF-44F4-B460-A30785E0256B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeletePtsSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeletePtsSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeletePtsSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeletePtsSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePtsSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeletePtsSceneResponseBody) SetCode(v string) *DeletePtsSceneResponseBody {
	s.Code = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetHttpStatusCode(v int32) *DeletePtsSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetMessage(v string) *DeletePtsSceneResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetRequestId(v string) *DeletePtsSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePtsSceneResponseBody) SetSuccess(v bool) *DeletePtsSceneResponseBody {
	s.Success = &v
	return s
}

func (s *DeletePtsSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
