// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppSecuritiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v int64) *DescribeAppSecuritiesRequest
	GetAppId() *int64
	SetSecurityToken(v string) *DescribeAppSecuritiesRequest
	GetSecurityToken() *string
}

type DescribeAppSecuritiesRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 110862931
	AppId         *int64  `json:"AppId,omitempty" xml:"AppId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeAppSecuritiesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppSecuritiesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppSecuritiesRequest) GetAppId() *int64 {
	return s.AppId
}

func (s *DescribeAppSecuritiesRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeAppSecuritiesRequest) SetAppId(v int64) *DescribeAppSecuritiesRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppSecuritiesRequest) SetSecurityToken(v string) *DescribeAppSecuritiesRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeAppSecuritiesRequest) Validate() error {
	return dara.Validate(s)
}
