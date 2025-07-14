// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeApplicationConfigRequest
	GetAppId() *string
	SetVersionId(v string) *DescribeApplicationConfigRequest
	GetVersionId() *string
}

type DescribeApplicationConfigRequest struct {
	// 7171a6ca-d1cd-4928-8642-7d5cfe69\\*\\*\\*\\*
	//
	// This parameter is required.
	//
	// example:
	//
	// 7171a6ca-d1cd-4928-8642-7d5cfe69****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// 0026ff7f-2b57-4127-bdd0-9bf202bb\\*\\*\\*\\*
	//
	// example:
	//
	// 0026ff7f-2b57-4127-bdd0-9bf202bb****
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DescribeApplicationConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeApplicationConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeApplicationConfigRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *DescribeApplicationConfigRequest) SetAppId(v string) *DescribeApplicationConfigRequest {
	s.AppId = &v
	return s
}

func (s *DescribeApplicationConfigRequest) SetVersionId(v string) *DescribeApplicationConfigRequest {
	s.VersionId = &v
	return s
}

func (s *DescribeApplicationConfigRequest) Validate() error {
	return dara.Validate(s)
}
