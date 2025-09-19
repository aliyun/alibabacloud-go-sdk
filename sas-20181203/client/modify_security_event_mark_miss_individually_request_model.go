// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityEventMarkMissIndividuallyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteMarkMissParam(v string) *ModifySecurityEventMarkMissIndividuallyRequest
	GetDeleteMarkMissParam() *string
	SetFrom(v string) *ModifySecurityEventMarkMissIndividuallyRequest
	GetFrom() *string
	SetInsertMarkMissParam(v string) *ModifySecurityEventMarkMissIndividuallyRequest
	GetInsertMarkMissParam() *string
	SetLang(v string) *ModifySecurityEventMarkMissIndividuallyRequest
	GetLang() *string
	SetSourceIp(v string) *ModifySecurityEventMarkMissIndividuallyRequest
	GetSourceIp() *string
}

type ModifySecurityEventMarkMissIndividuallyRequest struct {
	// The alert handling rule that you want to delete.
	//
	// example:
	//
	// [{\\"field\\":\\"loginSourceIp\\",\\"operate\\":\\"contains\\",\\"eventType\\":\\"SIL_AI_ALERT\\",\\"eventName\\":\\"login_common_ip\\",\\"fieldValue\\":\\"10.12.XX.XX\\",\\"uuids\\":\\"\\"}]
	DeleteMarkMissParam *string `json:"DeleteMarkMissParam,omitempty" xml:"DeleteMarkMissParam,omitempty"`
	// The ID of the request source. Set the value to **sas**.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The alert handling that you want to add.
	//
	// example:
	//
	// [{\\"field\\":\\"location\\",\\"operate\\":\\"contains\\",\\"eventType\\":\\"SIL_AI_ALERT\\",\\"eventName\\":\\"login_common_ip\\",\\"fieldValue\\":\\"xx\\",\\"uuids\\":\\"4296ee47-bf19-4fa4-a4a6-6bxxxxxxxxx\\"}]
	InsertMarkMissParam *string `json:"InsertMarkMissParam,omitempty" xml:"InsertMarkMissParam,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source IP address of the request. You do not need to specify this parameter. It is automatically obtained by the system.
	//
	// example:
	//
	// 127.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s ModifySecurityEventMarkMissIndividuallyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityEventMarkMissIndividuallyRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) GetDeleteMarkMissParam() *string {
	return s.DeleteMarkMissParam
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) GetFrom() *string {
	return s.From
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) GetInsertMarkMissParam() *string {
	return s.InsertMarkMissParam
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) GetLang() *string {
	return s.Lang
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) SetDeleteMarkMissParam(v string) *ModifySecurityEventMarkMissIndividuallyRequest {
	s.DeleteMarkMissParam = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) SetFrom(v string) *ModifySecurityEventMarkMissIndividuallyRequest {
	s.From = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) SetInsertMarkMissParam(v string) *ModifySecurityEventMarkMissIndividuallyRequest {
	s.InsertMarkMissParam = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) SetLang(v string) *ModifySecurityEventMarkMissIndividuallyRequest {
	s.Lang = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) SetSourceIp(v string) *ModifySecurityEventMarkMissIndividuallyRequest {
	s.SourceIp = &v
	return s
}

func (s *ModifySecurityEventMarkMissIndividuallyRequest) Validate() error {
	return dara.Validate(s)
}
