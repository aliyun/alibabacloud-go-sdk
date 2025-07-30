// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHandleComplaintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *HandleComplaintRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *HandleComplaintRequest
	GetJsonStr() *string
}

type HandleComplaintRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s HandleComplaintRequest) String() string {
	return dara.Prettify(s)
}

func (s HandleComplaintRequest) GoString() string {
	return s.String()
}

func (s *HandleComplaintRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *HandleComplaintRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *HandleComplaintRequest) SetBaseMeAgentId(v int64) *HandleComplaintRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *HandleComplaintRequest) SetJsonStr(v string) *HandleComplaintRequest {
	s.JsonStr = &v
	return s
}

func (s *HandleComplaintRequest) Validate() error {
	return dara.Validate(s)
}
