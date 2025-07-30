// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNextResultToVerifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetNextResultToVerifyRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetNextResultToVerifyRequest
	GetJsonStr() *string
}

type GetNextResultToVerifyRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// "{"pageNumber":1,"pageSize":1,"taskId":"593A04C0-E6E9-4CDC-8C9*****","original":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetNextResultToVerifyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNextResultToVerifyRequest) GoString() string {
	return s.String()
}

func (s *GetNextResultToVerifyRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetNextResultToVerifyRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetNextResultToVerifyRequest) SetBaseMeAgentId(v int64) *GetNextResultToVerifyRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetNextResultToVerifyRequest) SetJsonStr(v string) *GetNextResultToVerifyRequest {
	s.JsonStr = &v
	return s
}

func (s *GetNextResultToVerifyRequest) Validate() error {
	return dara.Validate(s)
}
