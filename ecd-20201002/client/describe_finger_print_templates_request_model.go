// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFingerPrintTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientId(v string) *DescribeFingerPrintTemplatesRequest
	GetClientId() *string
	SetLoginToken(v string) *DescribeFingerPrintTemplatesRequest
	GetLoginToken() *string
	SetRegionId(v string) *DescribeFingerPrintTemplatesRequest
	GetRegionId() *string
	SetSessionId(v string) *DescribeFingerPrintTemplatesRequest
	GetSessionId() *string
}

type DescribeFingerPrintTemplatesRequest struct {
	// The client ID. The system generates a unique ID for each client.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61e39dc6-0450-45f6-a372-2a09e938****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// The logon token.
	//
	// This parameter is required.
	//
	// example:
	//
	// v189646d6f329e4dfcbf51653542202890570fec26e4f9ee26427c5920fcd93871f017d2190199c4c7d0c0bf00f573****
	LoginToken *string `json:"LoginToken,omitempty" xml:"LoginToken,omitempty"`
	// The region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a5062d68-e550-4d09-8288-67c8ba9e****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeFingerPrintTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFingerPrintTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFingerPrintTemplatesRequest) GetClientId() *string {
	return s.ClientId
}

func (s *DescribeFingerPrintTemplatesRequest) GetLoginToken() *string {
	return s.LoginToken
}

func (s *DescribeFingerPrintTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFingerPrintTemplatesRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeFingerPrintTemplatesRequest) SetClientId(v string) *DescribeFingerPrintTemplatesRequest {
	s.ClientId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetLoginToken(v string) *DescribeFingerPrintTemplatesRequest {
	s.LoginToken = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetRegionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) SetSessionId(v string) *DescribeFingerPrintTemplatesRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeFingerPrintTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
