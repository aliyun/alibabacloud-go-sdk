// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppAgentTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteAppAgentTemplateRequest
	GetAppId() *string
	SetId(v string) *DeleteAppAgentTemplateRequest
	GetId() *string
}

type DeleteAppAgentTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1213123142124124124214
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAppAgentTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppAgentTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppAgentTemplateRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteAppAgentTemplateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteAppAgentTemplateRequest) SetAppId(v string) *DeleteAppAgentTemplateRequest {
	s.AppId = &v
	return s
}

func (s *DeleteAppAgentTemplateRequest) SetId(v string) *DeleteAppAgentTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteAppAgentTemplateRequest) Validate() error {
	return dara.Validate(s)
}
