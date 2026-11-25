// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *DescribeAccessConfigResponseBody
	GetConfigValue() *string
	SetRequestId(v string) *DescribeAccessConfigResponseBody
	GetRequestId() *string
}

type DescribeAccessConfigResponseBody struct {
	// The value of the queried configuration item. For more information, see the **ConfigValue*	- parameter in [ModifyAccessConfig](~~ModifyAccessConfig~~).
	//
	// example:
	//
	// {"status":1,"products":["alb"]}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccessConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessConfigResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeAccessConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessConfigResponseBody) SetConfigValue(v string) *DescribeAccessConfigResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *DescribeAccessConfigResponseBody) SetRequestId(v string) *DescribeAccessConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
