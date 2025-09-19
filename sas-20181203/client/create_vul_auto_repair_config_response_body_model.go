// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVulAutoRepairConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateVulAutoRepairConfigResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateVulAutoRepairConfigResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateVulAutoRepairConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateVulAutoRepairConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateVulAutoRepairConfigResponseBody
	GetSuccess() *bool
}

type CreateVulAutoRepairConfigResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6673D49C-A9AB-40DD-B4A2-B92306701AE7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateVulAutoRepairConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVulAutoRepairConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVulAutoRepairConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateVulAutoRepairConfigResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateVulAutoRepairConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateVulAutoRepairConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVulAutoRepairConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateVulAutoRepairConfigResponseBody) SetCode(v string) *CreateVulAutoRepairConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponseBody) SetHttpStatusCode(v int32) *CreateVulAutoRepairConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponseBody) SetMessage(v string) *CreateVulAutoRepairConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponseBody) SetRequestId(v string) *CreateVulAutoRepairConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponseBody) SetSuccess(v bool) *CreateVulAutoRepairConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateVulAutoRepairConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
