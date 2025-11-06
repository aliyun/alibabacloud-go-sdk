// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDiskSpecificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryClusterDiskSpecificationResponseBody
	GetCode() *int32
	SetData(v *QueryClusterDiskSpecificationResponseBodyData) *QueryClusterDiskSpecificationResponseBody
	GetData() *QueryClusterDiskSpecificationResponseBodyData
	SetDynamicMessage(v string) *QueryClusterDiskSpecificationResponseBody
	GetDynamicMessage() *string
	SetErrorCode(v string) *QueryClusterDiskSpecificationResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *QueryClusterDiskSpecificationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryClusterDiskSpecificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryClusterDiskSpecificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryClusterDiskSpecificationResponseBody
	GetSuccess() *bool
}

type QueryClusterDiskSpecificationResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data *QueryClusterDiskSpecificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The dynamic part in the error message. This parameter is used to replace the \\*\\*%s\\*\\	- variable in the **ErrMessage*	- parameter.
	//
	// >  If the return value of the **ErrMessage*	- parameter is **The Value of Input Parameter %s is not valid*	- and the return value of the **DynamicMessage*	- parameter is **DtsJobId**, the specified **DtsJobId*	- parameter is invalid.
	//
	// example:
	//
	// The specified parameter is invalid.
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E13A3A34-7201-50C4-B2D0-0D7DB891811E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterDiskSpecificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryClusterDiskSpecificationResponseBody) GetData() *QueryClusterDiskSpecificationResponseBodyData {
	return s.Data
}

func (s *QueryClusterDiskSpecificationResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *QueryClusterDiskSpecificationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryClusterDiskSpecificationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryClusterDiskSpecificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryClusterDiskSpecificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryClusterDiskSpecificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryClusterDiskSpecificationResponseBody) SetCode(v int32) *QueryClusterDiskSpecificationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetData(v *QueryClusterDiskSpecificationResponseBodyData) *QueryClusterDiskSpecificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetDynamicMessage(v string) *QueryClusterDiskSpecificationResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetErrorCode(v string) *QueryClusterDiskSpecificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetHttpStatusCode(v int32) *QueryClusterDiskSpecificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetMessage(v string) *QueryClusterDiskSpecificationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetRequestId(v string) *QueryClusterDiskSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) SetSuccess(v bool) *QueryClusterDiskSpecificationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryClusterDiskSpecificationResponseBodyData struct {
	// The maximum disk capacity. Unit: GB.
	//
	// example:
	//
	// 500
	Max *int32 `json:"Max,omitempty" xml:"Max,omitempty"`
	// The minimum disk capacity. Unit: GB.
	//
	// example:
	//
	// 1
	Min *int32 `json:"Min,omitempty" xml:"Min,omitempty"`
	// The step size of the disk capacity.
	//
	// example:
	//
	// 2
	Step *int32 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s QueryClusterDiskSpecificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDiskSpecificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterDiskSpecificationResponseBodyData) GetMax() *int32 {
	return s.Max
}

func (s *QueryClusterDiskSpecificationResponseBodyData) GetMin() *int32 {
	return s.Min
}

func (s *QueryClusterDiskSpecificationResponseBodyData) GetStep() *int32 {
	return s.Step
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetMax(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Max = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetMin(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Min = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBodyData) SetStep(v int32) *QueryClusterDiskSpecificationResponseBodyData {
	s.Step = &v
	return s
}

func (s *QueryClusterDiskSpecificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
