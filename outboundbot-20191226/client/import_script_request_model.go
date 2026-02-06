// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ImportScriptRequest
	GetInstanceId() *string
	SetNluEngine(v string) *ImportScriptRequest
	GetNluEngine() *string
	SetSignatureUrl(v string) *ImportScriptRequest
	GetSignatureUrl() *string
}

type ImportScriptRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	NluEngine  *string `json:"NluEngine,omitempty" xml:"NluEngine,omitempty"`
	// This parameter is required.
	SignatureUrl *string `json:"SignatureUrl,omitempty" xml:"SignatureUrl,omitempty"`
}

func (s ImportScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportScriptRequest) GoString() string {
	return s.String()
}

func (s *ImportScriptRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportScriptRequest) GetNluEngine() *string {
	return s.NluEngine
}

func (s *ImportScriptRequest) GetSignatureUrl() *string {
	return s.SignatureUrl
}

func (s *ImportScriptRequest) SetInstanceId(v string) *ImportScriptRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportScriptRequest) SetNluEngine(v string) *ImportScriptRequest {
	s.NluEngine = &v
	return s
}

func (s *ImportScriptRequest) SetSignatureUrl(v string) *ImportScriptRequest {
	s.SignatureUrl = &v
	return s
}

func (s *ImportScriptRequest) Validate() error {
	return dara.Validate(s)
}
