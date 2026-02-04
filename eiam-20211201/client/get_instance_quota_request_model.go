// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceQuotaRequest
	GetInstanceId() *string
	SetQuotaKey(v string) *GetInstanceQuotaRequest
	GetQuotaKey() *string
}

type GetInstanceQuotaRequest struct {
	// IDaaS EIAM实例的ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// 配额类型，QuotaEnum
	//
	// This parameter is required.
	//
	// example:
	//
	// userMaxNumber
	QuotaKey *string `json:"QuotaKey,omitempty" xml:"QuotaKey,omitempty"`
}

func (s GetInstanceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceQuotaRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceQuotaRequest) GetQuotaKey() *string {
	return s.QuotaKey
}

func (s *GetInstanceQuotaRequest) SetInstanceId(v string) *GetInstanceQuotaRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceQuotaRequest) SetQuotaKey(v string) *GetInstanceQuotaRequest {
	s.QuotaKey = &v
	return s
}

func (s *GetInstanceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
