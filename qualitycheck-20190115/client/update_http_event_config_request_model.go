// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpEventConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *UpdateHttpEventConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *UpdateHttpEventConfigRequest
	GetJsonStr() *string
}

type UpdateHttpEventConfigRequest struct {
	// example:
	//
	// 123456
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {
	//
	//     "switchFlag": true,
	//
	//     "callbackUrl": "http://123213.com",
	//
	//     "httpMethod": "GET",
	//
	//     "eventConfig": {
	//
	//         "taskCompleteEvent": true,
	//
	//         "fileReviewedEvent": true,
	//
	//         "complaintHandledEvent": true
	//
	//     }
	//
	// }
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s UpdateHttpEventConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpEventConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateHttpEventConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *UpdateHttpEventConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *UpdateHttpEventConfigRequest) SetBaseMeAgentId(v int64) *UpdateHttpEventConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *UpdateHttpEventConfigRequest) SetJsonStr(v string) *UpdateHttpEventConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *UpdateHttpEventConfigRequest) Validate() error {
	return dara.Validate(s)
}
