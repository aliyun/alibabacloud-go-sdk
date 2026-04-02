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
	Code       *string                `json:"code,omitempty" xml:"code,omitempty"`
	Data       *QueryAlertRulesResult `json:"data,omitempty" xml:"data,omitempty"`
	Message    *string                `json:"message,omitempty" xml:"message,omitempty"`
	NextToken  *string                `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	PageNumber *int32                 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize   *int32                 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	RequestId  *string                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success    *bool                  `json:"success,omitempty" xml:"success,omitempty"`
	Total      *int32                 `json:"total,omitempty" xml:"total,omitempty"`
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
