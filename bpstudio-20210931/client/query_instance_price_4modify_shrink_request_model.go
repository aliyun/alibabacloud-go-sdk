// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancePrice4ModifyShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *QueryInstancePrice4ModifyShrinkRequest
	GetApplicationId() *string
	SetConfigurationShrink(v string) *QueryInstancePrice4ModifyShrinkRequest
	GetConfigurationShrink() *string
	SetInstanceId(v string) *QueryInstancePrice4ModifyShrinkRequest
	GetInstanceId() *string
}

type QueryInstancePrice4ModifyShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId       *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	ConfigurationShrink *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstancePrice4ModifyShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancePrice4ModifyShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryInstancePrice4ModifyShrinkRequest) GetConfigurationShrink() *string {
	return s.ConfigurationShrink
}

func (s *QueryInstancePrice4ModifyShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetApplicationId(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetConfigurationShrink(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.ConfigurationShrink = &v
	return s
}

func (s *QueryInstancePrice4ModifyShrinkRequest) SetInstanceId(v string) *QueryInstancePrice4ModifyShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstancePrice4ModifyShrinkRequest) Validate() error {
	return dara.Validate(s)
}
