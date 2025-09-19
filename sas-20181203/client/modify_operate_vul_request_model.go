// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyOperateVulRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFrom(v string) *ModifyOperateVulRequest
	GetFrom() *string
	SetInfo(v string) *ModifyOperateVulRequest
	GetInfo() *string
	SetOperateType(v string) *ModifyOperateVulRequest
	GetOperateType() *string
	SetReason(v string) *ModifyOperateVulRequest
	GetReason() *string
	SetType(v string) *ModifyOperateVulRequest
	GetType() *string
}

type ModifyOperateVulRequest struct {
	// The request ID. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The details of the vulnerability. The value of this parameter is in the JSON format and contains the following fields:
	//
	// 	- **name**: the name of the vulnerability.
	//
	// 	- **uuid**: the UUID of the server on which the vulnerability is detected.
	//
	// 	- **tag**: the tag that is added to the vulnerability. Valid values:
	//
	//     	- **oval**: Linux software vulnerability
	//
	//     	- **system**: Windows system vulnerability
	//
	//     	- **cms**: Web-CMS vulnerability
	//
	// >  You can call the [DescribeVulList](~~DescribeVulList~~) operation to query the tags that are added to vulnerabilities of other types.
	//
	// 	- **isFront**: specifies whether a pre-patch is required to fix the Windows system vulnerability. This field is required only for Windows system vulnerabilities. Valid values:
	//
	//     	- **0**: no
	//
	//     	- **1**: yes
	//
	// >  You can fix multiple vulnerabilities at a time. Separate the details of multiple vulnerabilities with commas (,). You can call the [DescribeVulLIst](~~DescribeVulList~~) operation to query the details of vulnerabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"alilinux2:2.1903:ALINUX2-SA-2022:0007","uuid":"a3bb82a8-a3bd-4546-acce-45ac34af****","tag":"oval","isFront":0},{"name":"alilinux2:2.1903:ALINUX2-SA-2022:0007","uuid":"98a6fecc-88cd-46f2-8e35-f808a388****","tag":"oval","isFront":0}]
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// The operation that you want to perform on the vulnerability. Valid values:
	//
	// 	- **vul_fix**: fixes the vulnerability.
	//
	// 	- **vul_verify**: verifies the vulnerability fix.
	//
	// 	- **vul_ignore**: ignores the vulnerability.
	//
	// 	- **vul_undo_ignore**: cancels ignoring the vulnerability.
	//
	// 	- **vul_delete**: deletes the vulnerability.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The reason why the vulnerability is **ignored**.
	//
	// >  This parameter is required only when you set **OperateType*	- to **vul_ignore**.
	//
	// example:
	//
	// not operate
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// 	- **app**: application vulnerability
	//
	// 	- **sca**: vulnerability that is detected based on software component analysis
	//
	// >  You cannot fix the urgent vulnerabilities, application vulnerabilities, or vulnerabilities that are detected based on software component analysis.
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

func (s *ModifyOperateVulRequest) GetType() *string {
	return s.Type
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

func (s *ModifyOperateVulRequest) SetType(v string) *ModifyOperateVulRequest {
	s.Type = &v
	return s
}

func (s *ModifyOperateVulRequest) Validate() error {
	return dara.Validate(s)
}
