// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRDSToClickhouseDbResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMsg(v string) *CreateRDSToClickhouseDbResponseBody
	GetErrorMsg() *string
	SetRepeatedDbs(v []*string) *CreateRDSToClickhouseDbResponseBody
	GetRepeatedDbs() []*string
	SetRequestId(v string) *CreateRDSToClickhouseDbResponseBody
	GetRequestId() *string
	SetStatus(v int64) *CreateRDSToClickhouseDbResponseBody
	GetStatus() *int64
}

type CreateRDSToClickhouseDbResponseBody struct {
	// If the value of the **Status*	- parameter is -1, the cause of the creation failure is returned.
	//
	// example:
	//
	// ClickHouse exception, code: 49, host: 100.100.79.102, port: 14540; Code: 49, e.displayText() = DB::Exception: Logical error: there is no global context (version 20.8.17.25)n
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// Duplicate tables in the synchronization task.
	RepeatedDbs []*string `json:"RepeatedDbs,omitempty" xml:"RepeatedDbs,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 66676F54-1994-5DCF-993F-74536649628A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the synchronization task was created. Valid values:
	//
	// 	- **1**: Created.
	//
	// 	- **0**: Creation failed. The tables in the synchronization task are duplicate. The duplicate tables are returned for the **RepeatedDbs*	- parameter.
	//
	// 	- **-1**: Creation failed. The cause why the creation failed is returned for the **ErrorMsg*	- parameter.
	//
	// example:
	//
	// 0
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRDSToClickhouseDbResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRDSToClickhouseDbResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRDSToClickhouseDbResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateRDSToClickhouseDbResponseBody) GetRepeatedDbs() []*string {
	return s.RepeatedDbs
}

func (s *CreateRDSToClickhouseDbResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRDSToClickhouseDbResponseBody) GetStatus() *int64 {
	return s.Status
}

func (s *CreateRDSToClickhouseDbResponseBody) SetErrorMsg(v string) *CreateRDSToClickhouseDbResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetRepeatedDbs(v []*string) *CreateRDSToClickhouseDbResponseBody {
	s.RepeatedDbs = v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetRequestId(v string) *CreateRDSToClickhouseDbResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) SetStatus(v int64) *CreateRDSToClickhouseDbResponseBody {
	s.Status = &v
	return s
}

func (s *CreateRDSToClickhouseDbResponseBody) Validate() error {
	return dara.Validate(s)
}
