// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForTransferOutByAuthorizationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C560A803-B975-481D-A66B-A4395EA863A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) SetRequestId(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) SetTaskNo(v string) *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForTransferOutByAuthorizationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
