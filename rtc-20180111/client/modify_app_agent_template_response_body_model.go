// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAppAgentTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyAppAgentTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyAppAgentTemplateResponseBody
	GetRequestId() *string
}

type ModifyAppAgentTemplateResponseBody struct {
	// example:
	//
	// 1223131213231313213
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAppAgentTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAppAgentTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAppAgentTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyAppAgentTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAppAgentTemplateResponseBody) SetId(v string) *ModifyAppAgentTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyAppAgentTemplateResponseBody) SetRequestId(v string) *ModifyAppAgentTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAppAgentTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
