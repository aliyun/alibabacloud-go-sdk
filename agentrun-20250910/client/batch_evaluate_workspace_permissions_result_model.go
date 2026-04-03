// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchEvaluateWorkspacePermissionsResult interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchEvaluateWorkspacePermissionsResult
	GetCode() *string
	SetData(v *BatchEvaluateWorkspacePermissionsOutput) *BatchEvaluateWorkspacePermissionsResult
	GetData() *BatchEvaluateWorkspacePermissionsOutput
	SetRequestId(v string) *BatchEvaluateWorkspacePermissionsResult
	GetRequestId() *string
}

type BatchEvaluateWorkspacePermissionsResult struct {
	// OK 表示成功
	Code *string                                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *BatchEvaluateWorkspacePermissionsOutput `json:"data,omitempty" xml:"data,omitempty"`
	// 与响应头 x-funagent-request-id 对应
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s BatchEvaluateWorkspacePermissionsResult) String() string {
	return dara.Prettify(s)
}

func (s BatchEvaluateWorkspacePermissionsResult) GoString() string {
	return s.String()
}

func (s *BatchEvaluateWorkspacePermissionsResult) GetCode() *string {
	return s.Code
}

func (s *BatchEvaluateWorkspacePermissionsResult) GetData() *BatchEvaluateWorkspacePermissionsOutput {
	return s.Data
}

func (s *BatchEvaluateWorkspacePermissionsResult) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchEvaluateWorkspacePermissionsResult) SetCode(v string) *BatchEvaluateWorkspacePermissionsResult {
	s.Code = &v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsResult) SetData(v *BatchEvaluateWorkspacePermissionsOutput) *BatchEvaluateWorkspacePermissionsResult {
	s.Data = v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsResult) SetRequestId(v string) *BatchEvaluateWorkspacePermissionsResult {
	s.RequestId = &v
	return s
}

func (s *BatchEvaluateWorkspacePermissionsResult) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}
