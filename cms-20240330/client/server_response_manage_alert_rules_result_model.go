// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServerResponseManageAlertRulesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ServerResponseManageAlertRulesResult
	GetCode() *string
	SetData(v *ManageAlertRulesResult) *ServerResponseManageAlertRulesResult
	GetData() *ManageAlertRulesResult
	SetMessage(v string) *ServerResponseManageAlertRulesResult
	GetMessage() *string
	SetNextToken(v string) *ServerResponseManageAlertRulesResult
	GetNextToken() *string
	SetPageNumber(v int32) *ServerResponseManageAlertRulesResult
	GetPageNumber() *int32
	SetPageSize(v int32) *ServerResponseManageAlertRulesResult
	GetPageSize() *int32
	SetRequestId(v string) *ServerResponseManageAlertRulesResult
	GetRequestId() *string
	SetSuccess(v bool) *ServerResponseManageAlertRulesResult
	GetSuccess() *bool
	SetTotal(v int32) *ServerResponseManageAlertRulesResult
	GetTotal() *int32
}

type ServerResponseManageAlertRulesResult struct {
	// The status code of the API response. A value of 200 indicates a successful request. Other values indicate a request exception.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data returned by the API, which contains the execution result details of alert rule management operations such as batch enabling or disabling.
	//
	// example:
	//
	// {"updatedUuidList":["a1b2c3d4-e5f6-7890-abcd-ef1234567890"]}
	Data *ManageAlertRulesResult `json:"data,omitempty" xml:"data,omitempty"`
	// The error description returned when the request fails. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// The specified alert rule UUID does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The cursor for the next page in keyset-based pagination. A value of null indicates that no more data is available. Use this value to retrieve the next page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number of the returned data.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of records per page in the returned data.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The unique ID of the API request, which is used for troubleshooting and correlating server-side logs.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of alert rules actually affected by the operation.
	//
	// example:
	//
	// 8
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ServerResponseManageAlertRulesResult) String() string {
	return dara.Prettify(s)
}

func (s ServerResponseManageAlertRulesResult) GoString() string {
	return s.String()
}

func (s *ServerResponseManageAlertRulesResult) GetCode() *string {
	return s.Code
}

func (s *ServerResponseManageAlertRulesResult) GetData() *ManageAlertRulesResult {
	return s.Data
}

func (s *ServerResponseManageAlertRulesResult) GetMessage() *string {
	return s.Message
}

func (s *ServerResponseManageAlertRulesResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ServerResponseManageAlertRulesResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ServerResponseManageAlertRulesResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ServerResponseManageAlertRulesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ServerResponseManageAlertRulesResult) GetSuccess() *bool {
	return s.Success
}

func (s *ServerResponseManageAlertRulesResult) GetTotal() *int32 {
	return s.Total
}

func (s *ServerResponseManageAlertRulesResult) SetCode(v string) *ServerResponseManageAlertRulesResult {
	s.Code = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetData(v *ManageAlertRulesResult) *ServerResponseManageAlertRulesResult {
	s.Data = v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetMessage(v string) *ServerResponseManageAlertRulesResult {
	s.Message = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetNextToken(v string) *ServerResponseManageAlertRulesResult {
	s.NextToken = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetPageNumber(v int32) *ServerResponseManageAlertRulesResult {
	s.PageNumber = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetPageSize(v int32) *ServerResponseManageAlertRulesResult {
	s.PageSize = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetRequestId(v string) *ServerResponseManageAlertRulesResult {
	s.RequestId = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetSuccess(v bool) *ServerResponseManageAlertRulesResult {
	s.Success = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) SetTotal(v int32) *ServerResponseManageAlertRulesResult {
	s.Total = &v
	return s
}

func (s *ServerResponseManageAlertRulesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
