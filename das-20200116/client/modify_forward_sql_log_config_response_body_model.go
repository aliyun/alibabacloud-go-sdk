// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardSqlLogConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyForwardSqlLogConfigResponseBody
	GetCode() *string
	SetData(v *ModifyForwardSqlLogConfigResponseBodyData) *ModifyForwardSqlLogConfigResponseBody
	GetData() *ModifyForwardSqlLogConfigResponseBodyData
	SetMessage(v string) *ModifyForwardSqlLogConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyForwardSqlLogConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyForwardSqlLogConfigResponseBody
	GetSuccess() *string
}

type ModifyForwardSqlLogConfigResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ForwardSqlLogResult
	Data *ModifyForwardSqlLogConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an error message that contains information such as an error code is returned.
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
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true: The request is successful.
	//
	// 	- false: The request fails.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyForwardSqlLogConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardSqlLogConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyForwardSqlLogConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyForwardSqlLogConfigResponseBody) GetData() *ModifyForwardSqlLogConfigResponseBodyData {
	return s.Data
}

func (s *ModifyForwardSqlLogConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyForwardSqlLogConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyForwardSqlLogConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyForwardSqlLogConfigResponseBody) SetCode(v string) *ModifyForwardSqlLogConfigResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBody) SetData(v *ModifyForwardSqlLogConfigResponseBodyData) *ModifyForwardSqlLogConfigResponseBody {
	s.Data = v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBody) SetMessage(v string) *ModifyForwardSqlLogConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBody) SetRequestId(v string) *ModifyForwardSqlLogConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBody) SetSuccess(v string) *ModifyForwardSqlLogConfigResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyForwardSqlLogConfigResponseBodyData struct {
	// The LogStore name for real-time delivery to Simple Log Service.
	//
	// example:
	//
	// cdn222
	Logstore *string `json:"Logstore,omitempty" xml:"Logstore,omitempty"`
	// The project.
	//
	// example:
	//
	// facedetect7
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
	// The VPC endpoint of the component.
	//
	// example:
	//
	// cn-beijing-intranet.log.aliyuncs.com
	VpcEndpoint *string `json:"VpcEndpoint,omitempty" xml:"VpcEndpoint,omitempty"`
}

func (s ModifyForwardSqlLogConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardSqlLogConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) GetLogstore() *string {
	return s.Logstore
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) GetVpcEndpoint() *string {
	return s.VpcEndpoint
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) SetLogstore(v string) *ModifyForwardSqlLogConfigResponseBodyData {
	s.Logstore = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) SetProject(v string) *ModifyForwardSqlLogConfigResponseBodyData {
	s.Project = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) SetVpcEndpoint(v string) *ModifyForwardSqlLogConfigResponseBodyData {
	s.VpcEndpoint = &v
	return s
}

func (s *ModifyForwardSqlLogConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
