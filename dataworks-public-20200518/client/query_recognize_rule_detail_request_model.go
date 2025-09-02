// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecognizeRuleDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSensitiveName(v string) *QueryRecognizeRuleDetailRequest
	GetSensitiveName() *string
	SetTenantId(v string) *QueryRecognizeRuleDetailRequest
	GetTenantId() *string
}

type QueryRecognizeRuleDetailRequest struct {
	// The name of the sensitive field. To obtain the name of the sensitive field, call the [QuerySensNodeInfo](https://help.aliyun.com/document_detail/2747189.html) operation or log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Data Category and Sensitivity Level page.
	//
	// This parameter is required.
	//
	// example:
	//
	// Name
	SensitiveName *string `json:"SensitiveName,omitempty" xml:"SensitiveName,omitempty"`
	// The tenant ID. To obtain the tenant ID, perform the following steps: Log on to the [DataWorks console](https://workbench.data.aliyun.com/console). Find your workspace and go to the DataStudio page. On the DataStudio page, click the logon username in the upper-right corner and click User Info in the Menu section.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10241024
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s QueryRecognizeRuleDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRecognizeRuleDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryRecognizeRuleDetailRequest) GetSensitiveName() *string {
	return s.SensitiveName
}

func (s *QueryRecognizeRuleDetailRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *QueryRecognizeRuleDetailRequest) SetSensitiveName(v string) *QueryRecognizeRuleDetailRequest {
	s.SensitiveName = &v
	return s
}

func (s *QueryRecognizeRuleDetailRequest) SetTenantId(v string) *QueryRecognizeRuleDetailRequest {
	s.TenantId = &v
	return s
}

func (s *QueryRecognizeRuleDetailRequest) Validate() error {
	return dara.Validate(s)
}
