// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHadoopConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *DescribeHadoopConfigsResponseBody
	GetConfigName() *string
	SetConfigValue(v string) *DescribeHadoopConfigsResponseBody
	GetConfigValue() *string
	SetRequestId(v string) *DescribeHadoopConfigsResponseBody
	GetRequestId() *string
}

type DescribeHadoopConfigsResponseBody struct {
	// The name of the configuration file. Valid values:
	//
	// 	- hdfs-site
	//
	// 	- core-site
	//
	// 	- yarn-site
	//
	// 	- mapred-site
	//
	// 	- hive-site
	//
	// example:
	//
	// hdfs-site
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// <?xml version="1.0"?>
	//
	// <configuration>
	//
	//     <property>
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeHadoopConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeHadoopConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeHadoopConfigsResponseBody) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeHadoopConfigsResponseBody) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeHadoopConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeHadoopConfigsResponseBody) SetConfigName(v string) *DescribeHadoopConfigsResponseBody {
	s.ConfigName = &v
	return s
}

func (s *DescribeHadoopConfigsResponseBody) SetConfigValue(v string) *DescribeHadoopConfigsResponseBody {
	s.ConfigValue = &v
	return s
}

func (s *DescribeHadoopConfigsResponseBody) SetRequestId(v string) *DescribeHadoopConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeHadoopConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}
