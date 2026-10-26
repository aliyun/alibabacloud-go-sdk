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
	// rm-uf6wjk5****.mysql.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// The instance ID. If the **Amount*	- parameter is set to a value greater than **1**, the number of instance IDs returned corresponds to the value, separated by commas.
	//
	// For example, if the **Amount*	- parameter is set to **3**, three instance IDs are returned. Example:
	//
	// `rm-uf6wjk5*****1, rm-uf6wjk5*****2, rm-uf6wjk5*****3`
	//
	// example:
	//
	// rm-uf6wjk5****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Indicates whether the current request requires a precheck before creating the instance.
	//
	// 	- The return value is fixed to **true**.
	//
	// 	- If no precheck is performed, this parameter is not returned.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Indicates whether the precheck for creating the instance is successful. Valid values:
	//
	// 	- **true**: Passed.
	//
	// 	- **false**: Failed.
	//
	// > 	- If no precheck is performed, this parameter is not returned.
	//
	// > 	- If the precheck fails, the corresponding error is returned.
	//
	// example:
	//
	// true
	DryRunResult *bool `json:"DryRunResult,omitempty" xml:"DryRunResult,omitempty"`
	// The prompt message for the batch creation task.
	//
	// > This parameter is returned only when the **Amount*	- parameter is greater than 1.
	//
	// example:
	//
	// Batch Create DBInstance Task Is In Process.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The order ID.
	//
	// example:
	//
	// 1007893702****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The port number corresponding to the internal endpoint of the instance.
	//
	// example:
	//
	// 3306
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1E43AAE0-BEE8-43DA-860D-EAF2AA0724DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the instance is successfully tagged. Valid values:
	//
	// 	- **true**: Succeeded.
	//
	// 	- **false**: Failed.
	//
	// > If no tag is bound to the instance, this parameter is not returned.
	//
	// example:
	//
	// true
	TagResult *bool `json:"TagResult,omitempty" xml:"TagResult,omitempty"`
	// The task ID of the batch creation task.
	//
	// 	- This parameter is returned only when the **Amount*	- parameter is greater than 1.
	//
	// 	- Currently, you cannot view the task by using the value of **TaskId**.
	//
	// example:
	//
	// s2365879-a9d0-55af-fgae-f2****
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
