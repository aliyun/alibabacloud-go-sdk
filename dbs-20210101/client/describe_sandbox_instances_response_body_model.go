// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSandboxInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSandboxInstancesResponseBody
	GetCode() *string
	SetData(v string) *DescribeSandboxInstancesResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeSandboxInstancesResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeSandboxInstancesResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeSandboxInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSandboxInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeSandboxInstancesResponseBody
	GetSuccess() *string
}

type DescribeSandboxInstancesResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	//
	// 	- **connectionString**: the connection string of the sandbox instance, in the format of IP address:Port number. This parameter indicates the endpoint of the sandbox instance if the value of the SandboxType parameter is **Sandbox**. This parameter indicates the Network File System (NFS) mount address if the value of the SandboxType parameter is **NFS**.
	//
	// 	- **restoreSeconds**: the time required to create the sandbox instance. Unit: seconds.
	//
	// 	- **restoreTime**: the point in time to which the sandbox instance is restored. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **instanceId**: the ID of the sandbox instance.
	//
	// 	- **backupSetId**: the ID of the backup set.
	//
	// 	- **createTime**: the point in time when the sandbox instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mm:ssZ format. The time is displayed in UTC.
	//
	// 	- **backupPlanId**: the ID of the backup schedule.
	//
	// 	- **vpcId**: the ID of the virtual private cloud (VPC).
	//
	// 	- **vpcSwitchId**: the ID of the VSwitch.
	//
	// 	- **sandboxSpecification**: the specifications of the sandbox instance.
	//
	// 	- **status**: the status of the sandbox instance. Valid values: **running**, **check_pass**, and **stop**.
	//
	// example:
	//
	// {     "number": 1,     "size": 1,     "content": [       {         "connectionString": "172.26.178.229:3306",         "restoreSeconds": 15,         "restoreTime": "2021-08-11T07:26:24Z",         "instanceId": "1jxxxxx9xxxms",         "backupSetId": "1hxxxx8xxxxxa_20210811152624",         "createTime": "2021-08-12T07:40:29Z",         "backupPlanId": "1hxxxx8xxxxxa",         "vpcId": "vpc-bp1dxxxxxjy0xxxxx1xxp",         "sandboxSpecification": "MYSQL_1C_1M_SD",         "status": "running",         "vpcSwitchId": "vsw-bp1bxxxxxumxxxxxwxx2w"       }     ],     "totalElements": 1   }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned if the request fails.
	//
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4F1888AC-1138-4995-B9FE-D2734F61C058
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSandboxInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSandboxInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSandboxInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSandboxInstancesResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeSandboxInstancesResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeSandboxInstancesResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeSandboxInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSandboxInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSandboxInstancesResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeSandboxInstancesResponseBody) SetCode(v string) *DescribeSandboxInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetData(v string) *DescribeSandboxInstancesResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrCode(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetErrMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetMessage(v string) *DescribeSandboxInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetRequestId(v string) *DescribeSandboxInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) SetSuccess(v string) *DescribeSandboxInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSandboxInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
