// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyParametersRequest
	GetDBInstanceId() *string
	SetForceRestartInstance(v bool) *ModifyParametersRequest
	GetForceRestartInstance() *bool
	SetParameters(v string) *ModifyParametersRequest
	GetParameters() *string
}

type ModifyParametersRequest struct {
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in a specific region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp***************
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies whether to forcibly restart the instance. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	ForceRestartInstance *bool `json:"ForceRestartInstance,omitempty" xml:"ForceRestartInstance,omitempty"`
	// The name and value of the parameter to be modified. Specify the parameter in the `<Parameter name>:<Parameter value>` format.
	//
	// You can call the [DescribeParameters](https://help.aliyun.com/document_detail/208310.html) operation to query the parameters that can be modified.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"statement_timeout":"11800010"}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
}

func (s ModifyParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParametersRequest) GoString() string {
	return s.String()
}

func (s *ModifyParametersRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyParametersRequest) GetForceRestartInstance() *bool {
	return s.ForceRestartInstance
}

func (s *ModifyParametersRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParametersRequest) SetDBInstanceId(v string) *ModifyParametersRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParametersRequest) SetForceRestartInstance(v bool) *ModifyParametersRequest {
	s.ForceRestartInstance = &v
	return s
}

func (s *ModifyParametersRequest) SetParameters(v string) *ModifyParametersRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParametersRequest) Validate() error {
	return dara.Validate(s)
}
