// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsCopyWorkloadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsCopyWorkloadResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateApsCopyWorkloadResponseBody
	GetData() map[string]interface{}
	SetHttpStatusCode(v int32) *CreateApsCopyWorkloadResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApsCopyWorkloadResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsCopyWorkloadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsCopyWorkloadResponseBody
	GetSuccess() *bool
}

type CreateApsCopyWorkloadResponseBody struct {
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
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
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

func (s CreateApsCopyWorkloadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsCopyWorkloadResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsCopyWorkloadResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsCopyWorkloadResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateApsCopyWorkloadResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApsCopyWorkloadResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsCopyWorkloadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsCopyWorkloadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsCopyWorkloadResponseBody) SetCode(v string) *CreateApsCopyWorkloadResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) SetData(v map[string]interface{}) *CreateApsCopyWorkloadResponseBody {
	s.Data = v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) SetHttpStatusCode(v int32) *CreateApsCopyWorkloadResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) SetMessage(v string) *CreateApsCopyWorkloadResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) SetRequestId(v string) *CreateApsCopyWorkloadResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) SetSuccess(v bool) *CreateApsCopyWorkloadResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsCopyWorkloadResponseBody) Validate() error {
	return dara.Validate(s)
}
