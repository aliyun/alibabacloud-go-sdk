// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleActiveSQLRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *HandleActiveSQLRecordResponseBody
	GetDBInstanceId() *string
	SetRequestId(v string) *HandleActiveSQLRecordResponseBody
	GetRequestId() *string
	SetResults(v []*HandleActiveSQLRecordResponseBodyResults) *HandleActiveSQLRecordResponseBody
	GetResults() []*HandleActiveSQLRecordResponseBodyResults
	SetStatus(v string) *HandleActiveSQLRecordResponseBody
	GetStatus() *string
}

type HandleActiveSQLRecordResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The processing result of the active query.
	Results []*HandleActiveSQLRecordResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s HandleActiveSQLRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HandleActiveSQLRecordResponseBody) GoString() string {
	return s.String()
}

func (s *HandleActiveSQLRecordResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *HandleActiveSQLRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HandleActiveSQLRecordResponseBody) GetResults() []*HandleActiveSQLRecordResponseBodyResults {
	return s.Results
}

func (s *HandleActiveSQLRecordResponseBody) GetStatus() *string {
	return s.Status
}

func (s *HandleActiveSQLRecordResponseBody) SetDBInstanceId(v string) *HandleActiveSQLRecordResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *HandleActiveSQLRecordResponseBody) SetRequestId(v string) *HandleActiveSQLRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *HandleActiveSQLRecordResponseBody) SetResults(v []*HandleActiveSQLRecordResponseBodyResults) *HandleActiveSQLRecordResponseBody {
	s.Results = v
	return s
}

func (s *HandleActiveSQLRecordResponseBody) SetStatus(v string) *HandleActiveSQLRecordResponseBody {
	s.Status = &v
	return s
}

func (s *HandleActiveSQLRecordResponseBody) Validate() error {
	return dara.Validate(s)
}

type HandleActiveSQLRecordResponseBodyResults struct {
	// The process ID, which is a unique identifier of the query.
	//
	// example:
	//
	// 3003925
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// Indicates whether the processing was successful. Valid values:
	//
	// 	- **false**
	//
	// 	- **true**
	//
	// example:
	//
	// true
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s HandleActiveSQLRecordResponseBodyResults) String() string {
	return dara.Prettify(s)
}

func (s HandleActiveSQLRecordResponseBodyResults) GoString() string {
	return s.String()
}

func (s *HandleActiveSQLRecordResponseBodyResults) GetPid() *string {
	return s.Pid
}

func (s *HandleActiveSQLRecordResponseBodyResults) GetStatus() *string {
	return s.Status
}

func (s *HandleActiveSQLRecordResponseBodyResults) SetPid(v string) *HandleActiveSQLRecordResponseBodyResults {
	s.Pid = &v
	return s
}

func (s *HandleActiveSQLRecordResponseBodyResults) SetStatus(v string) *HandleActiveSQLRecordResponseBodyResults {
	s.Status = &v
	return s
}

func (s *HandleActiveSQLRecordResponseBodyResults) Validate() error {
	return dara.Validate(s)
}
