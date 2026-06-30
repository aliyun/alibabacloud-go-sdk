// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveConfigDataSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBaseMeAgentId(v int64) *SaveConfigDataSetRequest
	GetBaseMeAgentId() *int64
	SetJsonStr(v string) *SaveConfigDataSetRequest
	GetJsonStr() *string
}

type SaveConfigDataSetRequest struct {
	// baseMeAgentId
	BaseMeAgentId *int64 `json:"BaseMeAgentId,omitempty" xml:"BaseMeAgentId,omitempty"`
	// A complete JSON string. For details, see the following information.
	//
	// This parameter is required.
	//
	// example:
	//
	// "{"judgeType":1,"setId":851,"jsonStrForRule":{"conditions":[{"cid":"1","check_range":{},"lambda":"1","operators":[{"oid":1,"type":"HIT_ANY_KEYWORDS","param":{"keywords":["有什么可以帮您","客服中心"],"in_sentence":false}}]}],"rules":[{"externalProperty":0,"lambda":"1","rid":"1"}],"roleJudgeMethod":"keyword"},"channelType":1}"
	JsonStr *string `json:"JsonStr,omitempty" xml:"JsonStr,omitempty"`
}

func (s SaveConfigDataSetRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveConfigDataSetRequest) GoString() string {
	return s.String()
}

func (s *SaveConfigDataSetRequest) GetBaseMeAgentId() *int64 {
	return s.BaseMeAgentId
}

func (s *SaveConfigDataSetRequest) GetJsonStr() *string {
	return s.JsonStr
}

func (s *SaveConfigDataSetRequest) SetBaseMeAgentId(v int64) *SaveConfigDataSetRequest {
	s.BaseMeAgentId = &v
	return s
}

func (s *SaveConfigDataSetRequest) SetJsonStr(v string) *SaveConfigDataSetRequest {
	s.JsonStr = &v
	return s
}

func (s *SaveConfigDataSetRequest) Validate() error {
	return dara.Validate(s)
}
