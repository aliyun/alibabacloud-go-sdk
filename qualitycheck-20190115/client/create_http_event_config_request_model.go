// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpEventConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *CreateHttpEventConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *CreateHttpEventConfigRequest
	GetJsonStr() *string
}

type CreateHttpEventConfigRequest struct {
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

func (s CreateHttpEventConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpEventConfigRequest) GoString() string {
	return s.String()
}

func (s *CreateHttpEventConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *CreateHttpEventConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *CreateHttpEventConfigRequest) SetBaseMeAgentId(v int64) *CreateHttpEventConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *CreateHttpEventConfigRequest) SetJsonStr(v string) *CreateHttpEventConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *CreateHttpEventConfigRequest) Validate() error {
	return dara.Validate(s)
}
