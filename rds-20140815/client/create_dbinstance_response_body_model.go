// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionString(v string) *CreateDBInstanceResponseBody
	GetConnectionString() *string
	SetDBInstanceId(v string) *CreateDBInstanceResponseBody
	GetDBInstanceId() *string
	SetDryRun(v bool) *CreateDBInstanceResponseBody
	GetDryRun() *bool
	SetDryRunResult(v bool) *CreateDBInstanceResponseBody
	GetDryRunResult() *bool
	SetMessage(v string) *CreateDBInstanceResponseBody
	GetMessage() *string
	SetOrderId(v string) *CreateDBInstanceResponseBody
	GetOrderId() *string
	SetPort(v string) *CreateDBInstanceResponseBody
	GetPort() *string
	SetRequestId(v string) *CreateDBInstanceResponseBody
	GetRequestId() *string
	SetTagResult(v bool) *CreateDBInstanceResponseBody
	GetTagResult() *bool
	SetTaskId(v string) *CreateDBInstanceResponseBody
	GetTaskId() *string
}

type CreateDBInstanceResponseBody struct {
	// The internal endpoint of the instance.
	//
	// example:
	//
	// rm-uf6wjk5*****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID. If the value of the **Amount*	- parameter is greater than **1**, more than one instance ID is returned. The number of instance IDs that are returned is the same as the value of the Amount parameter. The returned instance IDs are separated by commas (,).
	//
	// For example, if the value of the **Amount*	- parameter is **3**, three instance IDs are returned. Examples: `rm-uf6wjk5*****1,rm-uf6wjk5*****2,rm-uf6wjk5*****3`
	//
	// example:
	//
	// rm-uf6wjk5*****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates that the system performed a dry run.
	//
	// 	- The value is fixed as **true**.
	//
	// 	- If the system does not perform a dry run, this parameter is not returned.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Indicates whether the request passed the dry run. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// > 	- If the system does not perform a dry run, this parameter is not returned.
	//
	// > 	- If the request failed the dry run, an error message is returned.
	//
	// example:
	//
	// true
	DryRunResult *bool `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty"`
	// The message that indicates whether multiple instances are created.
	//
	// > The parameter is returned only when the value of the **Amount*	- parameter is greater than 1.
	//
	// example:
	//
	// Batch Create DBInstance Task Is In Process.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 1007893702*****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The internal IP address and port number that are used to connect to the instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the specified tag is added to the instance. Valid values:
	//
	// 	- **true**: The specified tag is added to the instance.
	//
	// 	- **false**: The specified tag fails to be added to the instance.
	//
	// > If you do not add a tag to the instance, this parameter is not returned.
	//
	// example:
	//
	// true
	TagResult *bool `json:"TagResult,omitempty" xml:"TagResult,omitempty"`
	// The ID of the task that is run to create multiple instances.
	//
	// 	- This parameter is returned only when the value of **Amount*	- is greater than 1.
	//
	// 	- The **TaskID*	- parameter cannot be used to query a task.
	//
	// example:
	//
	// s2365879-a9d0-55af-fgae-f2*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateDBInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDBInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDBInstanceResponseBody) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *CreateDBInstanceResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *CreateDBInstanceResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDBInstanceResponseBody) GetDryRunResult() *bool {
	return s.DryRunResult
}

func (s *CreateDBInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDBInstanceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateDBInstanceResponseBody) GetPort() *string {
	return s.Port
}

func (s *CreateDBInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDBInstanceResponseBody) GetTagResult() *bool {
	return s.TagResult
}

func (s *CreateDBInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateDBInstanceResponseBody) SetConnectionString(v string) *CreateDBInstanceResponseBody {
	s.ConnectionString = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDBInstanceId(v string) *CreateDBInstanceResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDryRun(v bool) *CreateDBInstanceResponseBody {
	s.DryRun = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetDryRunResult(v bool) *CreateDBInstanceResponseBody {
	s.DryRunResult = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetMessage(v string) *CreateDBInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetOrderId(v string) *CreateDBInstanceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetPort(v string) *CreateDBInstanceResponseBody {
	s.Port = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetRequestId(v string) *CreateDBInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetTagResult(v bool) *CreateDBInstanceResponseBody {
	s.TagResult = &v
	return s
}

func (s *CreateDBInstanceResponseBody) SetTaskId(v string) *CreateDBInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateDBInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
