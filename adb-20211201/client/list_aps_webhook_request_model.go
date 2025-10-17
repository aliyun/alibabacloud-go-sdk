// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApsWebhookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *ListApsWebhookRequest
	GetDBClusterId() *string
	SetJobType(v string) *ListApsWebhookRequest
	GetJobType() *string
	SetRegionId(v string) *ListApsWebhookRequest
	GetRegionId() *string
}

type ListApsWebhookRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// amv-8vbuyjhrih**
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// example:
	//
	// ResultExport
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApsWebhookRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApsWebhookRequest) GoString() string {
	return s.String()
}

func (s *ListApsWebhookRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListApsWebhookRequest) GetJobType() *string {
	return s.JobType
}

func (s *ListApsWebhookRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApsWebhookRequest) SetDBClusterId(v string) *ListApsWebhookRequest {
	s.DBClusterId = &v
	return s
}

func (s *ListApsWebhookRequest) SetJobType(v string) *ListApsWebhookRequest {
	s.JobType = &v
	return s
}

func (s *ListApsWebhookRequest) SetRegionId(v string) *ListApsWebhookRequest {
	s.RegionId = &v
	return s
}

func (s *ListApsWebhookRequest) Validate() error {
	return dara.Validate(s)
}
