// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListGpuClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetListGpuClusterResponseBody
	GetCode() *string
	SetData(v interface{}) *GetListGpuClusterResponseBody
	GetData() interface{}
	SetMessage(v string) *GetListGpuClusterResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetListGpuClusterResponseBody
	GetRequestId() *string
	SetTotal(v int32) *GetListGpuClusterResponseBody
	GetTotal() *int32
}

type GetListGpuClusterResponseBody struct {
	// The code. Valid values: Success and Sysom.ACKClusterNotFound.
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of clusters.
	//
	// example:
	//
	// ["cluster1", "cluster2"]
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
	// The error message. If the code is not success, the message is the error information.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 10
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s GetListGpuClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetListGpuClusterResponseBody) GoString() string {
	return s.String()
}

func (s *GetListGpuClusterResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetListGpuClusterResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetListGpuClusterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetListGpuClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetListGpuClusterResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetListGpuClusterResponseBody) SetCode(v string) *GetListGpuClusterResponseBody {
	s.Code = &v
	return s
}

func (s *GetListGpuClusterResponseBody) SetData(v interface{}) *GetListGpuClusterResponseBody {
	s.Data = v
	return s
}

func (s *GetListGpuClusterResponseBody) SetMessage(v string) *GetListGpuClusterResponseBody {
	s.Message = &v
	return s
}

func (s *GetListGpuClusterResponseBody) SetRequestId(v string) *GetListGpuClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetListGpuClusterResponseBody) SetTotal(v int32) *GetListGpuClusterResponseBody {
	s.Total = &v
	return s
}

func (s *GetListGpuClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
