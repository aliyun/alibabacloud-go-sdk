// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealTimeConcurrencyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetRealTimeConcurrencyRequest
	GetInstanceId() *string
}

type GetRealTimeConcurrencyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c112c168ed664c0a851f9ca72d2f7999
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetRealTimeConcurrencyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRealTimeConcurrencyRequest) GoString() string {
	return s.String()
}

func (s *GetRealTimeConcurrencyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetRealTimeConcurrencyRequest) SetInstanceId(v string) *GetRealTimeConcurrencyRequest {
	s.InstanceId = &v
	return s
}

func (s *GetRealTimeConcurrencyRequest) Validate() error {
	return dara.Validate(s)
}
