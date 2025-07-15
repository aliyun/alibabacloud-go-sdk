// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOASResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessages(v *ImportOASResponseBodyErrorMessages) *ImportOASResponseBody
	GetErrorMessages() *ImportOASResponseBodyErrorMessages
	SetFailedApis(v *ImportOASResponseBodyFailedApis) *ImportOASResponseBody
	GetFailedApis() *ImportOASResponseBodyFailedApis
	SetFailedModels(v *ImportOASResponseBodyFailedModels) *ImportOASResponseBody
	GetFailedModels() *ImportOASResponseBodyFailedModels
	SetOperationId(v string) *ImportOASResponseBody
	GetOperationId() *string
	SetRequestId(v string) *ImportOASResponseBody
	GetRequestId() *string
	SetSuccessApis(v *ImportOASResponseBodySuccessApis) *ImportOASResponseBody
	GetSuccessApis() *ImportOASResponseBodySuccessApis
	SetSuccessModels(v *ImportOASResponseBodySuccessModels) *ImportOASResponseBody
	GetSuccessModels() *ImportOASResponseBodySuccessModels
	SetWarningMessages(v *ImportOASResponseBodyWarningMessages) *ImportOASResponseBody
	GetWarningMessages() *ImportOASResponseBodyWarningMessages
}

type ImportOASResponseBody struct {
	// The error messages that appear due to the invalid data in the imported file.
	ErrorMessages *ImportOASResponseBodyErrorMessages `json:"ErrorMessages,omitempty" xml:"ErrorMessages,omitempty" type:"Struct"`
	// The APIs that failed to pass the precheck.
	FailedApis *ImportOASResponseBodyFailedApis `json:"FailedApis,omitempty" xml:"FailedApis,omitempty" type:"Struct"`
	// The information about the models that failed to pass the precheck.
	FailedModels *ImportOASResponseBodyFailedModels `json:"FailedModels,omitempty" xml:"FailedModels,omitempty" type:"Struct"`
	// The ID of the asynchronous API import task that was generated during the import operation. This ID is used to query the execution status of the API import task.
	//
	// example:
	//
	// c16a1880f5164d779f6a54f64d997cd9
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E7FE7172-AA75-5880-B6F7-C00893E9BC06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the APIs that have passed the precheck.
	SuccessApis *ImportOASResponseBodySuccessApis `json:"SuccessApis,omitempty" xml:"SuccessApis,omitempty" type:"Struct"`
	// The information about the models that have passed the precheck.
	SuccessModels *ImportOASResponseBodySuccessModels `json:"SuccessModels,omitempty" xml:"SuccessModels,omitempty" type:"Struct"`
	// The warning messages that appear due to the invalid data in the imported file.
	WarningMessages *ImportOASResponseBodyWarningMessages `json:"WarningMessages,omitempty" xml:"WarningMessages,omitempty" type:"Struct"`
}

func (s ImportOASResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBody) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBody) GetErrorMessages() *ImportOASResponseBodyErrorMessages {
	return s.ErrorMessages
}

func (s *ImportOASResponseBody) GetFailedApis() *ImportOASResponseBodyFailedApis {
	return s.FailedApis
}

func (s *ImportOASResponseBody) GetFailedModels() *ImportOASResponseBodyFailedModels {
	return s.FailedModels
}

func (s *ImportOASResponseBody) GetOperationId() *string {
	return s.OperationId
}

func (s *ImportOASResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportOASResponseBody) GetSuccessApis() *ImportOASResponseBodySuccessApis {
	return s.SuccessApis
}

func (s *ImportOASResponseBody) GetSuccessModels() *ImportOASResponseBodySuccessModels {
	return s.SuccessModels
}

func (s *ImportOASResponseBody) GetWarningMessages() *ImportOASResponseBodyWarningMessages {
	return s.WarningMessages
}

func (s *ImportOASResponseBody) SetErrorMessages(v *ImportOASResponseBodyErrorMessages) *ImportOASResponseBody {
	s.ErrorMessages = v
	return s
}

func (s *ImportOASResponseBody) SetFailedApis(v *ImportOASResponseBodyFailedApis) *ImportOASResponseBody {
	s.FailedApis = v
	return s
}

func (s *ImportOASResponseBody) SetFailedModels(v *ImportOASResponseBodyFailedModels) *ImportOASResponseBody {
	s.FailedModels = v
	return s
}

func (s *ImportOASResponseBody) SetOperationId(v string) *ImportOASResponseBody {
	s.OperationId = &v
	return s
}

func (s *ImportOASResponseBody) SetRequestId(v string) *ImportOASResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportOASResponseBody) SetSuccessApis(v *ImportOASResponseBodySuccessApis) *ImportOASResponseBody {
	s.SuccessApis = v
	return s
}

func (s *ImportOASResponseBody) SetSuccessModels(v *ImportOASResponseBodySuccessModels) *ImportOASResponseBody {
	s.SuccessModels = v
	return s
}

func (s *ImportOASResponseBody) SetWarningMessages(v *ImportOASResponseBodyWarningMessages) *ImportOASResponseBody {
	s.WarningMessages = v
	return s
}

func (s *ImportOASResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyErrorMessages struct {
	ErrorMessage []*string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodyErrorMessages) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyErrorMessages) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyErrorMessages) GetErrorMessage() []*string {
	return s.ErrorMessage
}

func (s *ImportOASResponseBodyErrorMessages) SetErrorMessage(v []*string) *ImportOASResponseBodyErrorMessages {
	s.ErrorMessage = v
	return s
}

func (s *ImportOASResponseBodyErrorMessages) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyFailedApis struct {
	FailedApi []*ImportOASResponseBodyFailedApisFailedApi `json:"FailedApi,omitempty" xml:"FailedApi,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodyFailedApis) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyFailedApis) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyFailedApis) GetFailedApi() []*ImportOASResponseBodyFailedApisFailedApi {
	return s.FailedApi
}

func (s *ImportOASResponseBodyFailedApis) SetFailedApi(v []*ImportOASResponseBodyFailedApisFailedApi) *ImportOASResponseBodyFailedApis {
	s.FailedApi = v
	return s
}

func (s *ImportOASResponseBodyFailedApis) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyFailedApisFailedApi struct {
	// The error message.
	//
	// example:
	//
	// Invalid Api Definition.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP method configured when you created the API.
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request path configured when you created the API.
	//
	// example:
	//
	// /st1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ImportOASResponseBodyFailedApisFailedApi) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyFailedApisFailedApi) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyFailedApisFailedApi) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ImportOASResponseBodyFailedApisFailedApi) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ImportOASResponseBodyFailedApisFailedApi) GetPath() *string {
	return s.Path
}

func (s *ImportOASResponseBodyFailedApisFailedApi) SetErrorMsg(v string) *ImportOASResponseBodyFailedApisFailedApi {
	s.ErrorMsg = &v
	return s
}

func (s *ImportOASResponseBodyFailedApisFailedApi) SetHttpMethod(v string) *ImportOASResponseBodyFailedApisFailedApi {
	s.HttpMethod = &v
	return s
}

func (s *ImportOASResponseBodyFailedApisFailedApi) SetPath(v string) *ImportOASResponseBodyFailedApisFailedApi {
	s.Path = &v
	return s
}

func (s *ImportOASResponseBodyFailedApisFailedApi) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyFailedModels struct {
	FailedModel []*ImportOASResponseBodyFailedModelsFailedModel `json:"FailedModel,omitempty" xml:"FailedModel,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodyFailedModels) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyFailedModels) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyFailedModels) GetFailedModel() []*ImportOASResponseBodyFailedModelsFailedModel {
	return s.FailedModel
}

func (s *ImportOASResponseBodyFailedModels) SetFailedModel(v []*ImportOASResponseBodyFailedModelsFailedModel) *ImportOASResponseBodyFailedModels {
	s.FailedModel = v
	return s
}

func (s *ImportOASResponseBodyFailedModels) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyFailedModelsFailedModel struct {
	// The error message.
	//
	// example:
	//
	// Invalid Model Definition.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 2c1bc62e19614cc68c6b0b484bc9c5db
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s ImportOASResponseBodyFailedModelsFailedModel) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyFailedModelsFailedModel) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) GetModelName() *string {
	return s.ModelName
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) SetErrorMsg(v string) *ImportOASResponseBodyFailedModelsFailedModel {
	s.ErrorMsg = &v
	return s
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) SetGroupId(v string) *ImportOASResponseBodyFailedModelsFailedModel {
	s.GroupId = &v
	return s
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) SetModelName(v string) *ImportOASResponseBodyFailedModelsFailedModel {
	s.ModelName = &v
	return s
}

func (s *ImportOASResponseBodyFailedModelsFailedModel) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodySuccessApis struct {
	SuccessApi []*ImportOASResponseBodySuccessApisSuccessApi `json:"SuccessApi,omitempty" xml:"SuccessApi,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodySuccessApis) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodySuccessApis) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodySuccessApis) GetSuccessApi() []*ImportOASResponseBodySuccessApisSuccessApi {
	return s.SuccessApi
}

func (s *ImportOASResponseBodySuccessApis) SetSuccessApi(v []*ImportOASResponseBodySuccessApisSuccessApi) *ImportOASResponseBodySuccessApis {
	s.SuccessApi = v
	return s
}

func (s *ImportOASResponseBodySuccessApis) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodySuccessApisSuccessApi struct {
	// The ID of the API.
	//
	// example:
	//
	// 92af1abffc2443eaa2b815fdbd9c13f1
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// Indicates that the operation is CREATE or MODIFY.
	//
	// example:
	//
	// CREATE
	ApiOperation *string `json:"ApiOperation,omitempty" xml:"ApiOperation,omitempty"`
	// The HTTP method configured when you created the API.
	//
	// example:
	//
	// POST
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request path configured when you created the API.
	//
	// example:
	//
	// /st1
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s ImportOASResponseBodySuccessApisSuccessApi) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodySuccessApisSuccessApi) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) GetApiId() *string {
	return s.ApiId
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) GetApiOperation() *string {
	return s.ApiOperation
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) GetPath() *string {
	return s.Path
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) SetApiId(v string) *ImportOASResponseBodySuccessApisSuccessApi {
	s.ApiId = &v
	return s
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) SetApiOperation(v string) *ImportOASResponseBodySuccessApisSuccessApi {
	s.ApiOperation = &v
	return s
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) SetHttpMethod(v string) *ImportOASResponseBodySuccessApisSuccessApi {
	s.HttpMethod = &v
	return s
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) SetPath(v string) *ImportOASResponseBodySuccessApisSuccessApi {
	s.Path = &v
	return s
}

func (s *ImportOASResponseBodySuccessApisSuccessApi) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodySuccessModels struct {
	SuccessModel []*ImportOASResponseBodySuccessModelsSuccessModel `json:"SuccessModel,omitempty" xml:"SuccessModel,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodySuccessModels) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodySuccessModels) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodySuccessModels) GetSuccessModel() []*ImportOASResponseBodySuccessModelsSuccessModel {
	return s.SuccessModel
}

func (s *ImportOASResponseBodySuccessModels) SetSuccessModel(v []*ImportOASResponseBodySuccessModelsSuccessModel) *ImportOASResponseBodySuccessModels {
	s.SuccessModel = v
	return s
}

func (s *ImportOASResponseBodySuccessModels) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodySuccessModelsSuccessModel struct {
	// The ID of the API group.
	//
	// example:
	//
	// feaccf67040643bcbdedb253e59eb527
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// test
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The operation of the model. Valid values: CREATE and MODIFY.
	//
	// example:
	//
	// CREATE
	ModelOperation *string `json:"ModelOperation,omitempty" xml:"ModelOperation,omitempty"`
	// The UID of the model.
	//
	// example:
	//
	// 1r4efwee19614cc68c6b0b484bc9c5dbs
	ModelUid *string `json:"ModelUid,omitempty" xml:"ModelUid,omitempty"`
}

func (s ImportOASResponseBodySuccessModelsSuccessModel) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodySuccessModelsSuccessModel) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) GetModelName() *string {
	return s.ModelName
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) GetModelOperation() *string {
	return s.ModelOperation
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) GetModelUid() *string {
	return s.ModelUid
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) SetGroupId(v string) *ImportOASResponseBodySuccessModelsSuccessModel {
	s.GroupId = &v
	return s
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) SetModelName(v string) *ImportOASResponseBodySuccessModelsSuccessModel {
	s.ModelName = &v
	return s
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) SetModelOperation(v string) *ImportOASResponseBodySuccessModelsSuccessModel {
	s.ModelOperation = &v
	return s
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) SetModelUid(v string) *ImportOASResponseBodySuccessModelsSuccessModel {
	s.ModelUid = &v
	return s
}

func (s *ImportOASResponseBodySuccessModelsSuccessModel) Validate() error {
	return dara.Validate(s)
}

type ImportOASResponseBodyWarningMessages struct {
	WarningMessage []*string `json:"WarningMessage,omitempty" xml:"WarningMessage,omitempty" type:"Repeated"`
}

func (s ImportOASResponseBodyWarningMessages) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponseBodyWarningMessages) GoString() string {
	return s.String()
}

func (s *ImportOASResponseBodyWarningMessages) GetWarningMessage() []*string {
	return s.WarningMessage
}

func (s *ImportOASResponseBodyWarningMessages) SetWarningMessage(v []*string) *ImportOASResponseBodyWarningMessages {
	s.WarningMessage = v
	return s
}

func (s *ImportOASResponseBodyWarningMessages) Validate() error {
	return dara.Validate(s)
}
