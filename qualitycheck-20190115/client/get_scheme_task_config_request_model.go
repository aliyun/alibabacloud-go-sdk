// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSchemeTaskConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *GetSchemeTaskConfigRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *GetSchemeTaskConfigRequest
	GetJsonStr() *string
}

type GetSchemeTaskConfigRequest struct {
	// example:
	//
	// 12345
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// example:
	//
	// {"sourceDataType":3,"id":588}
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s GetSchemeTaskConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSchemeTaskConfigRequest) GoString() string {
	return s.String()
}

func (s *GetSchemeTaskConfigRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *GetSchemeTaskConfigRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *GetSchemeTaskConfigRequest) SetBaseMeAgentId(v int64) *GetSchemeTaskConfigRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *GetSchemeTaskConfigRequest) SetJsonStr(v string) *GetSchemeTaskConfigRequest {
	s.JsonStr = &v
	return s
}

func (s *GetSchemeTaskConfigRequest) Validate() error {
	return dara.Validate(s)
}
