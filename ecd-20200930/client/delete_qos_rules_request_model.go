// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetQosRuleId(v []*string) *DeleteQosRulesRequest
	GetQosRuleId() []*string
}

type DeleteQosRulesRequest struct {
	// This parameter is required.
	QosRuleId []*string `json:"QosRuleId,omitempty" xml:"QosRuleId,omitempty" type:"Repeated"`
}

func (s DeleteQosRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosRulesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQosRulesRequest) GetQosRuleId() []*string {
	return s.QosRuleId
}

func (s *DeleteQosRulesRequest) SetQosRuleId(v []*string) *DeleteQosRulesRequest {
	s.QosRuleId = v
	return s
}

func (s *DeleteQosRulesRequest) Validate() error {
	return dara.Validate(s)
}
