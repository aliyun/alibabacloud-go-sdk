// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppAgentFunctionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppAgentFunctionStatusRequest
	GetAppId() *string
}

type DescribeAppAgentFunctionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppAgentFunctionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppAgentFunctionStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppAgentFunctionStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppAgentFunctionStatusRequest) SetAppId(v string) *DescribeAppAgentFunctionStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppAgentFunctionStatusRequest) Validate() error {
	return dara.Validate(s)
}
