// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSaveArtExtensionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForSaveArtExtensionResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForSaveArtExtensionResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForSaveArtExtensionResponseBody struct {
	// example:
	//
	// E2598CAF-DBFE-494E-95EF-B42A33C178AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e893148f-6343-4ae1-9eba-6e2a4116e141
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSaveArtExtensionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSaveArtExtensionResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) SetRequestId(v string) *SaveSingleTaskForSaveArtExtensionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) SetTaskNo(v string) *SaveSingleTaskForSaveArtExtensionResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForSaveArtExtensionResponseBody) Validate() error {
	return dara.Validate(s)
}
