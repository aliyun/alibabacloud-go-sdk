// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFileUploadSignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallFrom(v string) *DescribeFileUploadSignatureRequest
	GetCallFrom() *string
	SetDmsUnit(v string) *DescribeFileUploadSignatureRequest
	GetDmsUnit() *string
	SetWorkspaceId(v string) *DescribeFileUploadSignatureRequest
	GetWorkspaceId() *string
}

type DescribeFileUploadSignatureRequest struct {
	// The parameter used only by the frontend.
	//
	// example:
	//
	// 仅前端使用
	CallFrom *string `json:"CallFrom,omitempty" xml:"CallFrom,omitempty"`
	// The current DMS unit.
	//
	// example:
	//
	// cn-hangzhou
	DmsUnit     *string `json:"DmsUnit,omitempty" xml:"DmsUnit,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DescribeFileUploadSignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFileUploadSignatureRequest) GoString() string {
	return s.String()
}

func (s *DescribeFileUploadSignatureRequest) GetCallFrom() *string {
	return s.CallFrom
}

func (s *DescribeFileUploadSignatureRequest) GetDmsUnit() *string {
	return s.DmsUnit
}

func (s *DescribeFileUploadSignatureRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *DescribeFileUploadSignatureRequest) SetCallFrom(v string) *DescribeFileUploadSignatureRequest {
	s.CallFrom = &v
	return s
}

func (s *DescribeFileUploadSignatureRequest) SetDmsUnit(v string) *DescribeFileUploadSignatureRequest {
	s.DmsUnit = &v
	return s
}

func (s *DescribeFileUploadSignatureRequest) SetWorkspaceId(v string) *DescribeFileUploadSignatureRequest {
	s.WorkspaceId = &v
	return s
}

func (s *DescribeFileUploadSignatureRequest) Validate() error {
	return dara.Validate(s)
}
