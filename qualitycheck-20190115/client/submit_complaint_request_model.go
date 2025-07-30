// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitComplaintRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SubmitComplaintRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SubmitComplaintRequest
	GetJsonStr() *string
}

type SubmitComplaintRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SubmitComplaintRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitComplaintRequest) GoString() string {
	return s.String()
}

func (s *SubmitComplaintRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SubmitComplaintRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SubmitComplaintRequest) SetBaseMeAgentId(v int64) *SubmitComplaintRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SubmitComplaintRequest) SetJsonStr(v string) *SubmitComplaintRequest {
	s.JsonStr = &v
	return s
}

func (s *SubmitComplaintRequest) Validate() error {
	return dara.Validate(s)
}
