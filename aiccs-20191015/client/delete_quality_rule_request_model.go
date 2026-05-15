// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteQualityRuleRequest
	GetId() *int64
	SetInstanceId(v string) *DeleteQualityRuleRequest
	GetInstanceId() *string
}

type DeleteQualityRuleRequest struct {
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteQualityRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityRuleRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteQualityRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteQualityRuleRequest) SetId(v int64) *DeleteQualityRuleRequest {
	s.Id = &v
	return s
}

func (s *DeleteQualityRuleRequest) SetInstanceId(v string) *DeleteQualityRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteQualityRuleRequest) Validate() error {
	return dara.Validate(s)
}
