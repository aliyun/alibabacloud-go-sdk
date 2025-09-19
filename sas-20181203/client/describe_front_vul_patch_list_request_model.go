// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFrontVulPatchListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInfo(v string) *DescribeFrontVulPatchListRequest
	GetInfo() *string
	SetLang(v string) *DescribeFrontVulPatchListRequest
	GetLang() *string
	SetOperateType(v string) *DescribeFrontVulPatchListRequest
	GetOperateType() *string
	SetType(v string) *DescribeFrontVulPatchListRequest
	GetType() *string
}

type DescribeFrontVulPatchListRequest struct {
	// The information about the Windows system vulnerability. The value is a JSON string that contains the following fields:
	//
	// 	- **name**: the name of the vulnerability.
	//
	// 	- **uuid**: the UUID of the server on which the vulnerability is detected.
	//
	// 	- **tag**: the tag that is added to the vulnerability. Set this field to **system**, which indicates Windows system vulnerabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"name":"5000803","uuid":"026c9296-1234-5678-b937-a7d81f05****","tag":"system"}]
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
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
	// The operation that you want to perform on the vulnerability. Set the value to **vul_fix**, which indicates vulnerability fixing.
	//
	// This parameter is required.
	//
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
	// The type of the vulnerability. Set the value to **sys**, which indicates Windows system vulnerabilities.
	//
	// This parameter is required.
	//
	// example:
	//
	// sys
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeFrontVulPatchListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFrontVulPatchListRequest) GoString() string {
	return s.String()
}

func (s *DescribeFrontVulPatchListRequest) GetInfo() *string {
	return s.Info
}

func (s *DescribeFrontVulPatchListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeFrontVulPatchListRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *DescribeFrontVulPatchListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeFrontVulPatchListRequest) SetInfo(v string) *DescribeFrontVulPatchListRequest {
	s.Info = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetLang(v string) *DescribeFrontVulPatchListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetOperateType(v string) *DescribeFrontVulPatchListRequest {
	s.OperateType = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) SetType(v string) *DescribeFrontVulPatchListRequest {
	s.Type = &v
	return s
}

func (s *DescribeFrontVulPatchListRequest) Validate() error {
	return dara.Validate(s)
}
