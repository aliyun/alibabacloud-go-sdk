// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserWafLogStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStoreName(v string) *ModifyUserWafLogStatusResponseBody
	GetLogStoreName() *string
	SetProjectName(v bool) *ModifyUserWafLogStatusResponseBody
	GetProjectName() *bool
	SetRequestId(v string) *ModifyUserWafLogStatusResponseBody
	GetRequestId() *string
}

type ModifyUserWafLogStatusResponseBody struct {
	// example:
	//
	// wafng-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// wafng-project-14316572********-cn-hangzhou
	ProjectName *bool `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// D7861F61-5B61-****-A47C-6B19160*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserWafLogStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserWafLogStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserWafLogStatusResponseBody) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ModifyUserWafLogStatusResponseBody) GetProjectName() *bool {
	return s.ProjectName
}

func (s *ModifyUserWafLogStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserWafLogStatusResponseBody) SetLogStoreName(v string) *ModifyUserWafLogStatusResponseBody {
	s.LogStoreName = &v
	return s
}

func (s *ModifyUserWafLogStatusResponseBody) SetProjectName(v bool) *ModifyUserWafLogStatusResponseBody {
	s.ProjectName = &v
	return s
}

func (s *ModifyUserWafLogStatusResponseBody) SetRequestId(v string) *ModifyUserWafLogStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserWafLogStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
