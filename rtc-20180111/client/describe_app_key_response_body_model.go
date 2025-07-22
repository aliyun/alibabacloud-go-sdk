// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v string) *DescribeAppKeyResponseBody
	GetAppKey() *string
	SetRequestId(v string) *DescribeAppKeyResponseBody
	GetRequestId() *string
}

type DescribeAppKeyResponseBody struct {
	// AppKey。
	//
	// example:
	//
	// ba133b2cee4ab9be424674892c33****
	AppKey *string `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 154EF5DE-3D08-1F2C-A482-281F78D74B7C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAppKeyResponseBody) GetAppKey() *string {
	return s.AppKey
}

func (s *DescribeAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAppKeyResponseBody) SetAppKey(v string) *DescribeAppKeyResponseBody {
	s.AppKey = &v
	return s
}

func (s *DescribeAppKeyResponseBody) SetRequestId(v string) *DescribeAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
