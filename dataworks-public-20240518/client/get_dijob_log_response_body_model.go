// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDIJobLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLog(v string) *GetDIJobLogResponseBody
	GetLog() *string
	SetRequestId(v string) *GetDIJobLogResponseBody
	GetRequestId() *string
}

type GetDIJobLogResponseBody struct {
	// The log.
	//
	// example:
	//
	// >>>>>>>> stdout:n++++++++++++++++++executing sql: create database if not exists jindo_test location \\"oss://pangbei-hdfs/tmp/hive\\" n++n
	Log *string `json:"Log,omitempty" xml:"Log,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 1AFAE64E-D1BE-432B-A9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDIJobLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDIJobLogResponseBody) GoString() string {
	return s.String()
}

func (s *GetDIJobLogResponseBody) GetLog() *string {
	return s.Log
}

func (s *GetDIJobLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDIJobLogResponseBody) SetLog(v string) *GetDIJobLogResponseBody {
	s.Log = &v
	return s
}

func (s *GetDIJobLogResponseBody) SetRequestId(v string) *GetDIJobLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDIJobLogResponseBody) Validate() error {
	return dara.Validate(s)
}
