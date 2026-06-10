// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackAppInstancePublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *RollbackAppInstancePublishRequest
	GetBizId() *string
	SetDeployChannel(v string) *RollbackAppInstancePublishRequest
	GetDeployChannel() *string
	SetPublishNumber(v string) *RollbackAppInstancePublishRequest
	GetPublishNumber() *string
	SetQuickRollback(v string) *RollbackAppInstancePublishRequest
	GetQuickRollback() *string
}

type RollbackAppInstancePublishRequest struct {
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Deployment channel
	//
	// example:
	//
	// PC
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// Publish number
	//
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// Quick rollback.
	//
	// example:
	//
	// true
	QuickRollback *string `json:"QuickRollback,omitempty" xml:"QuickRollback,omitempty"`
}

func (s RollbackAppInstancePublishRequest) String() string {
	return dara.Prettify(s)
}

func (s RollbackAppInstancePublishRequest) GoString() string {
	return s.String()
}

func (s *RollbackAppInstancePublishRequest) GetBizId() *string {
	return s.BizId
}

func (s *RollbackAppInstancePublishRequest) GetDeployChannel() *string {
	return s.DeployChannel
}

func (s *RollbackAppInstancePublishRequest) GetPublishNumber() *string {
	return s.PublishNumber
}

func (s *RollbackAppInstancePublishRequest) GetQuickRollback() *string {
	return s.QuickRollback
}

func (s *RollbackAppInstancePublishRequest) SetBizId(v string) *RollbackAppInstancePublishRequest {
	s.BizId = &v
	return s
}

func (s *RollbackAppInstancePublishRequest) SetDeployChannel(v string) *RollbackAppInstancePublishRequest {
	s.DeployChannel = &v
	return s
}

func (s *RollbackAppInstancePublishRequest) SetPublishNumber(v string) *RollbackAppInstancePublishRequest {
	s.PublishNumber = &v
	return s
}

func (s *RollbackAppInstancePublishRequest) SetQuickRollback(v string) *RollbackAppInstancePublishRequest {
	s.QuickRollback = &v
	return s
}

func (s *RollbackAppInstancePublishRequest) Validate() error {
	return dara.Validate(s)
}
