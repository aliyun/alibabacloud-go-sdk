// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePipelineManagementConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePipelineManagementConfigRequest
	GetClientToken() *string
}

type DescribePipelineManagementConfigRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
}

func (s DescribePipelineManagementConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePipelineManagementConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribePipelineManagementConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePipelineManagementConfigRequest) SetClientToken(v string) *DescribePipelineManagementConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePipelineManagementConfigRequest) Validate() error {
	return dara.Validate(s)
}
