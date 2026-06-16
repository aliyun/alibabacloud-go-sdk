// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCodeConfiguration interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *CodeConfiguration
	GetChecksum() *string
	SetCommand(v []*string) *CodeConfiguration
	GetCommand() []*string
	SetLanguage(v string) *CodeConfiguration
	GetLanguage() *string
	SetOssBucketName(v string) *CodeConfiguration
	GetOssBucketName() *string
	SetOssObjectName(v string) *CodeConfiguration
	GetOssObjectName() *string
	SetZipFile(v string) *CodeConfiguration
	GetZipFile() *string
}

type CodeConfiguration struct {
	// The CRC-64 checksum of the code package. If you provide `checksum`, Function Compute verifies that the code package\\"s computed checksum matches this value.
	//
	// example:
	//
	// 1234567890123456789
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// The command and arguments to run in the runtime.
	//
	// example:
	//
	// python,main.py
	Command []*string `json:"command" xml:"command" type:"Repeated"`
	// The programming language for the function\\"s runtime, such as python3 or nodejs.
	//
	// example:
	//
	// python3.12
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// The name of the OSS bucket that contains the function\\"s code package.
	//
	// example:
	//
	// my-agent-code-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// The name of the OSS object for the function\\"s code package.
	//
	// example:
	//
	// agent-code-v1.0.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// The base64-encoded content of the agent\\"s code package.
	//
	// example:
	//
	// UEsDBAoAAAAAANF
	ZipFile *string `json:"zipFile,omitempty" xml:"zipFile,omitempty"`
}

func (s CodeConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CodeConfiguration) GoString() string {
	return s.String()
}

func (s *CodeConfiguration) GetChecksum() *string {
	return s.Checksum
}

func (s *CodeConfiguration) GetCommand() []*string {
	return s.Command
}

func (s *CodeConfiguration) GetLanguage() *string {
	return s.Language
}

func (s *CodeConfiguration) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *CodeConfiguration) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *CodeConfiguration) GetZipFile() *string {
	return s.ZipFile
}

func (s *CodeConfiguration) SetChecksum(v string) *CodeConfiguration {
	s.Checksum = &v
	return s
}

func (s *CodeConfiguration) SetCommand(v []*string) *CodeConfiguration {
	s.Command = v
	return s
}

func (s *CodeConfiguration) SetLanguage(v string) *CodeConfiguration {
	s.Language = &v
	return s
}

func (s *CodeConfiguration) SetOssBucketName(v string) *CodeConfiguration {
	s.OssBucketName = &v
	return s
}

func (s *CodeConfiguration) SetOssObjectName(v string) *CodeConfiguration {
	s.OssObjectName = &v
	return s
}

func (s *CodeConfiguration) SetZipFile(v string) *CodeConfiguration {
	s.ZipFile = &v
	return s
}

func (s *CodeConfiguration) Validate() error {
	return dara.Validate(s)
}
