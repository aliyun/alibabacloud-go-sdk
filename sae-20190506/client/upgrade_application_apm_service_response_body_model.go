// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeApplicationApmServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpgradeApplicationApmServiceResponseBody
	GetCode() *string
	SetData(v *UpgradeApplicationApmServiceResponseBodyData) *UpgradeApplicationApmServiceResponseBody
	GetData() *UpgradeApplicationApmServiceResponseBodyData
	SetErrorCode(v string) *UpgradeApplicationApmServiceResponseBody
	GetErrorCode() *string
	SetMessage(v string) *UpgradeApplicationApmServiceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpgradeApplicationApmServiceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpgradeApplicationApmServiceResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *UpgradeApplicationApmServiceResponseBody
	GetTraceId() *string
}

type UpgradeApplicationApmServiceResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result.
	Data *UpgradeApplicationApmServiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The status code. Value values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	//
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message returned. The following limits are imposed on the ID:
	//
	// 	- If the request was successful, **success*	- is returned.
	//
	// 	- An error code is returned when a request failed.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. The ID is used to query the details of a request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s UpgradeApplicationApmServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationApmServiceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationApmServiceResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpgradeApplicationApmServiceResponseBody) GetData() *UpgradeApplicationApmServiceResponseBodyData {
	return s.Data
}

func (s *UpgradeApplicationApmServiceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpgradeApplicationApmServiceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpgradeApplicationApmServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeApplicationApmServiceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpgradeApplicationApmServiceResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *UpgradeApplicationApmServiceResponseBody) SetCode(v string) *UpgradeApplicationApmServiceResponseBody {
	s.Code = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetData(v *UpgradeApplicationApmServiceResponseBodyData) *UpgradeApplicationApmServiceResponseBody {
	s.Data = v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetErrorCode(v string) *UpgradeApplicationApmServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetMessage(v string) *UpgradeApplicationApmServiceResponseBody {
	s.Message = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetRequestId(v string) *UpgradeApplicationApmServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetSuccess(v bool) *UpgradeApplicationApmServiceResponseBody {
	s.Success = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) SetTraceId(v string) *UpgradeApplicationApmServiceResponseBody {
	s.TraceId = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpgradeApplicationApmServiceResponseBodyData struct {
	// Indicates whether ARMS advanced monitoring is enabled. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpgradeApplicationApmServiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpgradeApplicationApmServiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpgradeApplicationApmServiceResponseBodyData) GetStatus() *bool {
	return s.Status
}

func (s *UpgradeApplicationApmServiceResponseBodyData) SetStatus(v bool) *UpgradeApplicationApmServiceResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpgradeApplicationApmServiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}
