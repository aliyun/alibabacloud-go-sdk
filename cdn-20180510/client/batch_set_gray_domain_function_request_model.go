// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetGrayDomainFunctionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v string) *BatchSetGrayDomainFunctionRequest
	GetConfigs() *string
	SetDomainNames(v string) *BatchSetGrayDomainFunctionRequest
	GetDomainNames() *string
}

type BatchSetGrayDomainFunctionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs": [{"argName": "key","argValue": "Content-Encoding"},{"argName": "value","argValue": "gzip"}],"functionName": "set_resp_header"} ]
	Configs *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com,xxx.org.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
}

func (s BatchSetGrayDomainFunctionRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetGrayDomainFunctionRequest) GoString() string {
	return s.String()
}

func (s *BatchSetGrayDomainFunctionRequest) GetConfigs() *string {
	return s.Configs
}

func (s *BatchSetGrayDomainFunctionRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetGrayDomainFunctionRequest) SetConfigs(v string) *BatchSetGrayDomainFunctionRequest {
	s.Configs = &v
	return s
}

func (s *BatchSetGrayDomainFunctionRequest) SetDomainNames(v string) *BatchSetGrayDomainFunctionRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetGrayDomainFunctionRequest) Validate() error {
	return dara.Validate(s)
}
