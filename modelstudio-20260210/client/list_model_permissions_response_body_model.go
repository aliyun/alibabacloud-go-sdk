// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelPermissionsResponseBody
	GetCode() *string
	SetErrorMessage(v string) *ListModelPermissionsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListModelPermissionsResponseBody
	GetHttpStatusCode() *int32
	SetList(v []*ListModelPermissionsResponseBodyList) *ListModelPermissionsResponseBody
	GetList() []*ListModelPermissionsResponseBodyList
	SetMaxResults(v int32) *ListModelPermissionsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListModelPermissionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelPermissionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListModelPermissionsResponseBody
	GetTotalCount() *int64
}

type ListModelPermissionsResponseBody struct {
	// The error code. This value is empty when the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter is invalid
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The list of workspace permissions.
	List []*ListModelPermissionsResponseBodyList `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	// The maximum number of entries returned per request.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next request.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique request ID.
	//
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelPermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListModelPermissionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListModelPermissionsResponseBody) GetList() []*ListModelPermissionsResponseBodyList {
	return s.List
}

func (s *ListModelPermissionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelPermissionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelPermissionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelPermissionsResponseBody) SetCode(v string) *ListModelPermissionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetErrorMessage(v string) *ListModelPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetHttpStatusCode(v int32) *ListModelPermissionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetList(v []*ListModelPermissionsResponseBodyList) *ListModelPermissionsResponseBody {
	s.List = v
	return s
}

func (s *ListModelPermissionsResponseBody) SetMaxResults(v int32) *ListModelPermissionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetNextToken(v string) *ListModelPermissionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetRequestId(v string) *ListModelPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetSuccess(v bool) *ListModelPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelPermissionsResponseBody) SetTotalCount(v int64) *ListModelPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelPermissionsResponseBody) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelPermissionsResponseBodyList struct {
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The authorization status.
	Permissions *ListModelPermissionsResponseBodyListPermissions `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Struct"`
}

func (s ListModelPermissionsResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsResponseBodyList) GetModel() *string {
	return s.Model
}

func (s *ListModelPermissionsResponseBodyList) GetName() *string {
	return s.Name
}

func (s *ListModelPermissionsResponseBodyList) GetPermissions() *ListModelPermissionsResponseBodyListPermissions {
	return s.Permissions
}

func (s *ListModelPermissionsResponseBodyList) SetModel(v string) *ListModelPermissionsResponseBodyList {
	s.Model = &v
	return s
}

func (s *ListModelPermissionsResponseBodyList) SetName(v string) *ListModelPermissionsResponseBodyList {
	s.Name = &v
	return s
}

func (s *ListModelPermissionsResponseBodyList) SetPermissions(v *ListModelPermissionsResponseBodyListPermissions) *ListModelPermissionsResponseBodyList {
	s.Permissions = v
	return s
}

func (s *ListModelPermissionsResponseBodyList) Validate() error {
	if s.Permissions != nil {
		if err := s.Permissions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelPermissionsResponseBodyListPermissions struct {
	// The model deployment authorization. A value of true indicates that the model has been granted authorization. A value of false indicates that the model has not been granted authorization.
	//
	// example:
	//
	// true
	Deploy *bool `json:"deploy,omitempty" xml:"deploy,omitempty"`
	// The model training authorization. A value of true indicates that the model has been granted training authorization. A value of false indicates that the model has not been granted authorization.
	//
	// example:
	//
	// true
	FineTune *bool `json:"fineTune,omitempty" xml:"fineTune,omitempty"`
	// Indicates whether the model has inference permission. A value of true indicates that the model is authorized. A value of false indicates that the model is not authorized.
	//
	// example:
	//
	// true
	Inference *bool `json:"inference,omitempty" xml:"inference,omitempty"`
}

func (s ListModelPermissionsResponseBodyListPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListModelPermissionsResponseBodyListPermissions) GoString() string {
	return s.String()
}

func (s *ListModelPermissionsResponseBodyListPermissions) GetDeploy() *bool {
	return s.Deploy
}

func (s *ListModelPermissionsResponseBodyListPermissions) GetFineTune() *bool {
	return s.FineTune
}

func (s *ListModelPermissionsResponseBodyListPermissions) GetInference() *bool {
	return s.Inference
}

func (s *ListModelPermissionsResponseBodyListPermissions) SetDeploy(v bool) *ListModelPermissionsResponseBodyListPermissions {
	s.Deploy = &v
	return s
}

func (s *ListModelPermissionsResponseBodyListPermissions) SetFineTune(v bool) *ListModelPermissionsResponseBodyListPermissions {
	s.FineTune = &v
	return s
}

func (s *ListModelPermissionsResponseBodyListPermissions) SetInference(v bool) *ListModelPermissionsResponseBodyListPermissions {
	s.Inference = &v
	return s
}

func (s *ListModelPermissionsResponseBodyListPermissions) Validate() error {
	return dara.Validate(s)
}
