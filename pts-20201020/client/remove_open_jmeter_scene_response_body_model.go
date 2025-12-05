// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOpenJMeterSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveOpenJMeterSceneResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveOpenJMeterSceneResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveOpenJMeterSceneResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveOpenJMeterSceneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveOpenJMeterSceneResponseBody
	GetSuccess() *bool
}

type RemoveOpenJMeterSceneResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
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

func (s RemoveOpenJMeterSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveOpenJMeterSceneResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveOpenJMeterSceneResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveOpenJMeterSceneResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveOpenJMeterSceneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveOpenJMeterSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveOpenJMeterSceneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveOpenJMeterSceneResponseBody) SetCode(v string) *RemoveOpenJMeterSceneResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetHttpStatusCode(v int32) *RemoveOpenJMeterSceneResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetMessage(v string) *RemoveOpenJMeterSceneResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetRequestId(v string) *RemoveOpenJMeterSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) SetSuccess(v bool) *RemoveOpenJMeterSceneResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveOpenJMeterSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
