// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateModuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ValidateModuleRequest
	GetClientToken() *string
	SetCode(v string) *ValidateModuleRequest
	GetCode() *string
	SetCodeMap(v map[string]interface{}) *ValidateModuleRequest
	GetCodeMap() map[string]interface{}
	SetSource(v string) *ValidateModuleRequest
	GetSource() *string
	SetSourcePath(v string) *ValidateModuleRequest
	GetSourcePath() *string
}

type ValidateModuleRequest struct {
	// The idempotency token. Format: [0-9a-zA-Z-]{1,64}. We recommend that you use a UUID.
	//
	// example:
	//
	// 2daf4227f747cbf11a5501f18cc5e004
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	// The template text that is directly passed in when source is set to Upload.
	//
	// example:
	//
	// code
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Specifies multiple configuration files to pass in.
	//
	// example:
	//
	// {"main.tf":"xxx"}
	CodeMap map[string]interface{} `json:"codeMap,omitempty" xml:"codeMap,omitempty"`
	// The template source. Valid values:
	//
	// Upload: text upload.
	//
	// example:
	//
	// Upload
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The source path.
	//
	// example:
	//
	// test
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
}

func (s ValidateModuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateModuleRequest) GoString() string {
	return s.String()
}

func (s *ValidateModuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ValidateModuleRequest) GetCode() *string {
	return s.Code
}

func (s *ValidateModuleRequest) GetCodeMap() map[string]interface{} {
	return s.CodeMap
}

func (s *ValidateModuleRequest) GetSource() *string {
	return s.Source
}

func (s *ValidateModuleRequest) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ValidateModuleRequest) SetClientToken(v string) *ValidateModuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ValidateModuleRequest) SetCode(v string) *ValidateModuleRequest {
	s.Code = &v
	return s
}

func (s *ValidateModuleRequest) SetCodeMap(v map[string]interface{}) *ValidateModuleRequest {
	s.CodeMap = v
	return s
}

func (s *ValidateModuleRequest) SetSource(v string) *ValidateModuleRequest {
	s.Source = &v
	return s
}

func (s *ValidateModuleRequest) SetSourcePath(v string) *ValidateModuleRequest {
	s.SourcePath = &v
	return s
}

func (s *ValidateModuleRequest) Validate() error {
	return dara.Validate(s)
}
