// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsWorkloadNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyApsWorkloadNameResponseBody
	GetCode() *string
	SetData(v string) *ModifyApsWorkloadNameResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *ModifyApsWorkloadNameResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyApsWorkloadNameResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyApsWorkloadNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyApsWorkloadNameResponseBody
	GetSuccess() *bool
}

type ModifyApsWorkloadNameResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// InvalidInput
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 123
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The status code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// SUCCESS
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApsWorkloadNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsWorkloadNameResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApsWorkloadNameResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyApsWorkloadNameResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyApsWorkloadNameResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyApsWorkloadNameResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyApsWorkloadNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApsWorkloadNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyApsWorkloadNameResponseBody) SetCode(v string) *ModifyApsWorkloadNameResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) SetData(v string) *ModifyApsWorkloadNameResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) SetHttpStatusCode(v int32) *ModifyApsWorkloadNameResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) SetMessage(v string) *ModifyApsWorkloadNameResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) SetRequestId(v string) *ModifyApsWorkloadNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) SetSuccess(v bool) *ModifyApsWorkloadNameResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApsWorkloadNameResponseBody) Validate() error {
	return dara.Validate(s)
}
