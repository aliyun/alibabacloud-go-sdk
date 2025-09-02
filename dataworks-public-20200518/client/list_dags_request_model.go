// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpSeq(v int64) *ListDagsRequest
	GetOpSeq() *int64
	SetProjectEnv(v string) *ListDagsRequest
	GetProjectEnv() *string
}

type ListDagsRequest struct {
	// The sequence number that uniquely identifies the data backfill operation. You can call the [GetDag](https://help.aliyun.com/document_detail/189753.html) operation to query the sequence number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	OpSeq *int64 `json:"OpSeq,omitempty" xml:"OpSeq,omitempty"`
	// The environment of the workspace. Valid values: PROD and DEV. The value PROD indicates the production environment, and the value DEV indicates the development environment.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROD
	ProjectEnv *string `json:"ProjectEnv,omitempty" xml:"ProjectEnv,omitempty"`
}

func (s ListDagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDagsRequest) GoString() string {
	return s.String()
}

func (s *ListDagsRequest) GetOpSeq() *int64 {
	return s.OpSeq
}

func (s *ListDagsRequest) GetProjectEnv() *string {
	return s.ProjectEnv
}

func (s *ListDagsRequest) SetOpSeq(v int64) *ListDagsRequest {
	s.OpSeq = &v
	return s
}

func (s *ListDagsRequest) SetProjectEnv(v string) *ListDagsRequest {
	s.ProjectEnv = &v
	return s
}

func (s *ListDagsRequest) Validate() error {
	return dara.Validate(s)
}
