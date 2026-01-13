// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeRecallManagementServiceVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ChangeRecallManagementServiceVersionRequest
	GetInstanceId() *string
	SetRecallManagementServiceVersionId(v string) *ChangeRecallManagementServiceVersionRequest
	GetRecallManagementServiceVersionId() *string
}

type ChangeRecallManagementServiceVersionRequest struct {
	InstanceId                       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RecallManagementServiceVersionId *string `json:"RecallManagementServiceVersionId,omitempty" xml:"RecallManagementServiceVersionId,omitempty"`
}

func (s ChangeRecallManagementServiceVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeRecallManagementServiceVersionRequest) GoString() string {
	return s.String()
}

func (s *ChangeRecallManagementServiceVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ChangeRecallManagementServiceVersionRequest) GetRecallManagementServiceVersionId() *string {
	return s.RecallManagementServiceVersionId
}

func (s *ChangeRecallManagementServiceVersionRequest) SetInstanceId(v string) *ChangeRecallManagementServiceVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *ChangeRecallManagementServiceVersionRequest) SetRecallManagementServiceVersionId(v string) *ChangeRecallManagementServiceVersionRequest {
	s.RecallManagementServiceVersionId = &v
	return s
}

func (s *ChangeRecallManagementServiceVersionRequest) Validate() error {
	return dara.Validate(s)
}
