// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftVersion(v bool) *BeginSessionRequest
	GetDraftVersion() *bool
	SetInstanceId(v string) *BeginSessionRequest
	GetInstanceId() *string
	SetScriptId(v string) *BeginSessionRequest
	GetScriptId() *string
	SetVendorParams(v string) *BeginSessionRequest
	GetVendorParams() *string
}

type BeginSessionRequest struct {
	// example:
	//
	// false
	DraftVersion *bool `json:"DraftVersion,omitempty" xml:"DraftVersion,omitempty"`
	// example:
	//
	// 36e9a4cd-a821-4f78-86fa-a9a4aefeea2e
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 0a88e269-01f5-49ac-97af-5830f0ccb271
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// {"key1":"value1"}
	VendorParams *string `json:"VendorParams,omitempty" xml:"VendorParams,omitempty"`
}

func (s BeginSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s BeginSessionRequest) GoString() string {
	return s.String()
}

func (s *BeginSessionRequest) GetDraftVersion() *bool {
	return s.DraftVersion
}

func (s *BeginSessionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BeginSessionRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *BeginSessionRequest) GetVendorParams() *string {
	return s.VendorParams
}

func (s *BeginSessionRequest) SetDraftVersion(v bool) *BeginSessionRequest {
	s.DraftVersion = &v
	return s
}

func (s *BeginSessionRequest) SetInstanceId(v string) *BeginSessionRequest {
	s.InstanceId = &v
	return s
}

func (s *BeginSessionRequest) SetScriptId(v string) *BeginSessionRequest {
	s.ScriptId = &v
	return s
}

func (s *BeginSessionRequest) SetVendorParams(v string) *BeginSessionRequest {
	s.VendorParams = &v
	return s
}

func (s *BeginSessionRequest) Validate() error {
	return dara.Validate(s)
}
