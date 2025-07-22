// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDasProServiceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetDasProServiceUsageRequest
	GetInstanceId() *string
	SetUserId(v string) *GetDasProServiceUsageRequest
	GetUserId() *string
}

type GetDasProServiceUsageRequest struct {
	// The database instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that is used to create the database instance.
	//
	// >  This parameter is optional. The system can automatically obtain the account ID based on the value of InstanceId when you call this operation.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDasProServiceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDasProServiceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetDasProServiceUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDasProServiceUsageRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetDasProServiceUsageRequest) SetInstanceId(v string) *GetDasProServiceUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDasProServiceUsageRequest) SetUserId(v string) *GetDasProServiceUsageRequest {
	s.UserId = &v
	return s
}

func (s *GetDasProServiceUsageRequest) Validate() error {
	return dara.Validate(s)
}
