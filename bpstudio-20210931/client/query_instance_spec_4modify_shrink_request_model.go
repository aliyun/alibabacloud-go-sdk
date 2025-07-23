// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceSpec4ModifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *QueryInstanceSpec4ModifyShrinkRequest
	GetApplicationId() *string
	SetInstanceId(v string) *QueryInstanceSpec4ModifyShrinkRequest
	GetInstanceId() *string
	SetMethodName(v string) *QueryInstanceSpec4ModifyShrinkRequest
	GetMethodName() *string
	SetParametersShrink(v string) *QueryInstanceSpec4ModifyShrinkRequest
	GetParametersShrink() *string
}

type QueryInstanceSpec4ModifyShrinkRequest struct {
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
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MethodName       *string `json:"MethodName,omitempty" xml:"MethodName,omitempty"`
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s QueryInstanceSpec4ModifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) GetMethodName() *string {
	return s.MethodName
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetApplicationId(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetInstanceId(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetMethodName(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.MethodName = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) SetParametersShrink(v string) *QueryInstanceSpec4ModifyShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *QueryInstanceSpec4ModifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
