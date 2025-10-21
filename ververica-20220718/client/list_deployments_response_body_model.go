// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeploymentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*Deployment) *ListDeploymentsResponseBody
	GetData() []*Deployment
	SetErrorCode(v string) *ListDeploymentsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDeploymentsResponseBody
	GetErrorMessage() *string
	SetHttpCode(v int32) *ListDeploymentsResponseBody
	GetHttpCode() *int32
	SetPageIndex(v int32) *ListDeploymentsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ListDeploymentsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDeploymentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDeploymentsResponseBody
	GetSuccess() *bool
	SetTotalSize(v int32) *ListDeploymentsResponseBody
	GetTotalSize() *int32
}

type ListDeploymentsResponseBody struct {
	// 	- If the value of success was true, the list of all deployments was returned.
	//
	// 	- If the value of success was false, a null value was returned.
	Data []*Deployment `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ListDeploymentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDeploymentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeploymentsResponseBody) GetData() []*Deployment {
	return s.Data
}

func (s *ListDeploymentsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDeploymentsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDeploymentsResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListDeploymentsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ListDeploymentsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDeploymentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDeploymentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDeploymentsResponseBody) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListDeploymentsResponseBody) SetData(v []*Deployment) *ListDeploymentsResponseBody {
	s.Data = v
	return s
}

func (s *ListDeploymentsResponseBody) SetErrorCode(v string) *ListDeploymentsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetErrorMessage(v string) *ListDeploymentsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetHttpCode(v int32) *ListDeploymentsResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetPageIndex(v int32) *ListDeploymentsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetPageSize(v int32) *ListDeploymentsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetRequestId(v string) *ListDeploymentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetSuccess(v bool) *ListDeploymentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDeploymentsResponseBody) SetTotalSize(v int32) *ListDeploymentsResponseBody {
	s.TotalSize = &v
	return s
}

func (s *ListDeploymentsResponseBody) Validate() error {
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
