// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRecognizeRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSensitiveId(v string) *DeleteRecognizeRuleRequest
	GetSensitiveId() *string
	SetTenantId(v string) *DeleteRecognizeRuleRequest
	GetTenantId() *string
}

type DeleteRecognizeRuleRequest struct {
	// The ID of the sensitive field. You can call the [QuerySensNodeInfo](https://help.aliyun.com/document_detail/2747189.html) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 27f5f5e2-ec60-4567-b1e4-779ac3681024
	SensitiveId *string `json:"SensitiveId,omitempty" xml:"SensitiveId,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s DeleteRecognizeRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRecognizeRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteRecognizeRuleRequest) GetSensitiveId() *string {
	return s.SensitiveId
}

func (s *DeleteRecognizeRuleRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *DeleteRecognizeRuleRequest) SetSensitiveId(v string) *DeleteRecognizeRuleRequest {
	s.SensitiveId = &v
	return s
}

func (s *DeleteRecognizeRuleRequest) SetTenantId(v string) *DeleteRecognizeRuleRequest {
	s.TenantId = &v
	return s
}

func (s *DeleteRecognizeRuleRequest) Validate() error {
	return dara.Validate(s)
}
