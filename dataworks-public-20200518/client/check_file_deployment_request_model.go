// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckFileDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckDetailUrl(v string) *CheckFileDeploymentRequest
	GetCheckDetailUrl() *string
	SetCheckerInstanceId(v string) *CheckFileDeploymentRequest
	GetCheckerInstanceId() *string
	SetStatus(v string) *CheckFileDeploymentRequest
	GetStatus() *string
}

type CheckFileDeploymentRequest struct {
	// Deprecated.
	//
	// example:
	//
	// https://result.aliyun.com/?checkerInstanceId=
	CheckDetailUrl *string `json:"CheckDetailUrl,omitempty" xml:"CheckDetailUrl,omitempty"`
	// The instance ID to which the file checker belongs. You can obtain this value from the CheckerInstanceId field in the file publish check event.
	//
	// This parameter is required.
	//
	// example:
	//
	// 66_123455623_2
	CheckerInstanceId *string `json:"CheckerInstanceId,omitempty" xml:"CheckerInstanceId,omitempty"`
	// The check status of the file pending deployment. Valid values:
	//
	// - OK: The file passed the check.
	//
	// - WARN: The file passed the check but has warnings.
	//
	// - FAIL: The file failed the check.
	//
	// This parameter is required.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckFileDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s CheckFileDeploymentRequest) GoString() string {
	return s.String()
}

func (s *CheckFileDeploymentRequest) GetCheckDetailUrl() *string {
	return s.CheckDetailUrl
}

func (s *CheckFileDeploymentRequest) GetCheckerInstanceId() *string {
	return s.CheckerInstanceId
}

func (s *CheckFileDeploymentRequest) GetStatus() *string {
	return s.Status
}

func (s *CheckFileDeploymentRequest) SetCheckDetailUrl(v string) *CheckFileDeploymentRequest {
	s.CheckDetailUrl = &v
	return s
}

func (s *CheckFileDeploymentRequest) SetCheckerInstanceId(v string) *CheckFileDeploymentRequest {
	s.CheckerInstanceId = &v
	return s
}

func (s *CheckFileDeploymentRequest) SetStatus(v string) *CheckFileDeploymentRequest {
	s.Status = &v
	return s
}

func (s *CheckFileDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
