// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmGlobalAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmGlobalAlertRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeCloudGtmGlobalAlertRequest
	GetClientToken() *string
}

type DescribeCloudGtmGlobalAlertRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeCloudGtmGlobalAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmGlobalAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmGlobalAlertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmGlobalAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmGlobalAlertRequest) SetAcceptLanguage(v string) *DescribeCloudGtmGlobalAlertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertRequest) SetClientToken(v string) *DescribeCloudGtmGlobalAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmGlobalAlertRequest) Validate() error {
	return dara.Validate(s)
}
