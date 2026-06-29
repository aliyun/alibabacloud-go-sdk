// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDatasetProxyAppendDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizCode(v string) *OpenDatasetProxyAppendDataRequest
	GetBizCode() *string
	SetDataMeta(v []map[string]*string) *OpenDatasetProxyAppendDataRequest
	GetDataMeta() []map[string]*string
	SetTaskId(v string) *OpenDatasetProxyAppendDataRequest
	GetTaskId() *string
	SetTraceId(v string) *OpenDatasetProxyAppendDataRequest
	GetTraceId() *string
	SetUUID(v string) *OpenDatasetProxyAppendDataRequest
	GetUUID() *string
}

type OpenDatasetProxyAppendDataRequest struct {
	BizCode *string `json:"BizCode,omitempty" xml:"BizCode,omitempty"`
	// A list of data records. A single invocation can contain up to 100 records. Each element in the array is a map.
	DataMeta []map[string]*string `json:"DataMeta,omitempty" xml:"DataMeta,omitempty" type:"Repeated"`
	// Task ID, indicating the task to which data is appended.
	//
	// example:
	//
	// 154***518306500608
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// TraceId
	//
	// example:
	//
	// 0bc1ec3916825622257033399edb6b
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
	// Unique identifier ID, controlled by the business side.
	//
	// example:
	//
	// e5c9db3f-f27c-445e-a52b-06ba6d1ba00f
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s OpenDatasetProxyAppendDataRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDatasetProxyAppendDataRequest) GoString() string {
	return s.String()
}

func (s *OpenDatasetProxyAppendDataRequest) GetBizCode() *string {
	return s.BizCode
}

func (s *OpenDatasetProxyAppendDataRequest) GetDataMeta() []map[string]*string {
	return s.DataMeta
}

func (s *OpenDatasetProxyAppendDataRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *OpenDatasetProxyAppendDataRequest) GetTraceId() *string {
	return s.TraceId
}

func (s *OpenDatasetProxyAppendDataRequest) GetUUID() *string {
	return s.UUID
}

func (s *OpenDatasetProxyAppendDataRequest) SetBizCode(v string) *OpenDatasetProxyAppendDataRequest {
	s.BizCode = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetDataMeta(v []map[string]*string) *OpenDatasetProxyAppendDataRequest {
	s.DataMeta = v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetTaskId(v string) *OpenDatasetProxyAppendDataRequest {
	s.TaskId = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetTraceId(v string) *OpenDatasetProxyAppendDataRequest {
	s.TraceId = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) SetUUID(v string) *OpenDatasetProxyAppendDataRequest {
	s.UUID = &v
	return s
}

func (s *OpenDatasetProxyAppendDataRequest) Validate() error {
	return dara.Validate(s)
}
