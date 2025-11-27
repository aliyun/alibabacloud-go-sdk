// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitTenantAliasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasInfo(v *InitTenantAliasResponseBodyAliasInfo) *InitTenantAliasResponseBody
	GetAliasInfo() *InitTenantAliasResponseBodyAliasInfo
	SetData(v string) *InitTenantAliasResponseBody
	GetData() *string
	SetRequestId(v string) *InitTenantAliasResponseBody
	GetRequestId() *string
}

type InitTenantAliasResponseBody struct {
	// The data returned.
	AliasInfo *InitTenantAliasResponseBodyAliasInfo `json:"AliasInfo,omitempty" xml:"AliasInfo,omitempty" type:"Struct"`
	// The generated ID of the organization.
	//
	// example:
	//
	// WY23***
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitTenantAliasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitTenantAliasResponseBody) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponseBody) GetAliasInfo() *InitTenantAliasResponseBodyAliasInfo {
	return s.AliasInfo
}

func (s *InitTenantAliasResponseBody) GetData() *string {
	return s.Data
}

func (s *InitTenantAliasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitTenantAliasResponseBody) SetAliasInfo(v *InitTenantAliasResponseBodyAliasInfo) *InitTenantAliasResponseBody {
	s.AliasInfo = v
	return s
}

func (s *InitTenantAliasResponseBody) SetData(v string) *InitTenantAliasResponseBody {
	s.Data = &v
	return s
}

func (s *InitTenantAliasResponseBody) SetRequestId(v string) *InitTenantAliasResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitTenantAliasResponseBody) Validate() error {
	if s.AliasInfo != nil {
		if err := s.AliasInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitTenantAliasResponseBodyAliasInfo struct {
	// The reason why modification is not allowed.
	//
	// example:
	//
	// FrequencyExceedsLimit
	AliasEditDisabledReason *string `json:"AliasEditDisabledReason,omitempty" xml:"AliasEditDisabledReason,omitempty"`
	// Indicates whether modification is allowed.
	//
	// example:
	//
	// False
	AliasEditable *bool `json:"AliasEditable,omitempty" xml:"AliasEditable,omitempty"`
	// The source of the organization ID.
	//
	// Valid values:
	//
	// 	- Generated: auto-generated.
	//
	// 	- Customized: user-defined.
	//
	// example:
	//
	// Customized
	AliasSourceType *string `json:"AliasSourceType,omitempty" xml:"AliasSourceType,omitempty"`
	// The time window during which modification is allowed.
	//
	// example:
	//
	// 2025-04-17 20:31:48
	NextModifyTime *string `json:"NextModifyTime,omitempty" xml:"NextModifyTime,omitempty"`
}

func (s InitTenantAliasResponseBodyAliasInfo) String() string {
	return dara.Prettify(s)
}

func (s InitTenantAliasResponseBodyAliasInfo) GoString() string {
	return s.String()
}

func (s *InitTenantAliasResponseBodyAliasInfo) GetAliasEditDisabledReason() *string {
	return s.AliasEditDisabledReason
}

func (s *InitTenantAliasResponseBodyAliasInfo) GetAliasEditable() *bool {
	return s.AliasEditable
}

func (s *InitTenantAliasResponseBodyAliasInfo) GetAliasSourceType() *string {
	return s.AliasSourceType
}

func (s *InitTenantAliasResponseBodyAliasInfo) GetNextModifyTime() *string {
	return s.NextModifyTime
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasEditDisabledReason(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasEditDisabledReason = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasEditable(v bool) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasEditable = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetAliasSourceType(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.AliasSourceType = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) SetNextModifyTime(v string) *InitTenantAliasResponseBodyAliasInfo {
	s.NextModifyTime = &v
	return s
}

func (s *InitTenantAliasResponseBodyAliasInfo) Validate() error {
	return dara.Validate(s)
}
