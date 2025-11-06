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
	// E2598CAF-DBFE-494E-95EF-B42A33C178AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e893148f-6343-4ae1-9eba-6e2a4116e142
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
