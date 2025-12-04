// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckModifyConfigNeedRestartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNeedRestart(v bool) *CheckModifyConfigNeedRestartResponseBody
	GetNeedRestart() *bool
	SetRequestId(v string) *CheckModifyConfigNeedRestartResponseBody
	GetRequestId() *string
}

type CheckModifyConfigNeedRestartResponseBody struct {
	// Indicates whether the cluster was restarted after you modified the configuration parameters. Valid values:
	//
	// 	- **true**: The cluster was restarted.
	//
	// 	- **false**: The cluster was not restarted.
	//
	// example:
	//
	// true
	NeedRestart *bool `json:"NeedRestart,omitempty" xml:"NeedRestart,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06798FEE-BEF2-5FAF-A30D-728973BBE97C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckModifyConfigNeedRestartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckModifyConfigNeedRestartResponseBody) GoString() string {
	return s.String()
}

func (s *CheckModifyConfigNeedRestartResponseBody) GetNeedRestart() *bool {
	return s.NeedRestart
}

func (s *CheckModifyConfigNeedRestartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckModifyConfigNeedRestartResponseBody) SetNeedRestart(v bool) *CheckModifyConfigNeedRestartResponseBody {
	s.NeedRestart = &v
	return s
}

func (s *CheckModifyConfigNeedRestartResponseBody) SetRequestId(v string) *CheckModifyConfigNeedRestartResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckModifyConfigNeedRestartResponseBody) Validate() error {
	return dara.Validate(s)
}
