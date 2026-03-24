// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseSceneConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigValue(v string) *DescribeDefenseSceneConfigResponseBody
	GetConfigValue() *string
	SetRequestId(v string) *DescribeDefenseSceneConfigResponseBody
	GetRequestId() *string
}

type DescribeDefenseSceneConfigResponseBody struct {
	// The value of the configuration item. For more information, see the **ConfigValue*	- parameter in [ModifyDefenseSceneConfig](https://help.aliyun.com/document_detail/2968435.html).
	//
	// example:
	//
	// true
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D7861F61-5B61-46CE-A47C-6B19160D5EB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDefenseSceneConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseSceneConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseSceneConfigResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeDefenseSceneConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseSceneConfigResponseBody) SetConfigValue(v string) *DescribeDefenseSceneConfigResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *DescribeDefenseSceneConfigResponseBody) SetRequestId(v string) *DescribeDefenseSceneConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseSceneConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
