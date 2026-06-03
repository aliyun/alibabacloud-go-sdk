// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDisassociatingEnsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForDisassociatingEnsResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForDisassociatingEnsResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForDisassociatingEnsResponseBody struct {
	// example:
	//
	// A9C35C47-3366-482E-B872-8C9EA4733FE9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 561bc091-f16f-4132-8d63-f15edce45731
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDisassociatingEnsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDisassociatingEnsResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) SetRequestId(v string) *SaveSingleTaskForDisassociatingEnsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) SetTaskNo(v string) *SaveSingleTaskForDisassociatingEnsResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForDisassociatingEnsResponseBody) Validate() error {
	return dara.Validate(s)
}
