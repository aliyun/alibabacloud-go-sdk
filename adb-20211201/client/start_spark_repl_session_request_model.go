// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSparkReplSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *StartSparkReplSessionRequest
	GetConfig() *string
	SetDBClusterId(v string) *StartSparkReplSessionRequest
	GetDBClusterId() *string
	SetResourceGroupName(v string) *StartSparkReplSessionRequest
	GetResourceGroupName() *string
}

type StartSparkReplSessionRequest struct {
	// The configuration parameters that are used to start the Spark session, which are in the JSON format. For more information, see [Spark application configuration parameters](https://help.aliyun.com/document_detail/471203.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// { "spark.shuffle.timeout": ":0s" }
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1mfe9qm****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the job resource group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s StartSparkReplSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSparkReplSessionRequest) GoString() string {
	return s.String()
}

func (s *StartSparkReplSessionRequest) GetConfig() *string {
	return s.Config
}

func (s *StartSparkReplSessionRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *StartSparkReplSessionRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *StartSparkReplSessionRequest) SetConfig(v string) *StartSparkReplSessionRequest {
	s.Config = &v
	return s
}

func (s *StartSparkReplSessionRequest) SetDBClusterId(v string) *StartSparkReplSessionRequest {
	s.DBClusterId = &v
	return s
}

func (s *StartSparkReplSessionRequest) SetResourceGroupName(v string) *StartSparkReplSessionRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *StartSparkReplSessionRequest) Validate() error {
	return dara.Validate(s)
}
