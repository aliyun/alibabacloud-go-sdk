// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportSwaggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailed(v *ImportSwaggerResponseBodyFailed) *ImportSwaggerResponseBody
	GetFailed() *ImportSwaggerResponseBodyFailed
	SetModelFailed(v *ImportSwaggerResponseBodyModelFailed) *ImportSwaggerResponseBody
	GetModelFailed() *ImportSwaggerResponseBodyModelFailed
	SetModelSuccess(v *ImportSwaggerResponseBodyModelSuccess) *ImportSwaggerResponseBody
	GetModelSuccess() *ImportSwaggerResponseBodyModelSuccess
	SetRequestId(v string) *ImportSwaggerResponseBody
	GetRequestId() *string
	SetSuccess(v *ImportSwaggerResponseBodySuccess) *ImportSwaggerResponseBody
	GetSuccess() *ImportSwaggerResponseBodySuccess
}

type ImportSwaggerResponseBody struct {
	// The APIs that failed to be created based on the Swagger-compliant data imported this time.
	Failed *ImportSwaggerResponseBodyFailed `json:"Failed,omitempty" xml:"Failed,omitempty" type:"Struct"`
	// The models that failed to be imported through the Swagger-compliant data this time.
	ModelFailed *ImportSwaggerResponseBodyModelFailed `json:"ModelFailed,omitempty" xml:"ModelFailed,omitempty" type:"Struct"`
	// The models that were imported through the Swagger-compliant data this time.
	ModelSuccess *ImportSwaggerResponseBodyModelSuccess `json:"ModelSuccess,omitempty" xml:"ModelSuccess,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 647CEF05-404C-4125-B3D7-44792EB77392
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The APIs that are created based on the Swagger-compliant data imported this time.
	Success *ImportSwaggerResponseBodySuccess `json:"Success,omitempty" xml:"Success,omitempty" type:"Struct"`
}

func (s ImportSwaggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBody) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBody) GetFailed() *ImportSwaggerResponseBodyFailed {
	return s.Failed
}

func (s *ImportSwaggerResponseBody) GetModelFailed() *ImportSwaggerResponseBodyModelFailed {
	return s.ModelFailed
}

func (s *ImportSwaggerResponseBody) GetModelSuccess() *ImportSwaggerResponseBodyModelSuccess {
	return s.ModelSuccess
}

func (s *ImportSwaggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportSwaggerResponseBody) GetSuccess() *ImportSwaggerResponseBodySuccess {
	return s.Success
}

func (s *ImportSwaggerResponseBody) SetFailed(v *ImportSwaggerResponseBodyFailed) *ImportSwaggerResponseBody {
	s.Failed = v
	return s
}

func (s *ImportSwaggerResponseBody) SetModelFailed(v *ImportSwaggerResponseBodyModelFailed) *ImportSwaggerResponseBody {
	s.ModelFailed = v
	return s
}

func (s *ImportSwaggerResponseBody) SetModelSuccess(v *ImportSwaggerResponseBodyModelSuccess) *ImportSwaggerResponseBody {
	s.ModelSuccess = v
	return s
}

func (s *ImportSwaggerResponseBody) SetRequestId(v string) *ImportSwaggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportSwaggerResponseBody) SetSuccess(v *ImportSwaggerResponseBodySuccess) *ImportSwaggerResponseBody {
	s.Success = v
	return s
}

func (s *ImportSwaggerResponseBody) Validate() error {
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

type ImportSwaggerResponseBodyFailed struct {
	ApiImportSwaggerFailed []*ImportSwaggerResponseBodyFailedApiImportSwaggerFailed `json:"ApiImportSwaggerFailed,omitempty" xml:"ApiImportSwaggerFailed,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyFailed) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyFailed) GetApiImportSwaggerFailed() []*ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	return s.ApiImportSwaggerFailed
}

func (s *ImportSwaggerResponseBodyFailed) SetApiImportSwaggerFailed(v []*ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) *ImportSwaggerResponseBodyFailed {
	s.ApiImportSwaggerFailed = v
	return s
}

func (s *ImportSwaggerResponseBodyFailed) Validate() error {
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

type ImportSwaggerResponseBodyFailedApiImportSwaggerFailed struct {
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

func (s ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) GetPath() *string {
	return s.Path
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetErrorMsg(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.ErrorMsg = &v
	return s
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetHttpMethod(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.HttpMethod = &v
	return s
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) SetPath(v string) *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed {
	s.Path = &v
	return s
}

func (s *ImportSwaggerResponseBodyFailedApiImportSwaggerFailed) Validate() error {
	return dara.Validate(s)
}

type ImportSwaggerResponseBodyModelFailed struct {
	ApiImportModelFailed []*ImportSwaggerResponseBodyModelFailedApiImportModelFailed `json:"ApiImportModelFailed,omitempty" xml:"ApiImportModelFailed,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyModelFailed) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelFailed) GetApiImportModelFailed() []*ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	return s.ApiImportModelFailed
}

func (s *ImportSwaggerResponseBodyModelFailed) SetApiImportModelFailed(v []*ImportSwaggerResponseBodyModelFailedApiImportModelFailed) *ImportSwaggerResponseBodyModelFailed {
	s.ApiImportModelFailed = v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailed) Validate() error {
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

type ImportSwaggerResponseBodyModelFailedApiImportModelFailed struct {
	// The error message.
	//
	// example:
	//
	// error msg
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

func (s ImportSwaggerResponseBodyModelFailedApiImportModelFailed) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelFailedApiImportModelFailed) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) GetModelName() *string {
	return s.ModelName
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetErrorMsg(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ErrorMsg = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetGroupId(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) SetModelName(v string) *ImportSwaggerResponseBodyModelFailedApiImportModelFailed {
	s.ModelName = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelFailedApiImportModelFailed) Validate() error {
	return dara.Validate(s)
}

type ImportSwaggerResponseBodyModelSuccess struct {
	ApiImportModelSuccess []*ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess `json:"ApiImportModelSuccess,omitempty" xml:"ApiImportModelSuccess,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodyModelSuccess) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelSuccess) GetApiImportModelSuccess() []*ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	return s.ApiImportModelSuccess
}

func (s *ImportSwaggerResponseBodyModelSuccess) SetApiImportModelSuccess(v []*ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) *ImportSwaggerResponseBodyModelSuccess {
	s.ApiImportModelSuccess = v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccess) Validate() error {
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

type ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess struct {
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
	// The model operation
	//
	// example:
	//
	// CREATE
	ModelOperation *string `json:"ModelOperation,omitempty" xml:"ModelOperation,omitempty"`
	// The UID of the model.
	//
	// example:
	//
	// d4bcfaec1946e1870d
	ModelUid *string `json:"ModelUid,omitempty" xml:"ModelUid,omitempty"`
}

func (s ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetGroupId() *string {
	return s.GroupId
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelName() *string {
	return s.ModelName
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelOperation() *string {
	return s.ModelOperation
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) GetModelUid() *string {
	return s.ModelUid
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetGroupId(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.GroupId = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelName(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelName = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelOperation(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelOperation = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) SetModelUid(v string) *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess {
	s.ModelUid = &v
	return s
}

func (s *ImportSwaggerResponseBodyModelSuccessApiImportModelSuccess) Validate() error {
	return dara.Validate(s)
}

type ImportSwaggerResponseBodySuccess struct {
	ApiImportSwaggerSuccess []*ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess `json:"ApiImportSwaggerSuccess,omitempty" xml:"ApiImportSwaggerSuccess,omitempty" type:"Repeated"`
}

func (s ImportSwaggerResponseBodySuccess) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodySuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodySuccess) GetApiImportSwaggerSuccess() []*ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	return s.ApiImportSwaggerSuccess
}

func (s *ImportSwaggerResponseBodySuccess) SetApiImportSwaggerSuccess(v []*ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) *ImportSwaggerResponseBodySuccess {
	s.ApiImportSwaggerSuccess = v
	return s
}

func (s *ImportSwaggerResponseBodySuccess) Validate() error {
	if s.ApiImportSwaggerSuccess != nil {
		for _, item := range s.ApiImportSwaggerSuccess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess struct {
	// Specifies whether the operation is CREATE or MODIFY.
	//
	// example:
	//
	// CREATE
	ApiOperation *string `json:"ApiOperation,omitempty" xml:"ApiOperation,omitempty"`
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

func (s ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) String() string {
	return dara.Prettify(s)
}

func (s ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GoString() string {
	return s.String()
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GetApiOperation() *string {
	return s.ApiOperation
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GetApiUid() *string {
	return s.ApiUid
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) GetPath() *string {
	return s.Path
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetApiOperation(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.ApiOperation = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetApiUid(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.ApiUid = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetHttpMethod(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.HttpMethod = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) SetPath(v string) *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess {
	s.Path = &v
	return s
}

func (s *ImportSwaggerResponseBodySuccessApiImportSwaggerSuccess) Validate() error {
	return dara.Validate(s)
}
