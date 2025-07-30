// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *AssignReviewerRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *AssignReviewerRequest
	GetJsonStr() *string
}

type AssignReviewerRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"assignmentList":[{"taskId":"1C21CF1E-2917-4236-A046-20E37B293B69","fileId":"7981b466fd6a4c70a7065da159739a5b"},{"taskId":"0111DDBC-5F10-47E0-B7D4-7175F47D626F","fileId":"1814eeae3cff41e989e31ab547f07561"}],"assignments":[{"reviewer":"255746168704895558"},{"reviewer":"268370362815185444"}]}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s AssignReviewerRequest) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerRequest) GoString() string {
	return s.String()
}

func (s *AssignReviewerRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *AssignReviewerRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *AssignReviewerRequest) SetBaseMeAgentId(v int64) *AssignReviewerRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *AssignReviewerRequest) SetJsonStr(v string) *AssignReviewerRequest {
	s.JsonStr = &v
	return s
}

func (s *AssignReviewerRequest) Validate() error {
	return dara.Validate(s)
}
