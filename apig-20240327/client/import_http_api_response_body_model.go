// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHttpApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportHttpApiResponseBody
	GetCode() *string
	SetData(v *ImportHttpApiResponseBodyData) *ImportHttpApiResponseBody
	GetData() *ImportHttpApiResponseBodyData
	SetMessage(v string) *ImportHttpApiResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportHttpApiResponseBody
	GetRequestId() *string
}

type ImportHttpApiResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The API information.
	Data *ImportHttpApiResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71E20
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ImportHttpApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBody) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportHttpApiResponseBody) GetData() *ImportHttpApiResponseBodyData {
	return s.Data
}

func (s *ImportHttpApiResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportHttpApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportHttpApiResponseBody) SetCode(v string) *ImportHttpApiResponseBody {
	s.Code = &v
	return s
}

func (s *ImportHttpApiResponseBody) SetData(v *ImportHttpApiResponseBodyData) *ImportHttpApiResponseBody {
	s.Data = v
	return s
}

func (s *ImportHttpApiResponseBody) SetMessage(v string) *ImportHttpApiResponseBody {
	s.Message = &v
	return s
}

func (s *ImportHttpApiResponseBody) SetRequestId(v string) *ImportHttpApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportHttpApiResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportHttpApiResponseBodyData struct {
	// The dry run result.
	DryRunInfo *ImportHttpApiResponseBodyDataDryRunInfo `json:"dryRunInfo,omitempty" xml:"dryRunInfo,omitempty" type:"Struct"`
	// The API ID.
	//
	// example:
	//
	// api-xxx
	HttpApiId *string `json:"httpApiId,omitempty" xml:"httpApiId,omitempty"`
	// The API name.
	//
	// example:
	//
	// import-test
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyData) GetDryRunInfo() *ImportHttpApiResponseBodyDataDryRunInfo {
	return s.DryRunInfo
}

func (s *ImportHttpApiResponseBodyData) GetHttpApiId() *string {
	return s.HttpApiId
}

func (s *ImportHttpApiResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ImportHttpApiResponseBodyData) SetDryRunInfo(v *ImportHttpApiResponseBodyDataDryRunInfo) *ImportHttpApiResponseBodyData {
	s.DryRunInfo = v
	return s
}

func (s *ImportHttpApiResponseBodyData) SetHttpApiId(v string) *ImportHttpApiResponseBodyData {
	s.HttpApiId = &v
	return s
}

func (s *ImportHttpApiResponseBodyData) SetName(v string) *ImportHttpApiResponseBodyData {
	s.Name = &v
	return s
}

func (s *ImportHttpApiResponseBodyData) Validate() error {
	if s.DryRunInfo != nil {
		if err := s.DryRunInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportHttpApiResponseBodyDataDryRunInfo struct {
	// The error messages. If an error message is returned, the API fails to be imported.
	ErrorMessages []*string `json:"errorMessages,omitempty" xml:"errorMessages,omitempty" type:"Repeated"`
	// The existing APIs. If an existing API is returned, the import updates the existing API.
	ExistHttpApiInfo *HttpApiApiInfo `json:"existHttpApiInfo,omitempty" xml:"existHttpApiInfo,omitempty"`
	// The data structs that fail the dry run.
	FailureComponents []*ImportHttpApiResponseBodyDataDryRunInfoFailureComponents `json:"failureComponents,omitempty" xml:"failureComponents,omitempty" type:"Repeated"`
	// The operations that fail the dry run.
	FailureOperations []*ImportHttpApiResponseBodyDataDryRunInfoFailureOperations `json:"failureOperations,omitempty" xml:"failureOperations,omitempty" type:"Repeated"`
	// The data structs that pass the dry run.
	SuccessComponents []*ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents `json:"successComponents,omitempty" xml:"successComponents,omitempty" type:"Repeated"`
	// The operations that pass the dry run.
	SuccessOperations []*ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations `json:"successOperations,omitempty" xml:"successOperations,omitempty" type:"Repeated"`
	// The alerts. If an alert is returned, specific operations or structs may fail to be imported.
	WarningMessages []*string `json:"warningMessages,omitempty" xml:"warningMessages,omitempty" type:"Repeated"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfo) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfo) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetErrorMessages() []*string {
	return s.ErrorMessages
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetExistHttpApiInfo() *HttpApiApiInfo {
	return s.ExistHttpApiInfo
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetFailureComponents() []*ImportHttpApiResponseBodyDataDryRunInfoFailureComponents {
	return s.FailureComponents
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetFailureOperations() []*ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	return s.FailureOperations
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetSuccessComponents() []*ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents {
	return s.SuccessComponents
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetSuccessOperations() []*ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	return s.SuccessOperations
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) GetWarningMessages() []*string {
	return s.WarningMessages
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetErrorMessages(v []*string) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.ErrorMessages = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetExistHttpApiInfo(v *HttpApiApiInfo) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.ExistHttpApiInfo = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetFailureComponents(v []*ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.FailureComponents = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetFailureOperations(v []*ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.FailureOperations = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetSuccessComponents(v []*ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.SuccessComponents = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetSuccessOperations(v []*ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.SuccessOperations = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) SetWarningMessages(v []*string) *ImportHttpApiResponseBodyDataDryRunInfo {
	s.WarningMessages = v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfo) Validate() error {
	if s.ExistHttpApiInfo != nil {
		if err := s.ExistHttpApiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FailureComponents != nil {
		for _, item := range s.FailureComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FailureOperations != nil {
		for _, item := range s.FailureOperations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessComponents != nil {
		for _, item := range s.SuccessComponents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessOperations != nil {
		for _, item := range s.SuccessOperations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportHttpApiResponseBodyDataDryRunInfoFailureComponents struct {
	// The error message.
	//
	// example:
	//
	// The data struct is incorrectly defined.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The data struct name.
	//
	// example:
	//
	// orderDTO
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) GetName() *string {
	return s.Name
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) SetErrorMessage(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents {
	s.ErrorMessage = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents {
	s.Name = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureComponents) Validate() error {
	return dara.Validate(s)
}

type ImportHttpApiResponseBodyDataDryRunInfoFailureOperations struct {
	// The error message.
	//
	// example:
	//
	// Missing response definition.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP method of the operation.
	//
	// example:
	//
	// GET
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The operation path.
	//
	// example:
	//
	// /v1/orders
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) GetMethod() *string {
	return s.Method
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) GetPath() *string {
	return s.Path
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetErrorMessage(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.ErrorMessage = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetMethod(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.Method = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) SetPath(v string) *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations {
	s.Path = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoFailureOperations) Validate() error {
	return dara.Validate(s)
}

type ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents struct {
	// The action that will be performed for the data struct after the dry run.
	//
	// 	- Create: The data struct is created.
	//
	// 	- Update: The data struct is updated.
	//
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The data struct name.
	//
	// example:
	//
	// userDTO
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) GetAction() *string {
	return s.Action
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) GetName() *string {
	return s.Name
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) SetAction(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents {
	s.Action = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents {
	s.Name = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessComponents) Validate() error {
	return dara.Validate(s)
}

type ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations struct {
	// The action that will be performed for the operation after the dry run.
	//
	// 	- Create: The operation is created.
	//
	// 	- Update: The operation is updated.
	//
	// example:
	//
	// Create
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The HTTP method of the operation.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The operation name.
	//
	// example:
	//
	// CreateUser
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The operation path.
	//
	// example:
	//
	// /v1/users
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GetAction() *string {
	return s.Action
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GetMethod() *string {
	return s.Method
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GetName() *string {
	return s.Name
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) GetPath() *string {
	return s.Path
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetAction(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Action = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetMethod(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Method = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetName(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Name = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) SetPath(v string) *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations {
	s.Path = &v
	return s
}

func (s *ImportHttpApiResponseBodyDataDryRunInfoSuccessOperations) Validate() error {
	return dara.Validate(s)
}
