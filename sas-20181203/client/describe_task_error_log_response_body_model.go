// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTaskErrorLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*DescribeTaskErrorLogResponseBodyLogs) *DescribeTaskErrorLogResponseBody
	GetLogs() []*DescribeTaskErrorLogResponseBodyLogs
	SetRequestId(v string) *DescribeTaskErrorLogResponseBody
	GetRequestId() *string
}

type DescribeTaskErrorLogResponseBody struct {
	// An array that consists of the error logs.
	Logs []*DescribeTaskErrorLogResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F929E952-EBFC-56C3-BD35-BF8B59024C69
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeTaskErrorLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskErrorLogResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskErrorLogResponseBody) GetLogs() []*DescribeTaskErrorLogResponseBodyLogs {
	return s.Logs
}

func (s *DescribeTaskErrorLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskErrorLogResponseBody) SetLogs(v []*DescribeTaskErrorLogResponseBodyLogs) *DescribeTaskErrorLogResponseBody {
	s.Logs = v
	return s
}

func (s *DescribeTaskErrorLogResponseBody) SetRequestId(v string) *DescribeTaskErrorLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskErrorLogResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTaskErrorLogResponseBodyLogs struct {
	// The text content of the log.
	//
	// example:
	//
	// mv: cannot move \\"CentOS-Base.repo\\" to \\"CentOS-Base.repo.backup\\": Permission denied
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DescribeTaskErrorLogResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskErrorLogResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *DescribeTaskErrorLogResponseBodyLogs) GetText() *string {
	return s.Text
}

func (s *DescribeTaskErrorLogResponseBodyLogs) SetText(v string) *DescribeTaskErrorLogResponseBodyLogs {
	s.Text = &v
	return s
}

func (s *DescribeTaskErrorLogResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}
