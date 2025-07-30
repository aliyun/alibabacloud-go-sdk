// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerBySessionGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AssignReviewerBySessionGroupRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *AssignReviewerBySessionGroupRequest
	GetJsonStr() *string
}

type AssignReviewerBySessionGroupRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"sessionGroupParam":{"isSchemeData":1,"callStartTime":"2022-09-17 00:00:00","callEndTime":"2022-09-23 23:59:59","schemeTaskConfigId":24},"assignments":[{"reviewer":63,"count":4}],"isSchemeData":1}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s AssignReviewerBySessionGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerBySessionGroupRequest) GoString() string {
	return s.String()
}

func (s *AssignReviewerBySessionGroupRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AssignReviewerBySessionGroupRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *AssignReviewerBySessionGroupRequest) SetBaseMeAgentId(v int64) *AssignReviewerBySessionGroupRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AssignReviewerBySessionGroupRequest) SetJsonStr(v string) *AssignReviewerBySessionGroupRequest {
	s.JsonStr = &v
	return s
}

func (s *AssignReviewerBySessionGroupRequest) Validate() error {
	return dara.Validate(s)
}
