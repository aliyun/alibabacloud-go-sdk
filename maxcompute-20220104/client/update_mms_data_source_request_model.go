// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsDataSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateMmsDataSourceRequest
	GetAction() *string
	SetConfig(v map[string]interface{}) *UpdateMmsDataSourceRequest
	GetConfig() map[string]interface{}
	SetName(v string) *UpdateMmsDataSourceRequest
	GetName() *string
	SetTest(v bool) *UpdateMmsDataSourceRequest
	GetTest() *bool
}

type UpdateMmsDataSourceRequest struct {
	Action *string                `json:"action,omitempty" xml:"action,omitempty"`
	Config map[string]interface{} `json:"config,omitempty" xml:"config,omitempty"`
	Name   *string                `json:"name,omitempty" xml:"name,omitempty"`
	Test   *bool                  `json:"test,omitempty" xml:"test,omitempty"`
}

func (s UpdateMmsDataSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsDataSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateMmsDataSourceRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateMmsDataSourceRequest) GetConfig() map[string]interface{} {
	return s.Config
}

func (s *UpdateMmsDataSourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateMmsDataSourceRequest) GetTest() *bool {
	return s.Test
}

func (s *UpdateMmsDataSourceRequest) SetAction(v string) *UpdateMmsDataSourceRequest {
	s.Action = &v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetConfig(v map[string]interface{}) *UpdateMmsDataSourceRequest {
	s.Config = v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetName(v string) *UpdateMmsDataSourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateMmsDataSourceRequest) SetTest(v bool) *UpdateMmsDataSourceRequest {
	s.Test = &v
	return s
}

func (s *UpdateMmsDataSourceRequest) Validate() error {
	return dara.Validate(s)
}
