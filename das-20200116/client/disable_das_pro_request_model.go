// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDasProRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableDasProRequest
	GetInstanceId() *string
	SetUserId(v string) *DisableDasProRequest
	GetUserId() *string
}

type DisableDasProRequest struct {
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
	// >  This parameter is optional. The system can automatically obtain the account ID based on the value of InstanceId that you set when you call this operation.
	//
	// example:
	//
	// 196278346919****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DisableDasProRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableDasProRequest) GoString() string {
	return s.String()
}

func (s *DisableDasProRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableDasProRequest) GetUserId() *string {
	return s.UserId
}

func (s *DisableDasProRequest) SetInstanceId(v string) *DisableDasProRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableDasProRequest) SetUserId(v string) *DisableDasProRequest {
	s.UserId = &v
	return s
}

func (s *DisableDasProRequest) Validate() error {
	return dara.Validate(s)
}
