// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTTSConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTTSConfigRequest
	GetInstanceId() *string
	SetScriptId(v string) *DescribeTTSConfigRequest
	GetScriptId() *string
}

type DescribeTTSConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0ec0c897-b92c-40e4-9ad7-e6e4f5ce13bb
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// d7fbd0a0-27bc-49c4-a456-ecb75e79122b
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s DescribeTTSConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTTSConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeTTSConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTTSConfigRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeTTSConfigRequest) SetInstanceId(v string) *DescribeTTSConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTTSConfigRequest) SetScriptId(v string) *DescribeTTSConfigRequest {
	s.ScriptId = &v
	return s
}

func (s *DescribeTTSConfigRequest) Validate() error {
	return dara.Validate(s)
}
