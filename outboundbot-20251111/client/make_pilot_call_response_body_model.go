// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakePilotCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MakePilotCallResponseBody
	GetCode() *string
	SetData(v *MakePilotCallResponseBodyData) *MakePilotCallResponseBody
	GetData() *MakePilotCallResponseBodyData
	SetHttpStatusCode(v int32) *MakePilotCallResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *MakePilotCallResponseBody
	GetMessage() *string
	SetParams(v []*string) *MakePilotCallResponseBody
	GetParams() []*string
	SetRequestId(v string) *MakePilotCallResponseBody
	GetRequestId() *string
}

type MakePilotCallResponseBody struct {
	// The business processing status code. OK indicates success. When an exception occurs, the corresponding business error code is returned. Error types and mapping rules follow the current API configuration. No error codes are added or modified in this operation.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The session ID of the current real-number outbound call test. The return of this object does not indicate that the called party has answered.
	//
	// example:
	//
	// {"SessionId":"test-11111111-2222-4333-8444-555555555555"}
	Data *MakePilotCallResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code recorded in the business response body. A normal return value is 200. To determine whether the business operation is successful, you must also check Code. You cannot determine the success of the call solely based on the presence of this field.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The business response or error description. This field is usually empty or not returned upon success. A fixed text cannot be required upon success. Upon failure, this field contains an error description to assist in troubleshooting. Combine this field with Code and Params to understand the cause of the error. The example shows an error message when the instance does not exist, where Code is NotExists.InstanceId. This is for format reference only.
	//
	// example:
	//
	// Instance instance-example-001 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of parameters used for the localized display of error messages. This is not a scenario variable or business result. A normal successful response usually returns an empty array. When an error occurs, parameters corresponding to the error template may be returned.
	//
	// example:
	//
	// []
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The ID used to associate and troubleshoot the current POP request. The value in the response corresponds to the current request. System request parameters are injected by the gateway. Callers do not need to fill in system fields.
	//
	// example:
	//
	// AAAAAAAA-BBBB-4CCC-8DDD-EEEEEEEEEEEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MakePilotCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MakePilotCallResponseBody) GoString() string {
	return s.String()
}

func (s *MakePilotCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *MakePilotCallResponseBody) GetData() *MakePilotCallResponseBodyData {
	return s.Data
}

func (s *MakePilotCallResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *MakePilotCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MakePilotCallResponseBody) GetParams() []*string {
	return s.Params
}

func (s *MakePilotCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MakePilotCallResponseBody) SetCode(v string) *MakePilotCallResponseBody {
	s.Code = &v
	return s
}

func (s *MakePilotCallResponseBody) SetData(v *MakePilotCallResponseBodyData) *MakePilotCallResponseBody {
	s.Data = v
	return s
}

func (s *MakePilotCallResponseBody) SetHttpStatusCode(v int32) *MakePilotCallResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *MakePilotCallResponseBody) SetMessage(v string) *MakePilotCallResponseBody {
	s.Message = &v
	return s
}

func (s *MakePilotCallResponseBody) SetParams(v []*string) *MakePilotCallResponseBody {
	s.Params = v
	return s
}

func (s *MakePilotCallResponseBody) SetRequestId(v string) *MakePilotCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *MakePilotCallResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type MakePilotCallResponseBodyData struct {
	// The session ID used to uniquely associate the current interaction. For a text session, it is the text- ID returned when the session is started. For web voice, it is the ws- ID. For a real-number outbound call test, it is the test- ID. Subsequent calls must use the actual return value.
	//
	// example:
	//
	// test-11111111-2222-4333-8444-555555555555
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s MakePilotCallResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MakePilotCallResponseBodyData) GoString() string {
	return s.String()
}

func (s *MakePilotCallResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *MakePilotCallResponseBodyData) SetSessionId(v string) *MakePilotCallResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *MakePilotCallResponseBodyData) Validate() error {
	return dara.Validate(s)
}
