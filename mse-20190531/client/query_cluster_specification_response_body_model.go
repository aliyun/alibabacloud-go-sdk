// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterSpecificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryClusterSpecificationResponseBody
	GetCode() *int32
	SetData(v []*QueryClusterSpecificationResponseBodyData) *QueryClusterSpecificationResponseBody
	GetData() []*QueryClusterSpecificationResponseBodyData
	SetErrorCode(v string) *QueryClusterSpecificationResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *QueryClusterSpecificationResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *QueryClusterSpecificationResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryClusterSpecificationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryClusterSpecificationResponseBody
	GetSuccess() *bool
}

type QueryClusterSpecificationResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The details of the data.
	Data []*QueryClusterSpecificationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EE5C32A1-BC0E-4B79-817C-103E4EDF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryClusterSpecificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterSpecificationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryClusterSpecificationResponseBody) GetData() []*QueryClusterSpecificationResponseBodyData {
	return s.Data
}

func (s *QueryClusterSpecificationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *QueryClusterSpecificationResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *QueryClusterSpecificationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryClusterSpecificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryClusterSpecificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryClusterSpecificationResponseBody) SetCode(v int32) *QueryClusterSpecificationResponseBody {
	s.Code = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetData(v []*QueryClusterSpecificationResponseBodyData) *QueryClusterSpecificationResponseBody {
	s.Data = v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetErrorCode(v string) *QueryClusterSpecificationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetHttpStatusCode(v int32) *QueryClusterSpecificationResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetMessage(v string) *QueryClusterSpecificationResponseBody {
	s.Message = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetRequestId(v string) *QueryClusterSpecificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) SetSuccess(v bool) *QueryClusterSpecificationResponseBody {
	s.Success = &v
	return s
}

func (s *QueryClusterSpecificationResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryClusterSpecificationResponseBodyData struct {
	// The engine specifications that can be used.
	//
	// example:
	//
	// MSE_SC_1_2_200_c
	ClusterSpecificationName *string `json:"ClusterSpecificationName,omitempty" xml:"ClusterSpecificationName,omitempty"`
	// The number of vCPUs in the specifications.
	//
	// example:
	//
	// 1
	CpuCapacity *string `json:"CpuCapacity,omitempty" xml:"CpuCapacity,omitempty"`
	// The memory size in the specifications. Unit: GB.
	//
	// example:
	//
	// 2
	MemoryCapacity *string `json:"MemoryCapacity,omitempty" xml:"MemoryCapacity,omitempty"`
}

func (s QueryClusterSpecificationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterSpecificationResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryClusterSpecificationResponseBodyData) GetClusterSpecificationName() *string {
	return s.ClusterSpecificationName
}

func (s *QueryClusterSpecificationResponseBodyData) GetCpuCapacity() *string {
	return s.CpuCapacity
}

func (s *QueryClusterSpecificationResponseBodyData) GetMemoryCapacity() *string {
	return s.MemoryCapacity
}

func (s *QueryClusterSpecificationResponseBodyData) SetClusterSpecificationName(v string) *QueryClusterSpecificationResponseBodyData {
	s.ClusterSpecificationName = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetCpuCapacity(v string) *QueryClusterSpecificationResponseBodyData {
	s.CpuCapacity = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) SetMemoryCapacity(v string) *QueryClusterSpecificationResponseBodyData {
	s.MemoryCapacity = &v
	return s
}

func (s *QueryClusterSpecificationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
