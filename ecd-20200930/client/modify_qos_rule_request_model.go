// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDownload(v int32) *ModifyQosRuleRequest
	GetDownload() *int32
	SetQosRuleId(v string) *ModifyQosRuleRequest
	GetQosRuleId() *string
	SetQosRuleName(v string) *ModifyQosRuleRequest
	GetQosRuleName() *string
	SetUpload(v int32) *ModifyQosRuleRequest
	GetUpload() *int32
}

type ModifyQosRuleRequest struct {
	// example:
	//
	// 10
	Download *int32 `json:"Download,omitempty" xml:"Download,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// qos-5605u0gelk200****
	QosRuleId *string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty"`
	// example:
	//
	// test
	QosRuleName *string `json:"QosRuleName,omitempty" xml:"QosRuleName,omitempty"`
	// example:
	//
	// 10
	Upload *int32 `json:"Upload,omitempty" xml:"Upload,omitempty"`
}

func (s ModifyQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyQosRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifyQosRuleRequest) GetDownload() *int32 {
	return s.Download
}

func (s *ModifyQosRuleRequest) GetQosRuleId() *string {
	return s.QosRuleId
}

func (s *ModifyQosRuleRequest) GetQosRuleName() *string {
	return s.QosRuleName
}

func (s *ModifyQosRuleRequest) GetUpload() *int32 {
	return s.Upload
}

func (s *ModifyQosRuleRequest) SetDownload(v int32) *ModifyQosRuleRequest {
	s.Download = &v
	return s
}

func (s *ModifyQosRuleRequest) SetQosRuleId(v string) *ModifyQosRuleRequest {
	s.QosRuleId = &v
	return s
}

func (s *ModifyQosRuleRequest) SetQosRuleName(v string) *ModifyQosRuleRequest {
	s.QosRuleName = &v
	return s
}

func (s *ModifyQosRuleRequest) SetUpload(v int32) *ModifyQosRuleRequest {
	s.Upload = &v
	return s
}

func (s *ModifyQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
