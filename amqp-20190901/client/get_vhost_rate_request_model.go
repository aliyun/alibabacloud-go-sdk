// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVhostRateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *GetVhostRateRequest
	GetConsoleSessionId() *string
	SetInstanceId(v string) *GetVhostRateRequest
	GetInstanceId() *string
	SetVhostNames(v map[string]interface{}) *GetVhostRateRequest
	GetVhostNames() map[string]interface{}
}

type GetVhostRateRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	VhostNames map[string]interface{} `json:"VhostNames,omitempty" xml:"VhostNames,omitempty"`
}

func (s GetVhostRateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVhostRateRequest) GoString() string {
	return s.String()
}

func (s *GetVhostRateRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *GetVhostRateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetVhostRateRequest) GetVhostNames() map[string]interface{} {
	return s.VhostNames
}

func (s *GetVhostRateRequest) SetConsoleSessionId(v string) *GetVhostRateRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *GetVhostRateRequest) SetInstanceId(v string) *GetVhostRateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetVhostRateRequest) SetVhostNames(v map[string]interface{}) *GetVhostRateRequest {
	s.VhostNames = v
	return s
}

func (s *GetVhostRateRequest) Validate() error {
	return dara.Validate(s)
}
