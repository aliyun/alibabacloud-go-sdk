// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForAssociatingEnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForAssociatingEnsResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForAssociatingEnsResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForAssociatingEnsResponseBody struct {
	// example:
	//
	// A9C35C47-3366-482E-B872-8C9EA4733FE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 561bc091-f16f-4132-8d63-f15edce45731
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForAssociatingEnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForAssociatingEnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) SetRequestId(v string) *SaveSingleTaskForAssociatingEnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) SetTaskNo(v string) *SaveSingleTaskForAssociatingEnsResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForAssociatingEnsResponseBody) Validate() error {
	return dara.Validate(s)
}
