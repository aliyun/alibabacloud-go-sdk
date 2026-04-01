// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAttachScriptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeClusterAttachScriptsResponseBody
	GetRequestId() *string
	SetScript(v string) *DescribeClusterAttachScriptsResponseBody
	GetScript() *string
}

type DescribeClusterAttachScriptsResponseBody struct {
	// example:
	//
	// CEF72CEB-54B6-4AE8-B225-F876FF7BA984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Y2QgWnhiU0o=
	Script *string `json:"Script,omitempty" xml:"Script,omitempty"`
}

func (s DescribeClusterAttachScriptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAttachScriptsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAttachScriptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterAttachScriptsResponseBody) GetScript() *string {
	return s.Script
}

func (s *DescribeClusterAttachScriptsResponseBody) SetRequestId(v string) *DescribeClusterAttachScriptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterAttachScriptsResponseBody) SetScript(v string) *DescribeClusterAttachScriptsResponseBody {
	s.Script = &v
	return s
}

func (s *DescribeClusterAttachScriptsResponseBody) Validate() error {
	return dara.Validate(s)
}
