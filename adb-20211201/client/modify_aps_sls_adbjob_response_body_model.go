// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApsSlsADBJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyApsSlsADBJobResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *ModifyApsSlsADBJobResponseBody
	GetData() map[string]interface{}
	SetHttpStatusCode(v int32) *ModifyApsSlsADBJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModifyApsSlsADBJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyApsSlsADBJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyApsSlsADBJobResponseBody
	GetSuccess() *bool
}

type ModifyApsSlsADBJobResponseBody struct {
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
	// -
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
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
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyApsSlsADBJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyApsSlsADBJobResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyApsSlsADBJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyApsSlsADBJobResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *ModifyApsSlsADBJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModifyApsSlsADBJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyApsSlsADBJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyApsSlsADBJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyApsSlsADBJobResponseBody) SetCode(v string) *ModifyApsSlsADBJobResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) SetData(v map[string]interface{}) *ModifyApsSlsADBJobResponseBody {
	s.Data = v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) SetHttpStatusCode(v int32) *ModifyApsSlsADBJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) SetMessage(v string) *ModifyApsSlsADBJobResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) SetRequestId(v string) *ModifyApsSlsADBJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) SetSuccess(v bool) *ModifyApsSlsADBJobResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyApsSlsADBJobResponseBody) Validate() error {
	return dara.Validate(s)
}
