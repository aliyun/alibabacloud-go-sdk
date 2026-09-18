// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServerResponseQueryAlertRulesResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ServerResponseQueryAlertRulesResult
	GetCode() *string
	SetData(v *QueryAlertRulesResult) *ServerResponseQueryAlertRulesResult
	GetData() *QueryAlertRulesResult
	SetMessage(v string) *ServerResponseQueryAlertRulesResult
	GetMessage() *string
	SetNextToken(v string) *ServerResponseQueryAlertRulesResult
	GetNextToken() *string
	SetPageNumber(v int32) *ServerResponseQueryAlertRulesResult
	GetPageNumber() *int32
	SetPageSize(v int32) *ServerResponseQueryAlertRulesResult
	GetPageSize() *int32
	SetRequestId(v string) *ServerResponseQueryAlertRulesResult
	GetRequestId() *string
	SetSuccess(v bool) *ServerResponseQueryAlertRulesResult
	GetSuccess() *bool
	SetTotal(v int32) *ServerResponseQueryAlertRulesResult
	GetTotal() *int32
}

type ServerResponseQueryAlertRulesResult struct {
	// The status code of the API response. A value of 200 indicates a successful request. Other values indicate an exception.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The business data returned by the API, including the list of alert rule query results and pagination information.
	//
	// example:
	//
	// {"alertRules":[],"totalCount":0}
	Data *QueryAlertRulesResult `json:"data,omitempty" xml:"data,omitempty"`
	// The detailed error description returned when the request fails. This parameter is empty when the request succeeds.
	//
	// example:
	//
	// The request is invalid.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The token for the next page. A value of null indicates that no more pages are available.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The page number of the current response.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of records per page in the current response.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The unique ID of the API request, used for troubleshooting and server-side log tracing.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of alert rules that match the query conditions.
	//
	// example:
	//
	// 1
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ServerResponseQueryAlertRulesResult) String() string {
	return dara.Prettify(s)
}

func (s ServerResponseQueryAlertRulesResult) GoString() string {
	return s.String()
}

func (s *ServerResponseQueryAlertRulesResult) GetCode() *string {
	return s.Code
}

func (s *ServerResponseQueryAlertRulesResult) GetData() *QueryAlertRulesResult {
	return s.Data
}

func (s *ServerResponseQueryAlertRulesResult) GetMessage() *string {
	return s.Message
}

func (s *ServerResponseQueryAlertRulesResult) GetNextToken() *string {
	return s.NextToken
}

func (s *ServerResponseQueryAlertRulesResult) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ServerResponseQueryAlertRulesResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ServerResponseQueryAlertRulesResult) GetRequestId() *string {
	return s.RequestId
}

func (s *ServerResponseQueryAlertRulesResult) GetSuccess() *bool {
	return s.Success
}

func (s *ServerResponseQueryAlertRulesResult) GetTotal() *int32 {
	return s.Total
}

func (s *ServerResponseQueryAlertRulesResult) SetCode(v string) *ServerResponseQueryAlertRulesResult {
	s.Code = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetData(v *QueryAlertRulesResult) *ServerResponseQueryAlertRulesResult {
	s.Data = v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetMessage(v string) *ServerResponseQueryAlertRulesResult {
	s.Message = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetNextToken(v string) *ServerResponseQueryAlertRulesResult {
	s.NextToken = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetPageNumber(v int32) *ServerResponseQueryAlertRulesResult {
	s.PageNumber = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetPageSize(v int32) *ServerResponseQueryAlertRulesResult {
	s.PageSize = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetRequestId(v string) *ServerResponseQueryAlertRulesResult {
	s.RequestId = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetSuccess(v bool) *ServerResponseQueryAlertRulesResult {
	s.Success = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) SetTotal(v int32) *ServerResponseQueryAlertRulesResult {
	s.Total = &v
	return s
}

func (s *ServerResponseQueryAlertRulesResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
