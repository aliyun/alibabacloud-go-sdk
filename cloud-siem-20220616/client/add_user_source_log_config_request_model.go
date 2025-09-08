// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserSourceLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleted(v int32) *AddUserSourceLogConfigRequest
	GetDeleted() *int32
	SetDisPlayLine(v string) *AddUserSourceLogConfigRequest
	GetDisPlayLine() *string
	SetRegionId(v string) *AddUserSourceLogConfigRequest
	GetRegionId() *string
	SetSourceLogCode(v string) *AddUserSourceLogConfigRequest
	GetSourceLogCode() *string
	SetSourceLogInfo(v string) *AddUserSourceLogConfigRequest
	GetSourceLogInfo() *string
	SetSourceProdCode(v string) *AddUserSourceLogConfigRequest
	GetSourceProdCode() *string
	SetSubUserId(v int64) *AddUserSourceLogConfigRequest
	GetSubUserId() *int64
}

type AddUserSourceLogConfigRequest struct {
	// Specifies whether to add logs or delete added logs. Valid values:
	//
	// 	- \\-1: deletes added logs.
	//
	// 	- 0: adds logs.
	//
	// example:
	//
	// 0
	Deleted *int32 `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	// The display details of the Logstore.
	//
	// example:
	//
	// cn-shanghai.siem-project.siem-logstore
	DisPlayLine *string `json:"DisPlayLine,omitempty" xml:"DisPlayLine,omitempty"`
	// The region in which the data management center of the threat analysis feature resides. Specify this parameter based on the regions in which your assets reside. Valid values:
	//
	// 	- cn-hangzhou: Your assets reside in regions in China.
	//
	// 	- ap-southeast-1: Your assets reside in regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The log code.
	//
	// example:
	//
	// cloud_siem_aegis_proc
	SourceLogCode *string `json:"SourceLogCode,omitempty" xml:"SourceLogCode,omitempty"`
	// The details of the Logstore that you want to use in the JSON string format.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"project":"wafnew-project-1335759343513432-cn-hangzhou","logStore":"wafnew-logstore","regionCode":"cn-hangzhou","prodCode":"waf"}
	SourceLogInfo *string `json:"SourceLogInfo,omitempty" xml:"SourceLogInfo,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// sas
	SourceProdCode *string `json:"SourceProdCode,omitempty" xml:"SourceProdCode,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123XXXXXX
	SubUserId *int64 `json:"SubUserId,omitempty" xml:"SubUserId,omitempty"`
}

func (s AddUserSourceLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s AddUserSourceLogConfigRequest) GoString() string {
	return s.String()
}

func (s *AddUserSourceLogConfigRequest) GetDeleted() *int32 {
	return s.Deleted
}

func (s *AddUserSourceLogConfigRequest) GetDisPlayLine() *string {
	return s.DisPlayLine
}

func (s *AddUserSourceLogConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddUserSourceLogConfigRequest) GetSourceLogCode() *string {
	return s.SourceLogCode
}

func (s *AddUserSourceLogConfigRequest) GetSourceLogInfo() *string {
	return s.SourceLogInfo
}

func (s *AddUserSourceLogConfigRequest) GetSourceProdCode() *string {
	return s.SourceProdCode
}

func (s *AddUserSourceLogConfigRequest) GetSubUserId() *int64 {
	return s.SubUserId
}

func (s *AddUserSourceLogConfigRequest) SetDeleted(v int32) *AddUserSourceLogConfigRequest {
	s.Deleted = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetDisPlayLine(v string) *AddUserSourceLogConfigRequest {
	s.DisPlayLine = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetRegionId(v string) *AddUserSourceLogConfigRequest {
	s.RegionId = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceLogCode(v string) *AddUserSourceLogConfigRequest {
	s.SourceLogCode = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceLogInfo(v string) *AddUserSourceLogConfigRequest {
	s.SourceLogInfo = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSourceProdCode(v string) *AddUserSourceLogConfigRequest {
	s.SourceProdCode = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) SetSubUserId(v int64) *AddUserSourceLogConfigRequest {
	s.SubUserId = &v
	return s
}

func (s *AddUserSourceLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
