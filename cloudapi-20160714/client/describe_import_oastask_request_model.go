// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImportOASTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOperationId(v string) *DescribeImportOASTaskRequest
	GetOperationId() *string
	SetSecurityToken(v string) *DescribeImportOASTaskRequest
	GetSecurityToken() *string
}

type DescribeImportOASTaskRequest struct {
	// The ID of the asynchronous API import task that was generated during the import operation. This ID is used to query the execution status of the API import task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4c53ccf1d40c489686d1adf5c2644a7f
	OperationId   *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeImportOASTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImportOASTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeImportOASTaskRequest) GetOperationId() *string {
	return s.OperationId
}

func (s *DescribeImportOASTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeImportOASTaskRequest) SetOperationId(v string) *DescribeImportOASTaskRequest {
	s.OperationId = &v
	return s
}

func (s *DescribeImportOASTaskRequest) SetSecurityToken(v string) *DescribeImportOASTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeImportOASTaskRequest) Validate() error {
	return dara.Validate(s)
}
