// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBusinessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetBusinessResponseBodyData) *GetBusinessResponseBody
	GetData() *GetBusinessResponseBodyData
	SetErrorCode(v string) *GetBusinessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetBusinessResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetBusinessResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetBusinessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBusinessResponseBody
	GetSuccess() *bool
}

type GetBusinessResponseBody struct {
	// Details of the workflow.
	Data *GetBusinessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. Used for troubleshooting when an error occurs.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: success.
	//
	// 	- false: failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBusinessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessResponseBody) GoString() string {
	return s.String()
}

func (s *GetBusinessResponseBody) GetData() *GetBusinessResponseBodyData {
	return s.Data
}

func (s *GetBusinessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetBusinessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetBusinessResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBusinessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBusinessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBusinessResponseBody) SetData(v *GetBusinessResponseBodyData) *GetBusinessResponseBody {
	s.Data = v
	return s
}

func (s *GetBusinessResponseBody) SetErrorCode(v string) *GetBusinessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetBusinessResponseBody) SetErrorMessage(v string) *GetBusinessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetBusinessResponseBody) SetHttpStatusCode(v int32) *GetBusinessResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBusinessResponseBody) SetRequestId(v string) *GetBusinessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBusinessResponseBody) SetSuccess(v bool) *GetBusinessResponseBody {
	s.Success = &v
	return s
}

func (s *GetBusinessResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBusinessResponseBodyData struct {
	// The workflow ID.
	//
	// example:
	//
	// 1000001
	BusinessId *int64 `json:"BusinessId,omitempty" xml:"BusinessId,omitempty"`
	// The name of the workflow. Workflow names must be unique within the same workspace.
	//
	// example:
	//
	// The first business process
	BusinessName *string `json:"BusinessName,omitempty" xml:"BusinessName,omitempty"`
	// The description of the workflow.
	//
	// example:
	//
	// This is my first business process.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The Alibaba Cloud account ID of the workflow owner.
	//
	// example:
	//
	// 20000****
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The ID of the workspace where the workflow resides.
	//
	// example:
	//
	// 10000
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The functional module to which the workflow belongs. Valid values: NORMAL (Data Studio) and MANUAL_BIZ (Manually Triggered Workflow)
	//
	// example:
	//
	// NORMAL
	UseType *string `json:"UseType,omitempty" xml:"UseType,omitempty"`
}

func (s GetBusinessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBusinessResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBusinessResponseBodyData) GetBusinessId() *int64 {
	return s.BusinessId
}

func (s *GetBusinessResponseBodyData) GetBusinessName() *string {
	return s.BusinessName
}

func (s *GetBusinessResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetBusinessResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetBusinessResponseBodyData) GetProjectId() *string {
	return s.ProjectId
}

func (s *GetBusinessResponseBodyData) GetUseType() *string {
	return s.UseType
}

func (s *GetBusinessResponseBodyData) SetBusinessId(v int64) *GetBusinessResponseBodyData {
	s.BusinessId = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetBusinessName(v string) *GetBusinessResponseBodyData {
	s.BusinessName = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetDescription(v string) *GetBusinessResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetOwner(v string) *GetBusinessResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetProjectId(v string) *GetBusinessResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetBusinessResponseBodyData) SetUseType(v string) *GetBusinessResponseBodyData {
	s.UseType = &v
	return s
}

func (s *GetBusinessResponseBodyData) Validate() error {
	return dara.Validate(s)
}
