// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDBInstanceConnectivityDiagnosisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDBInstanceConnectivityDiagnosisResponseBody
	GetCode() *string
	SetData(v *GetDBInstanceConnectivityDiagnosisResponseBodyData) *GetDBInstanceConnectivityDiagnosisResponseBody
	GetData() *GetDBInstanceConnectivityDiagnosisResponseBodyData
	SetMessage(v string) *GetDBInstanceConnectivityDiagnosisResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDBInstanceConnectivityDiagnosisResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetDBInstanceConnectivityDiagnosisResponseBody
	GetSuccess() *string
}

type GetDBInstanceConnectivityDiagnosisResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The detailed information.
	Data *GetDBInstanceConnectivityDiagnosisResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request was successful, **Successful*	- is returned. Otherwise, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDBInstanceConnectivityDiagnosisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceConnectivityDiagnosisResponseBody) GoString() string {
	return s.String()
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) GetData() *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	return s.Data
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) SetCode(v string) *GetDBInstanceConnectivityDiagnosisResponseBody {
	s.Code = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) SetData(v *GetDBInstanceConnectivityDiagnosisResponseBodyData) *GetDBInstanceConnectivityDiagnosisResponseBody {
	s.Data = v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) SetMessage(v string) *GetDBInstanceConnectivityDiagnosisResponseBody {
	s.Message = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) SetRequestId(v string) *GetDBInstanceConnectivityDiagnosisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) SetSuccess(v string) *GetDBInstanceConnectivityDiagnosisResponseBody {
	s.Success = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDBInstanceConnectivityDiagnosisResponseBodyData struct {
	// The exception detection items:
	//
	// 	- **SRC_IP_NOT_IN_USER_WHITELIST**: The source IP address is not added to the whitelist of the user.
	//
	// 	- **VIP_NOT_EXISTS**: The Application Load Balancer (ALB) instance corresponding to the virtual IP address (VIP) does not exist.
	//
	// 	- **RS_NOT_EXISTS**: The resource sharing (RS) is not properly mounted.
	//
	// 	- **VIP_TUNNEL_ID_NOT_CONSISTENT**: The tunnel ID used by the VIP of the virtual private cloud (VPC) type is different from the tunnel ID of the VPC.
	//
	// 	- **VIP_VPC_CLOUD_INSTANCE_NOT_EXISTS**: The VIP of the VPC type does not exist.
	//
	// 	- **VIP_IS_NOT_NGLB**: The NGLB mode is disabled for the VIP.
	//
	// 	- **CUSTINS_NOT_ASSOCIATE_ECS_SECURITY_GROUP**: No security group is associated with the instance.
	//
	// 	- **SRC_IP_NOT_IN_USER_WHITELIST**: The source IP address is not added to the whitelist of the user.
	//
	// 	- **SRC_IP_NOT_IN_ADMIN_WHITELIST**: The source IP address is not added to the whitelist of the instance.
	//
	// 	- **SRC_IP_NOT_IN_ECS_SECURITY_GROUP**: The source IP address is not added to the security group that is associated with the instance.
	//
	// 	- **VPC_INSTANCE_IP_NOT_WORKING_STATUS**: The IP address in the VPC is in an abnormal state.
	//
	// example:
	//
	// SRC_IP_NOT_IN_USER_WHITELIST
	ConnCheckErrorCode *string `json:"connCheckErrorCode,omitempty" xml:"connCheckErrorCode,omitempty"`
	// The details of the exception detection.
	//
	// example:
	//
	// Src ip:47.110.180.62 not in user whitelist
	ConnCheckErrorMessage *string `json:"connCheckErrorMessage,omitempty" xml:"connCheckErrorMessage,omitempty"`
	// The type of the exception:
	//
	// 	- **0**: an exception that can be handled by the user.
	//
	// 	- **1**: an exception that can be handled by a technical engineer.
	//
	// example:
	//
	// 0
	FailType *string `json:"failType,omitempty" xml:"failType,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// Indicates whether the connectivity test was passed:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetDBInstanceConnectivityDiagnosisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDBInstanceConnectivityDiagnosisResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) GetConnCheckErrorCode() *string {
	return s.ConnCheckErrorCode
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) GetConnCheckErrorMessage() *string {
	return s.ConnCheckErrorMessage
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) GetFailType() *string {
	return s.FailType
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) SetConnCheckErrorCode(v string) *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	s.ConnCheckErrorCode = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) SetConnCheckErrorMessage(v string) *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	s.ConnCheckErrorMessage = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) SetFailType(v string) *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	s.FailType = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) SetInstanceId(v string) *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) SetSuccess(v bool) *GetDBInstanceConnectivityDiagnosisResponseBodyData {
	s.Success = &v
	return s
}

func (s *GetDBInstanceConnectivityDiagnosisResponseBodyData) Validate() error {
	return dara.Validate(s)
}
