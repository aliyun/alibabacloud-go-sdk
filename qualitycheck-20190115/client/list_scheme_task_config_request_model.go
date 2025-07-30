// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *ListSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *ListSchemeTaskConfigRequest
	GetJsonStr() *string
}

type ListSchemeTaskConfigRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"pageNumber":1,"pageSize":10,"sourceDataType":"1"}
	JsonStr *string `json:"jsonStr,omitempty" xml:"jsonStr,omitempty"`
}

func (s ListSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *ListSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *ListSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *ListSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *ListSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *ListSchemeTaskConfigRequest) SetJsonStr(v string) *ListSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *ListSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
