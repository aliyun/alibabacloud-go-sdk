// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOutputFuncCode interface {
	dara.Model
	String() string
	GoString() string
	SetChecksum(v string) *OutputFuncCode
	GetChecksum() *string
	SetUrl(v string) *OutputFuncCode
	GetUrl() *string
}

type OutputFuncCode struct {
	// The CRC-64 value of the function code package.
	//
	// example:
	//
	// 1234567890
	Checksum *string `json:"checksum,omitempty" xml:"checksum,omitempty"`
	// The URL of the function code package.
	//
	// example:
	//
	// http://func-code.oss-cn-shanghai.aliyuncs.com/1a2b3c4d5e6f
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s OutputFuncCode) String() string {
	return dara.Prettify(s)
}

func (s OutputFuncCode) GoString() string {
	return s.String()
}

func (s *OutputFuncCode) GetChecksum() *string {
	return s.Checksum
}

func (s *OutputFuncCode) GetUrl() *string {
	return s.Url
}

func (s *OutputFuncCode) SetChecksum(v string) *OutputFuncCode {
	s.Checksum = &v
	return s
}

func (s *OutputFuncCode) SetUrl(v string) *OutputFuncCode {
	s.Url = &v
	return s
}

func (s *OutputFuncCode) Validate() error {
	return dara.Validate(s)
}
