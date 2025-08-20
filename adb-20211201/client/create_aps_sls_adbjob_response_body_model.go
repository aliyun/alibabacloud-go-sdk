// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApsSlsADBJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateApsSlsADBJobResponseBody
	GetCode() *string
	SetData(v string) *CreateApsSlsADBJobResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateApsSlsADBJobResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateApsSlsADBJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateApsSlsADBJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateApsSlsADBJobResponseBody
	GetSuccess() *bool
}

type CreateApsSlsADBJobResponseBody struct {
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
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
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
	// Indicates whether the dry run succeeds. Valid values:
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

func (s CreateApsSlsADBJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApsSlsADBJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApsSlsADBJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateApsSlsADBJobResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateApsSlsADBJobResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateApsSlsADBJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateApsSlsADBJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApsSlsADBJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateApsSlsADBJobResponseBody) SetCode(v string) *CreateApsSlsADBJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) SetData(v string) *CreateApsSlsADBJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) SetHttpStatusCode(v int32) *CreateApsSlsADBJobResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) SetMessage(v string) *CreateApsSlsADBJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) SetRequestId(v string) *CreateApsSlsADBJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) SetSuccess(v bool) *CreateApsSlsADBJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateApsSlsADBJobResponseBody) Validate() error {
	return dara.Validate(s)
}
