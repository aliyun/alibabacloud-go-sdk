// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancePrice4ModifyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *QueryInstancePrice4ModifyRequest
	GetApplicationId() *string
	SetConfiguration(v map[string]interface{}) *QueryInstancePrice4ModifyRequest
	GetConfiguration() map[string]interface{}
	SetInstanceId(v string) *QueryInstancePrice4ModifyRequest
	GetInstanceId() *string
}

type QueryInstancePrice4ModifyRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 002XWH7MXB8MJRU0
	ApplicationId *string                `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	Configuration map[string]interface{} `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rm-uf66k9143r2ch*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s QueryInstancePrice4ModifyRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancePrice4ModifyRequest) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *QueryInstancePrice4ModifyRequest) GetConfiguration() map[string]interface{} {
	return s.Configuration
}

func (s *QueryInstancePrice4ModifyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryInstancePrice4ModifyRequest) SetApplicationId(v string) *QueryInstancePrice4ModifyRequest {
	s.ApplicationId = &v
	return s
}

func (s *QueryInstancePrice4ModifyRequest) SetConfiguration(v map[string]interface{}) *QueryInstancePrice4ModifyRequest {
	s.Configuration = v
	return s
}

func (s *QueryInstancePrice4ModifyRequest) SetInstanceId(v string) *QueryInstancePrice4ModifyRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryInstancePrice4ModifyRequest) Validate() error {
	return dara.Validate(s)
}
