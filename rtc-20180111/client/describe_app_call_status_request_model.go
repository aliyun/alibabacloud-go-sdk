// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppCallStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DescribeAppCallStatusRequest
	GetAppId() *string
}

type DescribeAppCallStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 223***JQb
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s DescribeAppCallStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppCallStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeAppCallStatusRequest) GetAppId() *string {
	return s.AppId
}

func (s *DescribeAppCallStatusRequest) SetAppId(v string) *DescribeAppCallStatusRequest {
	s.AppId = &v
	return s
}

func (s *DescribeAppCallStatusRequest) Validate() error {
	return dara.Validate(s)
}
