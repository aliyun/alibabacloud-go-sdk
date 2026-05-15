// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineMessageLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcid(v string) *GetHotlineMessageLogRequest
	GetAcid() *string
	SetInstanceId(v string) *GetHotlineMessageLogRequest
	GetInstanceId() *string
}

type GetHotlineMessageLogRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 100****2077
	Acid *string `json:"Acid,omitempty" xml:"Acid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetHotlineMessageLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineMessageLogRequest) GoString() string {
	return s.String()
}

func (s *GetHotlineMessageLogRequest) GetAcid() *string {
	return s.Acid
}

func (s *GetHotlineMessageLogRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetHotlineMessageLogRequest) SetAcid(v string) *GetHotlineMessageLogRequest {
	s.Acid = &v
	return s
}

func (s *GetHotlineMessageLogRequest) SetInstanceId(v string) *GetHotlineMessageLogRequest {
	s.InstanceId = &v
	return s
}

func (s *GetHotlineMessageLogRequest) Validate() error {
	return dara.Validate(s)
}
