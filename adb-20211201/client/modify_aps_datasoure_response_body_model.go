// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsDatasoureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyApsDatasoureResponseBody
	GetCode() *string
	SetData(v string) *ModifyApsDatasoureResponseBody
	GetData() *string
	SetHttpStatusCode(v string) *ModifyApsDatasoureResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ModifyApsDatasoureResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyApsDatasoureResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyApsDatasoureResponseBody
	GetSuccess() *string
}

type ModifyApsDatasoureResponseBody struct {
	// The API status code or POP error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of resource group information.
	//
	// example:
	//
	// 150
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned. A status code of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The execution result of the application suggestion. Valid values:
	//
	// - If the request is normal, **SUCCESS*	- is returned.
	//
	// - If the request is abnormal, the specific error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A3EB3BCD-D974-52D4-B75C-BB06505916CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApsDatasoureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsDatasoureResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApsDatasoureResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyApsDatasoureResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyApsDatasoureResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyApsDatasoureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyApsDatasoureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApsDatasoureResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyApsDatasoureResponseBody) SetCode(v string) *ModifyApsDatasoureResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) SetData(v string) *ModifyApsDatasoureResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) SetHttpStatusCode(v string) *ModifyApsDatasoureResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) SetMessage(v string) *ModifyApsDatasoureResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) SetRequestId(v string) *ModifyApsDatasoureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) SetSuccess(v string) *ModifyApsDatasoureResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApsDatasoureResponseBody) Validate() error {
	return dara.Validate(s)
}
