// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDryRunSwaggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailed(v *DryRunSwaggerResponseBodyFailed) *DryRunSwaggerResponseBody
	GetFailed() *DryRunSwaggerResponseBodyFailed
	SetGlobalCondition(v string) *DryRunSwaggerResponseBody
	GetGlobalCondition() *string
	SetModelFailed(v *DryRunSwaggerResponseBodyModelFailed) *DryRunSwaggerResponseBody
	GetModelFailed() *DryRunSwaggerResponseBodyModelFailed
	SetModelSuccess(v *DryRunSwaggerResponseBodyModelSuccess) *DryRunSwaggerResponseBody
	GetModelSuccess() *DryRunSwaggerResponseBodyModelSuccess
	SetRequestId(v string) *DryRunSwaggerResponseBody
	GetRequestId() *string
	SetSuccess(v *DryRunSwaggerResponseBodySuccess) *DryRunSwaggerResponseBody
	GetSuccess() *DryRunSwaggerResponseBodySuccess
}

type DryRunSwaggerResponseBody struct {
	// The APIs that failed to be created based on the Swagger-compliant data imported this time.
	Failed *DryRunSwaggerResponseBodyFailed `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Struct"`
	// The global condition.
	//
	// example:
	//
	// {}
	GlobalCondition *string `json:"GlobalCondition,omitempty" xml:"GlobalCondition,omitempty"`
	// The models that failed to be imported through the Swagger-compliant data this time.
	ModelFailed *DryRunSwaggerResponseBodyModelFailed `json:"ModelFailed,omitempty" xml:"ModelFailed,omitempty" type:"Struct"`
	// The models that failed to be imported through the Swagger-compliant data this time.
	ModelSuccess *DryRunSwaggerResponseBodyModelSuccess `json:"ModelSuccess,omitempty" xml:"ModelSuccess,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// EF924FE4-2EDD-4CD3-89EC-34E4708574E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The APIs that are created based on the Swagger-compliant data imported this time.
	Success *DryRunSwaggerResponseBodySuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Struct"`
}

func (s DryRunSwaggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBody) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBody) GetFailed() *DryRunSwaggerResponseBodyFailed {
	return s.Failed
}

func (s *DryRunSwaggerResponseBody) GetGlobalCondition() *string {
	return s.GlobalCondition
}

func (s *DryRunSwaggerResponseBody) GetModelFailed() *DryRunSwaggerResponseBodyModelFailed {
	return s.ModelFailed
}

func (s *DryRunSwaggerResponseBody) GetModelSuccess() *DryRunSwaggerResponseBodyModelSuccess {
	return s.ModelSuccess
}

func (s *DryRunSwaggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DryRunSwaggerResponseBody) GetSuccess() *DryRunSwaggerResponseBodySuccess {
	return s.Success
}

func (s *DryRunSwaggerResponseBody) SetFailed(v *DryRunSwaggerResponseBodyFailed) *DryRunSwaggerResponseBody {
	s.Failed = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetGlobalCondition(v string) *DryRunSwaggerResponseBody {
	s.GlobalCondition = &v
	return s
}

func (s *DryRunSwaggerResponseBody) SetModelFailed(v *DryRunSwaggerResponseBodyModelFailed) *DryRunSwaggerResponseBody {
	s.ModelFailed = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetModelSuccess(v *DryRunSwaggerResponseBodyModelSuccess) *DryRunSwaggerResponseBody {
	s.ModelSuccess = v
	return s
}

func (s *DryRunSwaggerResponseBody) SetRequestId(v string) *DryRunSwaggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DryRunSwaggerResponseBody) SetSuccess(v *DryRunSwaggerResponseBodySuccess) *DryRunSwaggerResponseBody {
	s.Success = v
	return s
}

func (s *DryRunSwaggerResponseBody) Validate() error {
	if s.Failed != nil {
		if err := s.Failed.Validate(); err != nil {
			return err
		}
	}
	if s.ModelFailed != nil {
		if err := s.ModelFailed.Validate(); err != nil {
			return err
		}
	}
	if s.ModelSuccess != nil {
		if err := s.ModelSuccess.Validate(); err != nil {
			return err
		}
	}
	if s.Success != nil {
		if err := s.Success.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DryRunSwaggerResponseBodyFailed struct {
	ApiImportSwaggerFailed []*DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed `json:"ApiImportSwaggerFailed,omitempty" xml:"ApiImportSwaggerFailed,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyFailed) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyFailed) GetApiImportSwaggerFailed() []*DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	return s.ApiImportSwaggerFailed
}

func (s *DryRunSwaggerResponseBodyFailed) SetApiImportSwaggerFailed(v []*DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) *DryRunSwaggerResponseBodyFailed {
	s.ApiImportSwaggerFailed = v
	return s
}

func (s *DryRunSwaggerResponseBodyFailed) Validate() error {
	if s.ApiImportSwaggerFailed != nil {
		for _, item := range s.ApiImportSwaggerFailed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed struct {
	// The error message returned when the API is created.
	//
	// example:
	//
	// api already exists : apiUid ===> 8e274ec61cf6468e83b68371956831cb
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The HTTP method configured when the API is created.
	//
	// example:
	//
	// post
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request path configured when the API is created.
	//
	// example:
	//
	// /http/get/mapping
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) GetPath() *string {
	return s.Path
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetErrorMsg(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.ErrorMsg = &v
	return s
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetHttpMethod(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.HttpMethod = &v
	return s
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) SetPath(v string) *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.Path = &v
	return s
}

func (s *DryRunSwaggerResponseBodyFailedApiImportSwaggerFailed) Validate() error {
	return dara.Validate(s)
}

type DryRunSwaggerResponseBodyModelFailed struct {
	ApiImportModelFailed []*DryRunSwaggerResponseBodyModelFailedApiImportModelFailed `json:"ApiImportModelFailed,omitempty" xml:"ApiImportModelFailed,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyModelFailed) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelFailed) GetApiImportModelFailed() []*DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	return s.ApiImportModelFailed
}

func (s *DryRunSwaggerResponseBodyModelFailed) SetApiImportModelFailed(v []*DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) *DryRunSwaggerResponseBodyModelFailed {
	s.ApiImportModelFailed = v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailed) Validate() error {
	if s.ApiImportModelFailed != nil {
		for _, item := range s.ApiImportModelFailed {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DryRunSwaggerResponseBodyModelFailedApiImportModelFailed struct {
	// The error message.
	//
	// example:
	//
	// Not Found
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// The ID of the API group.
	//
	// example:
	//
	// 36d4bcfaec1946e1870d90b2d7519710
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// Region
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
}

func (s DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) GetGroupId() *string {
	return s.GroupId
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) GetModelName() *string {
	return s.ModelName
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetErrorMsg(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ErrorMsg = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetGroupId(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) SetModelName(v string) *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ModelName = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelFailedApiImportModelFailed) Validate() error {
	return dara.Validate(s)
}

type DryRunSwaggerResponseBodyModelSuccess struct {
	ApiImportModelSuccess []*DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess `json:"ApiImportModelSuccess,omitempty" xml:"ApiImportModelSuccess,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodyModelSuccess) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelSuccess) GetApiImportModelSuccess() []*DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	return s.ApiImportModelSuccess
}

func (s *DryRunSwaggerResponseBodyModelSuccess) SetApiImportModelSuccess(v []*DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) *DryRunSwaggerResponseBodyModelSuccess {
	s.ApiImportModelSuccess = v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccess) Validate() error {
	if s.ApiImportModelSuccess != nil {
		for _, item := range s.ApiImportModelSuccess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess struct {
	// The ID of the API group.
	//
	// example:
	//
	// b2d552ed90ca435b86f7bf8d45414793
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the model.
	//
	// example:
	//
	// NewInstance
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The model operation.
	//
	// example:
	//
	// CREATE
	ModelOperation *string `json:"ModelOperation,omitempty" xml:"ModelOperation,omitempty"`
	// The UID of the model.
	//
	// example:
	//
	// ec1946e1870d90b2d7519
	ModelUid *string `json:"ModelUid,omitempty" xml:"ModelUid,omitempty"`
}

func (s DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetGroupId() *string {
	return s.GroupId
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelName() *string {
	return s.ModelName
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelOperation() *string {
	return s.ModelOperation
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelUid() *string {
	return s.ModelUid
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetGroupId(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.GroupId = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelName(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelName = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelOperation(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelOperation = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelUid(v string) *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelUid = &v
	return s
}

func (s *DryRunSwaggerResponseBodyModelSuccessApiImportModelSuccess) Validate() error {
	return dara.Validate(s)
}

type DryRunSwaggerResponseBodySuccess struct {
	ApiDryRunSwaggerSuccess []*DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess `json:"ApiDryRunSwaggerSuccess,omitempty" xml:"ApiDryRunSwaggerSuccess,omitempty" type:"Repeated"`
}

func (s DryRunSwaggerResponseBodySuccess) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodySuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodySuccess) GetApiDryRunSwaggerSuccess() []*DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	return s.ApiDryRunSwaggerSuccess
}

func (s *DryRunSwaggerResponseBodySuccess) SetApiDryRunSwaggerSuccess(v []*DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) *DryRunSwaggerResponseBodySuccess {
	s.ApiDryRunSwaggerSuccess = v
	return s
}

func (s *DryRunSwaggerResponseBodySuccess) Validate() error {
	if s.ApiDryRunSwaggerSuccess != nil {
		for _, item := range s.ApiDryRunSwaggerSuccess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess struct {
	// Specifies whether the operation is CREATE or MODIFY.
	//
	// example:
	//
	// CREATE
	ApiOperation *string `json:"ApiOperation,omitempty" xml:"ApiOperation,omitempty"`
	// The API definition that complies with the Swagger specification.
	//
	// example:
	//
	// "A Swagger API definition in YAML"
	ApiSwagger *string `json:"ApiSwagger,omitempty" xml:"ApiSwagger,omitempty"`
	// The UID of the successfully imported API.
	//
	// example:
	//
	// 8e274ec61cf6468e83b68371956831cb
	ApiUid *string `json:"ApiUid,omitempty" xml:"ApiUid,omitempty"`
	// The HTTP method configured when the API is created.
	//
	// example:
	//
	// get
	HttpMethod *string `json:"HttpMethod,omitempty" xml:"HttpMethod,omitempty"`
	// The request path configured when the API is created.
	//
	// example:
	//
	// /http/get/mapping
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
}

func (s DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) String() string {
	return dara.Prettify(s)
}

func (s DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GoString() string {
	return s.String()
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GetApiOperation() *string {
	return s.ApiOperation
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GetApiSwagger() *string {
	return s.ApiSwagger
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GetApiUid() *string {
	return s.ApiUid
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) GetPath() *string {
	return s.Path
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiOperation(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiOperation = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiSwagger(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiSwagger = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetApiUid(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.ApiUid = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetHttpMethod(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.HttpMethod = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) SetPath(v string) *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess {
	s.Path = &v
	return s
}

func (s *DryRunSwaggerResponseBodySuccessApiDryRunSwaggerSuccess) Validate() error {
	return dara.Validate(s)
}
