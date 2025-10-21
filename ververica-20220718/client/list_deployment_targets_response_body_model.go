// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DeploymentTarget) *ListDeploymentTargetsResponseBody
	GetData() []*DeploymentTarget
	SetErrorCode(v string) *ListDeploymentTargetsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDeploymentTargetsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListDeploymentTargetsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListDeploymentTargetsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentTargetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDeploymentTargetsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeploymentTargetsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListDeploymentTargetsResponseBody
	GetTotalSize() *int32
}

type ListDeploymentTargetsResponseBody struct {
	// 	- If the value of success was true, a list of clusters in which the deployment is deployed was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*DeploymentTarget `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// 	- If the value of success was false, an error code was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// 	- If the value of success was false, an error message was returned.
	//
	// 	- If the value of success was true, a null value was returned.
	//
	// example:
	//
	// ""
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The value was fixed to 200.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBC799F0-AS7S-1D30-8A4F-882ED4DD****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDeploymentTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentTargetsResponseBody) GetData() []*DeploymentTarget {
	return s.Data
}

func (s *ListDeploymentTargetsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDeploymentTargetsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDeploymentTargetsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListDeploymentTargetsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentTargetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentTargetsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeploymentTargetsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListDeploymentTargetsResponseBody) SetData(v []*DeploymentTarget) *ListDeploymentTargetsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetErrorCode(v string) *ListDeploymentTargetsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetErrorMessage(v string) *ListDeploymentTargetsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetHttpCode(v int32) *ListDeploymentTargetsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetPageIndex(v int32) *ListDeploymentTargetsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetPageSize(v int32) *ListDeploymentTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetRequestId(v string) *ListDeploymentTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetSuccess(v bool) *ListDeploymentTargetsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) SetTotalSize(v int32) *ListDeploymentTargetsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListDeploymentTargetsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
