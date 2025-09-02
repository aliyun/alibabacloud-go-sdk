// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgPlatformQueryProjectsAndSchemaFromMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetData() []*DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData
	SetErrorCode(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody
	GetSuccess() *bool
}

type DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody struct {
	// The data returned.
	Data []*DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code.
	//
	// example:
	//
	// 1029030003
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// param error
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 400
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 102400001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GoString() string {
	return s.String()
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetData() []*DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData {
	return s.Data
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetData(v []*DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.Data = v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetErrorCode(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetErrorMessage(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetHttpStatusCode(v int32) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetRequestId(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) SetSuccess(v bool) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody {
	s.Success = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBody) Validate() error {
	return dara.Validate(s)
}

type DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData struct {
	// The ID of the EMR cluster. This parameter is generated only when the request parameter EngineName is set to EMR.
	//
	// example:
	//
	// c-12345
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The name of the compute engine.
	//
	// example:
	//
	// emr_test_project
	Project *string `json:"Project,omitempty" xml:"Project,omitempty"`
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) GoString() string {
	return s.String()
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) GetClusterId() *string {
	return s.ClusterId
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) GetProject() *string {
	return s.Project
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) SetClusterId(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData {
	s.ClusterId = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) SetProject(v string) *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData {
	s.Project = &v
	return s
}

func (s *DsgPlatformQueryProjectsAndSchemaFromMetaResponseBodyData) Validate() error {
	return dara.Validate(s)
}
