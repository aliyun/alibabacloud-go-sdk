// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppKeyRequest
	GetAppId() *string
	SetOwnerId(v int64) *DescribeAppKeyRequest
	GetOwnerId() *int64
}

type DescribeAppKeyRequest struct {
	// Application ID. You can create and query this ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0cho****
	AppId   *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s DescribeAppKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppKeyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAppKeyRequest) SetAppId(v string) *DescribeAppKeyRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppKeyRequest) SetOwnerId(v int64) *DescribeAppKeyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAppKeyRequest) Validate() error {
	return dara.Validate(s)
}
