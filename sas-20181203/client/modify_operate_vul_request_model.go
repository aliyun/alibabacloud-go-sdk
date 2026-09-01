// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOperateVulRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyOperateVulRequest
	GetClientToken() *string
	SetFrom(v string) *ModifyOperateVulRequest
	GetFrom() *string
	SetInfo(v string) *ModifyOperateVulRequest
	GetInfo() *string
	SetOperateType(v string) *ModifyOperateVulRequest
	GetOperateType() *string
	SetReason(v string) *ModifyOperateVulRequest
	GetReason() *string
	SetResourceDirectoryAccountId(v int64) *ModifyOperateVulRequest
	GetResourceDirectoryAccountId() *int64
	SetType(v string) *ModifyOperateVulRequest
	GetType() *string
}

type ModifyOperateVulRequest struct {
	// The client token that is used to ensure the idempotence of the request. Use a different token for each request. The token supports only ASCII characters and cannot exceed 64 characters in length.
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The source identifier of the request. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The information about the vulnerability to handle. This parameter is in JSON format and contains the following fields:
	//
	// - **name**: The name of the vulnerability.
	//
	// - **uuid**: The UUID of the server on which the vulnerability is detected.
	//
	// - **tag**: The tag of the vulnerability. Valid values:
	//
	//     - **oval**: Linux software vulnerability.
	//
	//     - **system**: Windows system vulnerability.
	//
	//     - **cms**: Web-CMS vulnerability.
	//
	// > For other vulnerability types, call the [DescribeVulList](~~DescribeVulList~~) operation to obtain vulnerability information.
	//
	// - **isFront**: Specifies whether the Windows patch is a prerequisite patch. This parameter is required only when you handle Windows system vulnerabilities. You can ignore this parameter for other vulnerability types. Valid values:
	//
	//     - **0**: No.
	//
	//     - **1**: Yes.
	//
	// > Batch processing of vulnerabilities is supported. Separate multiple vulnerability entries with commas (,). Call the [DescribeVulList](~~DescribeVulList~~) operation to obtain vulnerability information.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"alilinux2:2.1903:ALINUX2-SA-2022:0007","uuid":"a3bb82a8-a3bd-4546-acce-45ac34af****","tag":"oval","isFront":0},{"name":"alilinux2:2.1903:ALINUX2-SA-2022:0007","uuid":"98a6fecc-88cd-46f2-8e35-f808a388****","tag":"oval","isFront":0}]
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The operation to perform on the vulnerability. Valid values:
	//
	// - **vul_fix**: fixes the vulnerability.
	//
	// - **vul_verify**: verifies the vulnerability.
	//
	// - **vul_ignore**: ignores the vulnerability.
	//
	// - **vul_undo_ignore**: cancels ignoring the vulnerability.
	//
	// - **vul_delete**: deletes the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The reason for ignoring the vulnerability.
	//
	// > This parameter is required only when the operation type is **ignore*	- (OperateType is set to **vul_ignore**).
	//
	// example:
	//
	// not operate
	Reason                     *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	ResourceDirectoryAccountId *int64  `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The type of the vulnerability to handle. Valid values:
	//
	// - **cve**: Linux software vulnerability.
	//
	// - **sys**: Windows system vulnerability.
	//
	// - **cms**: Web-CMS vulnerability.
	//
	// - **emg**: emergency vulnerability.
	//
	// - **app**: application vulnerability.
	//
	// - **sca**: software constituency parsing vulnerability.
	//
	// > Emergency vulnerabilities (emg), application vulnerabilities (app), and software constituency parsing vulnerabilities (sca) do not support the execute vulnerability fix operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyOperateVulRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyOperateVulRequest) GoString() string {
	return s.String()
}

func (s *ModifyOperateVulRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyOperateVulRequest) GetFrom() *string {
	return s.From
}

func (s *ModifyOperateVulRequest) GetInfo() *string {
	return s.Info
}

func (s *ModifyOperateVulRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *ModifyOperateVulRequest) GetReason() *string {
	return s.Reason
}

func (s *ModifyOperateVulRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ModifyOperateVulRequest) GetType() *string {
	return s.Type
}

func (s *ModifyOperateVulRequest) SetClientToken(v string) *ModifyOperateVulRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyOperateVulRequest) SetFrom(v string) *ModifyOperateVulRequest {
	s.From = &v
	return s
}

func (s *ModifyOperateVulRequest) SetInfo(v string) *ModifyOperateVulRequest {
	s.Info = &v
	return s
}

func (s *ModifyOperateVulRequest) SetOperateType(v string) *ModifyOperateVulRequest {
	s.OperateType = &v
	return s
}

func (s *ModifyOperateVulRequest) SetReason(v string) *ModifyOperateVulRequest {
	s.Reason = &v
	return s
}

func (s *ModifyOperateVulRequest) SetResourceDirectoryAccountId(v int64) *ModifyOperateVulRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) Validate() error {
	return dara.Validate(s)
}
