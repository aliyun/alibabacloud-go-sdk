// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAtiCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentCertificateId(v string) *DescribeAtiCertificateRequest
	GetAgentCertificateId() *string
	SetClientToken(v string) *DescribeAtiCertificateRequest
	GetClientToken() *string
}

type DescribeAtiCertificateRequest struct {
	// The ID of the certificate to query. Call the ListAtiCertificates operation to query the target certificate information and obtain the ID from the response.
	//
	// example:
	//
	// 2074041094504542210
	AgentCertificateId *string `json:"AgentCertificateId,omitempty" xml:"AgentCertificateId,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure that the value is unique among different requests. ClientToken supports only ASCII characters.
	//
	// > If you do not specify this parameter, the system uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request is different.
	//
	// example:
	//
	// eyJhbGciOiJIUzI1NiIsInR5cC.....
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeAtiCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAtiCertificateRequest) GoString() string {
	return s.String()
}

func (s *DescribeAtiCertificateRequest) GetAgentCertificateId() *string {
	return s.AgentCertificateId
}

func (s *DescribeAtiCertificateRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAtiCertificateRequest) SetAgentCertificateId(v string) *DescribeAtiCertificateRequest {
	s.AgentCertificateId = &v
	return s
}

func (s *DescribeAtiCertificateRequest) SetClientToken(v string) *DescribeAtiCertificateRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAtiCertificateRequest) Validate() error {
	return dara.Validate(s)
}
