// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelKeyDeletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelKeyDeletionResponseBody
	GetRequestId() *string
}

type CancelKeyDeletionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3da5b8cc-8107-40ac-a170-793cd181d7b7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelKeyDeletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelKeyDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *CancelKeyDeletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelKeyDeletionResponseBody) SetRequestId(v string) *CancelKeyDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelKeyDeletionResponseBody) Validate() error {
	return dara.Validate(s)
}
