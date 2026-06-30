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
	// The business space ID.
	//
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A JSON string that contains the complaint details. For more information, see the following table.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"taskId":"ADCA1DE6-8117-472A-B3A1-352A248F90D0","fileId":"653e563d-774f-4f01-a809-cb8bb920c3e6","rid":1346,"comments":"请重新判定"}
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
