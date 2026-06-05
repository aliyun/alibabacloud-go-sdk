// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAppInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *PublishAppInstanceRequest
	GetBizId() *string
	SetDeployChannel(v string) *PublishAppInstanceRequest
	GetDeployChannel() *string
	SetDescription(v string) *PublishAppInstanceRequest
	GetDescription() *string
	SetLogicalNumber(v int32) *PublishAppInstanceRequest
	GetLogicalNumber() *int32
	SetPublishNumber(v string) *PublishAppInstanceRequest
	GetPublishNumber() *string
	SetWeappAction(v string) *PublishAppInstanceRequest
	GetWeappAction() *string
}

type PublishAppInstanceRequest struct {
	// example:
	//
	// WD20250821161210000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// PC
	DeployChannel *string `json:"DeployChannel,omitempty" xml:"DeployChannel,omitempty"`
	// example:
	//
	// CREATE_BY_NLB.nlb-jcvs5sm9l8um84zbfa
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 111
	LogicalNumber *int32 `json:"LogicalNumber,omitempty" xml:"LogicalNumber,omitempty"`
	// example:
	//
	// 123
	PublishNumber *string `json:"PublishNumber,omitempty" xml:"PublishNumber,omitempty"`
	// example:
	//
	// BUILD
	WeappAction *string `json:"WeappAction,omitempty" xml:"WeappAction,omitempty"`
}

func (s PublishAppInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishAppInstanceRequest) GoString() string {
	return s.String()
}

func (s *PublishAppInstanceRequest) GetBizId() *string {
	return s.BizId
}

func (s *PublishAppInstanceRequest) GetDeployChannel() *string {
	return s.DeployChannel
}

func (s *PublishAppInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *PublishAppInstanceRequest) GetLogicalNumber() *int32 {
	return s.LogicalNumber
}

func (s *PublishAppInstanceRequest) GetPublishNumber() *string {
	return s.PublishNumber
}

func (s *PublishAppInstanceRequest) GetWeappAction() *string {
	return s.WeappAction
}

func (s *PublishAppInstanceRequest) SetBizId(v string) *PublishAppInstanceRequest {
	s.BizId = &v
	return s
}

func (s *PublishAppInstanceRequest) SetDeployChannel(v string) *PublishAppInstanceRequest {
	s.DeployChannel = &v
	return s
}

func (s *PublishAppInstanceRequest) SetDescription(v string) *PublishAppInstanceRequest {
	s.Description = &v
	return s
}

func (s *PublishAppInstanceRequest) SetLogicalNumber(v int32) *PublishAppInstanceRequest {
	s.LogicalNumber = &v
	return s
}

func (s *PublishAppInstanceRequest) SetPublishNumber(v string) *PublishAppInstanceRequest {
	s.PublishNumber = &v
	return s
}

func (s *PublishAppInstanceRequest) SetWeappAction(v string) *PublishAppInstanceRequest {
	s.WeappAction = &v
	return s
}

func (s *PublishAppInstanceRequest) Validate() error {
	return dara.Validate(s)
}
