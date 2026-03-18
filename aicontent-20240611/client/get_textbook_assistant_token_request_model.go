// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextbookAssistantTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeviceId(v string) *GetTextbookAssistantTokenRequest
	GetDeviceId() *string
	SetModel(v string) *GetTextbookAssistantTokenRequest
	GetModel() *string
}

type GetTextbookAssistantTokenRequest struct {
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// 700d4d9411efbe6e0
	DeviceId *string `json:"deviceId,omitempty" xml:"deviceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 25032PS56C
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
}

func (s GetTextbookAssistantTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextbookAssistantTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTextbookAssistantTokenRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *GetTextbookAssistantTokenRequest) GetModel() *string {
	return s.Model
}

func (s *GetTextbookAssistantTokenRequest) SetDeviceId(v string) *GetTextbookAssistantTokenRequest {
	s.DeviceId = &v
	return s
}

func (s *GetTextbookAssistantTokenRequest) SetModel(v string) *GetTextbookAssistantTokenRequest {
	s.Model = &v
	return s
}

func (s *GetTextbookAssistantTokenRequest) Validate() error {
	return dara.Validate(s)
}
