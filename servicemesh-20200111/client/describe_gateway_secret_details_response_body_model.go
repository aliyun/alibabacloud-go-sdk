// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewaySecretDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGatewaySecretDetails(v []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) *DescribeGatewaySecretDetailsResponseBody
	GetGatewaySecretDetails() []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails
	SetRequestId(v string) *DescribeGatewaySecretDetailsResponseBody
	GetRequestId() *string
}

type DescribeGatewaySecretDetailsResponseBody struct {
	// The detailed information about the secret of the ASM gateway.
	GatewaySecretDetails []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails `json:"GatewaySecretDetails,omitempty" xml:"GatewaySecretDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 31d3a0f0-07ed-4f6e-9004-1804498c****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeGatewaySecretDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponseBody) GetGatewaySecretDetails() []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	return s.GatewaySecretDetails
}

func (s *DescribeGatewaySecretDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGatewaySecretDetailsResponseBody) SetGatewaySecretDetails(v []*DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) *DescribeGatewaySecretDetailsResponseBody {
	s.GatewaySecretDetails = v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBody) SetRequestId(v string) *DescribeGatewaySecretDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBody) Validate() error {
	if s.GatewaySecretDetails != nil {
		for _, item := range s.GatewaySecretDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails struct {
	// The time when the certificate expires.
	//
	// example:
	//
	// 2023-03-03 07:45
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The name of the gateway.
	//
	// example:
	//
	// bookinfo-gateway
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The time when the certificate was issued.
	//
	// example:
	//
	// 2022-03-03 07:45
	IssueTime *string `json:"IssueTime,omitempty" xml:"IssueTime,omitempty"`
	// 	- An error message is returned if the status of the gateway is abnormal. Examples: `tls.crt not exist`, `tls.key not exist`, and `secret type must be kubernetes.io/tls`.
	//
	// 	- An empty value is returned if the status of the gateway is normal.
	//
	// example:
	//
	// tls.crt not exist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The Server Name Indication (SNI) value.
	//
	// example:
	//
	// demo.com
	SNI *string `json:"SNI,omitempty" xml:"SNI,omitempty"`
	// The name of the secret.
	//
	// example:
	//
	// demo-secret
	SecretName *string `json:"SecretName,omitempty" xml:"SecretName,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- `normal`
	//
	// 	- `abnormal`
	//
	// example:
	//
	// normal
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GoString() string {
	return s.String()
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetGatewayName() *string {
	return s.GatewayName
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetIssueTime() *string {
	return s.IssueTime
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetMessage() *string {
	return s.Message
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetSNI() *string {
	return s.SNI
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetSecretName() *string {
	return s.SecretName
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) GetState() *string {
	return s.State
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetExpiredTime(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetGatewayName(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.GatewayName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetIssueTime(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.IssueTime = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetMessage(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.Message = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetSNI(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.SNI = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetSecretName(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.SecretName = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) SetState(v string) *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails {
	s.State = &v
	return s
}

func (s *DescribeGatewaySecretDetailsResponseBodyGatewaySecretDetails) Validate() error {
	return dara.Validate(s)
}
