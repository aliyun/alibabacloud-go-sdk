// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppRecordStatusRequest
	GetAppId() *string
	SetClientToken(v string) *DescribeAppRecordStatusRequest
	GetClientToken() *string
}

type DescribeAppRecordStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ac7N****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
}

func (s DescribeAppRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppRecordStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppRecordStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeAppRecordStatusRequest) SetAppId(v string) *DescribeAppRecordStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppRecordStatusRequest) SetClientToken(v string) *DescribeAppRecordStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeAppRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
