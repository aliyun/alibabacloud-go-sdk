// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTagApplyRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryTagApplyRuleResponseBody
	GetCode() *string
	SetData(v *QueryTagApplyRuleResponseBodyData) *QueryTagApplyRuleResponseBody
	GetData() *QueryTagApplyRuleResponseBodyData
	SetMessage(v string) *QueryTagApplyRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryTagApplyRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryTagApplyRuleResponseBody
	GetSuccess() *bool
}

type QueryTagApplyRuleResponseBody struct {
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *QueryTagApplyRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CC3BB6D2-****-****-9DCE-B38165CE4C47
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryTagApplyRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryTagApplyRuleResponseBody) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryTagApplyRuleResponseBody) GetData() *QueryTagApplyRuleResponseBodyData {
	return s.Data
}

func (s *QueryTagApplyRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryTagApplyRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryTagApplyRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryTagApplyRuleResponseBody) SetCode(v string) *QueryTagApplyRuleResponseBody {
	s.Code = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetData(v *QueryTagApplyRuleResponseBodyData) *QueryTagApplyRuleResponseBody {
	s.Data = v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetMessage(v string) *QueryTagApplyRuleResponseBody {
	s.Message = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetRequestId(v string) *QueryTagApplyRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) SetSuccess(v bool) *QueryTagApplyRuleResponseBody {
	s.Success = &v
	return s
}

func (s *QueryTagApplyRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryTagApplyRuleResponseBodyData struct {
	// The requirements for application materials.
	//
	// example:
	//
	// Requirements for application materials
	ApplyMaterialDesc *string `json:"ApplyMaterialDesc,omitempty" xml:"ApplyMaterialDesc,omitempty"`
	// Indicates whether the application is automatically approved.
	//
	// example:
	//
	// 16
	AutoAudit *int64 `json:"AutoAudit,omitempty" xml:"AutoAudit,omitempty"`
	// The URL for the billing documentation.
	//
	// example:
	//
	// aliyundoc.com
	ChargingStandardLink *string `json:"ChargingStandardLink,omitempty" xml:"ChargingStandardLink,omitempty"`
	// indicates whether encrypted queries are supported.
	//
	// example:
	//
	// 99
	EncryptedQuery *int64 `json:"EncryptedQuery,omitempty" xml:"EncryptedQuery,omitempty"`
	// Indicates whether application materials are required.
	//
	// example:
	//
	// 46
	NeedApplyMaterial *int64 `json:"NeedApplyMaterial,omitempty" xml:"NeedApplyMaterial,omitempty"`
	// The URL for the service agreement.
	//
	// example:
	//
	// example.aliyundoc.com
	SlaLink *string `json:"SlaLink,omitempty" xml:"SlaLink,omitempty"`
}

func (s QueryTagApplyRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryTagApplyRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryTagApplyRuleResponseBodyData) GetApplyMaterialDesc() *string {
	return s.ApplyMaterialDesc
}

func (s *QueryTagApplyRuleResponseBodyData) GetAutoAudit() *int64 {
	return s.AutoAudit
}

func (s *QueryTagApplyRuleResponseBodyData) GetChargingStandardLink() *string {
	return s.ChargingStandardLink
}

func (s *QueryTagApplyRuleResponseBodyData) GetEncryptedQuery() *int64 {
	return s.EncryptedQuery
}

func (s *QueryTagApplyRuleResponseBodyData) GetNeedApplyMaterial() *int64 {
	return s.NeedApplyMaterial
}

func (s *QueryTagApplyRuleResponseBodyData) GetSlaLink() *string {
	return s.SlaLink
}

func (s *QueryTagApplyRuleResponseBodyData) SetApplyMaterialDesc(v string) *QueryTagApplyRuleResponseBodyData {
	s.ApplyMaterialDesc = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetAutoAudit(v int64) *QueryTagApplyRuleResponseBodyData {
	s.AutoAudit = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetChargingStandardLink(v string) *QueryTagApplyRuleResponseBodyData {
	s.ChargingStandardLink = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetEncryptedQuery(v int64) *QueryTagApplyRuleResponseBodyData {
	s.EncryptedQuery = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetNeedApplyMaterial(v int64) *QueryTagApplyRuleResponseBodyData {
	s.NeedApplyMaterial = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) SetSlaLink(v string) *QueryTagApplyRuleResponseBodyData {
	s.SlaLink = &v
	return s
}

func (s *QueryTagApplyRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
