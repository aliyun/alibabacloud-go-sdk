// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteAppAgentTemplateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteAppAgentTemplateResponseBody
	GetRequestId() *string
}

type DeleteAppAgentTemplateResponseBody struct {
	// example:
	//
	// 1223131213231313213
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppAgentTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentTemplateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteAppAgentTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppAgentTemplateResponseBody) SetId(v string) *DeleteAppAgentTemplateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteAppAgentTemplateResponseBody) SetRequestId(v string) *DeleteAppAgentTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppAgentTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
