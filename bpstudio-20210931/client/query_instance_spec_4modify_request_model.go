// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceSpec4ModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *QueryInstanceSpec4ModifyRequest
	GetApplicationId() *string
	SetInstanceId(v string) *QueryInstanceSpec4ModifyRequest
	GetInstanceId() *string
	SetMethodName(v string) *QueryInstanceSpec4ModifyRequest
	GetMethodName() *string
	SetParameters(v map[string]interface{}) *QueryInstanceSpec4ModifyRequest
	GetParameters() map[string]interface{}
}

type QueryInstanceSpec4ModifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// BE68D71ZY5YYIU9R
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string                `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MethodName *string                `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s QueryInstanceSpec4ModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryInstanceSpec4ModifyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceSpec4ModifyRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *QueryInstanceSpec4ModifyRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *QueryInstanceSpec4ModifyRequest) SetApplicationId(v string) *QueryInstanceSpec4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetInstanceId(v string) *QueryInstanceSpec4ModifyRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetMethodName(v string) *QueryInstanceSpec4ModifyRequest {
	s.MethodName = &v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) SetParameters(v map[string]interface{}) *QueryInstanceSpec4ModifyRequest {
	s.Parameters = v
	return s
}

func (s *QueryInstanceSpec4ModifyRequest) Validate() error {
	return dara.Validate(s)
}
