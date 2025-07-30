// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarningConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListWarningConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListWarningConfigRequest
	GetJsonStr() *string
}

type ListWarningConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"pageNumber":1,"pageSize":10}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s ListWarningConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListWarningConfigRequest) GoString() string {
	return s.String()
}

func (s *ListWarningConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListWarningConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListWarningConfigRequest) SetBaseMeAgentId(v int64) *ListWarningConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListWarningConfigRequest) SetJsonStr(v string) *ListWarningConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *ListWarningConfigRequest) Validate() error {
	return dara.Validate(s)
}
