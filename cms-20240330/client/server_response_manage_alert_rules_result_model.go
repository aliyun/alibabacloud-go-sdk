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
	// 响应码
	Code *string                 `json:"code,omitempty" xml:"code,omitempty"`
	Data *ManageAlertRulesResult `json:"data,omitempty" xml:"data,omitempty"`
	// 错误消息
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 分页 Token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// 页码
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// 每页大小
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求 ID
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 是否成功
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// 总数
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
