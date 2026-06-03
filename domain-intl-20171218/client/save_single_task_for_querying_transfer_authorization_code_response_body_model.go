// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) SetRequestId(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) SetTaskNo(v string) *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForQueryingTransferAuthorizationCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
