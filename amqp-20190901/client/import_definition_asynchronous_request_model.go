// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDefinitionAsynchronousRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConsoleSessionId(v string) *ImportDefinitionAsynchronousRequest
	GetConsoleSessionId() *string
	SetImportType(v int32) *ImportDefinitionAsynchronousRequest
	GetImportType() *int32
	SetInstanceId(v string) *ImportDefinitionAsynchronousRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ImportDefinitionAsynchronousRequest
	GetInstanceName() *string
	SetOssUrl(v string) *ImportDefinitionAsynchronousRequest
	GetOssUrl() *string
	SetVhostName(v string) *ImportDefinitionAsynchronousRequest
	GetVhostName() *string
}

type ImportDefinitionAsynchronousRequest struct {
	ConsoleSessionId *string `json:"ConsoleSessionId,omitempty" xml:"ConsoleSessionId,omitempty"`
	// This parameter is required.
	ImportType *int32 `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	OssUrl    *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
	VhostName *string `json:"VhostName,omitempty" xml:"VhostName,omitempty"`
}

func (s ImportDefinitionAsynchronousRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDefinitionAsynchronousRequest) GoString() string {
	return s.String()
}

func (s *ImportDefinitionAsynchronousRequest) GetConsoleSessionId() *string {
	return s.ConsoleSessionId
}

func (s *ImportDefinitionAsynchronousRequest) GetImportType() *int32 {
	return s.ImportType
}

func (s *ImportDefinitionAsynchronousRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportDefinitionAsynchronousRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ImportDefinitionAsynchronousRequest) GetOssUrl() *string {
	return s.OssUrl
}

func (s *ImportDefinitionAsynchronousRequest) GetVhostName() *string {
	return s.VhostName
}

func (s *ImportDefinitionAsynchronousRequest) SetConsoleSessionId(v string) *ImportDefinitionAsynchronousRequest {
	s.ConsoleSessionId = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) SetImportType(v int32) *ImportDefinitionAsynchronousRequest {
	s.ImportType = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) SetInstanceId(v string) *ImportDefinitionAsynchronousRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) SetInstanceName(v string) *ImportDefinitionAsynchronousRequest {
	s.InstanceName = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) SetOssUrl(v string) *ImportDefinitionAsynchronousRequest {
	s.OssUrl = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) SetVhostName(v string) *ImportDefinitionAsynchronousRequest {
	s.VhostName = &v
	return s
}

func (s *ImportDefinitionAsynchronousRequest) Validate() error {
	return dara.Validate(s)
}
