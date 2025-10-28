// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationInfo(v *InsertApplicationResponseBodyApplicationInfo) *InsertApplicationResponseBody
	GetApplicationInfo() *InsertApplicationResponseBodyApplicationInfo
	SetCode(v int32) *InsertApplicationResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertApplicationResponseBody
	GetRequestId() *string
}

type InsertApplicationResponseBody struct {
	// The information about the created application.
	ApplicationInfo *InsertApplicationResponseBodyApplicationInfo `json:"ApplicationInfo,omitempty" xml:"ApplicationInfo,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// The application name test-hsy-C5039-paas-6 had been created successfully.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4264F69C-686C-4107-B493-0599C8xxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InsertApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponseBody) GetApplicationInfo() *InsertApplicationResponseBodyApplicationInfo {
	return s.ApplicationInfo
}

func (s *InsertApplicationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertApplicationResponseBody) SetApplicationInfo(v *InsertApplicationResponseBodyApplicationInfo) *InsertApplicationResponseBody {
	s.ApplicationInfo = v
	return s
}

func (s *InsertApplicationResponseBody) SetCode(v int32) *InsertApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *InsertApplicationResponseBody) SetMessage(v string) *InsertApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *InsertApplicationResponseBody) SetRequestId(v string) *InsertApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertApplicationResponseBody) Validate() error {
	if s.ApplicationInfo != nil {
		if err := s.ApplicationInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InsertApplicationResponseBodyApplicationInfo struct {
	// The ID of the application. The ID is the unique identifier of the application in EDAS.
	//
	// example:
	//
	// 6c733bcd-6efb-47a1-8226-cf722c******
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the application.
	//
	// example:
	//
	// hello-edas-test-1
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The ID of the change process.
	//
	// example:
	//
	// d0cf569e-dce3-4efb-****-08b70021****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// Indicates whether the application is a Docker application. Valid values:
	//
	// 	- **true**: The application is a Docker application.
	//
	// 	- **false**: The application is not a Docker application.
	//
	// example:
	//
	// false
	Dockerize *bool `json:"Dockerize,omitempty" xml:"Dockerize,omitempty"`
	// The owner of the application. The owner is the user who created the application.
	//
	// example:
	//
	// 249763358688********
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The port used by the created application. Default value: 8080. You can call the UpdateContainerConfiguration operation to change the port. For more information, see [UpdateContainerConfiguration](https://help.aliyun.com/document_detail/149403.html).
	//
	// example:
	//
	// 8080
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The name of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the user who created the application.
	//
	// example:
	//
	// tdy218@1362469756xxxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s InsertApplicationResponseBodyApplicationInfo) String() string {
	return dara.Prettify(s)
}

func (s InsertApplicationResponseBodyApplicationInfo) GoString() string {
	return s.String()
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetAppId() *string {
	return s.AppId
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetAppName() *string {
	return s.AppName
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetDockerize() *bool {
	return s.Dockerize
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetOwner() *string {
	return s.Owner
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetPort() *int32 {
	return s.Port
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetRegionName() *string {
	return s.RegionName
}

func (s *InsertApplicationResponseBodyApplicationInfo) GetUserId() *string {
	return s.UserId
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetAppId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.AppId = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetAppName(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.AppName = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetChangeOrderId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.ChangeOrderId = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetDockerize(v bool) *InsertApplicationResponseBodyApplicationInfo {
	s.Dockerize = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetOwner(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.Owner = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetPort(v int32) *InsertApplicationResponseBodyApplicationInfo {
	s.Port = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetRegionName(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.RegionName = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) SetUserId(v string) *InsertApplicationResponseBodyApplicationInfo {
	s.UserId = &v
	return s
}

func (s *InsertApplicationResponseBodyApplicationInfo) Validate() error {
	return dara.Validate(s)
}
