// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecordDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *GetRecordDataRequest
	GetAcid() *string
	SetInstanceId(v string) *GetRecordDataRequest
	GetInstanceId() *string
}

type GetRecordDataRequest struct {
	// Session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1001067****
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// Instance ID.
	//
	// Log on to the [Artificial Intelligence Cloud Call Service console](https://aiccs.console.aliyun.com/overview) and view the instance ID in **Instance Management**.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRecordDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecordDataRequest) GoString() string {
	return s.String()
}

func (s *GetRecordDataRequest) GetAcid() *string {
	return s.Acid
}

func (s *GetRecordDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRecordDataRequest) SetAcid(v string) *GetRecordDataRequest {
	s.Acid = &v
	return s
}

func (s *GetRecordDataRequest) SetInstanceId(v string) *GetRecordDataRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRecordDataRequest) Validate() error {
	return dara.Validate(s)
}
