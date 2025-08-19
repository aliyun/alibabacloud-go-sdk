// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeletionType(v string) *DeleteAccountResponseBody
	GetDeletionType() *string
	SetRequestId(v string) *DeleteAccountResponseBody
	GetRequestId() *string
}

type DeleteAccountResponseBody struct {
	// The type of the deletion. Valid values:
	//
	// 	- 0: direct deletion. If the member does not have pay-as-you-go resources that are purchased within the previous 30 days, the system directly deletes the member.
	//
	// 	- 1: deletion with a silence period. If the member has pay-as-you-go resources that are purchased within the previous 30 days, the member enters a silence period. The system starts to delete the member until the silence period ends. For more information about the silence period, see [What is the silence period for member deletion?](https://help.aliyun.com/document_detail/446079.html)
	//
	// example:
	//
	// 0
	DeletionType *string `json:"DeletionType,omitempty" xml:"DeletionType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 009429F8-C1C0-5872-B674-A6C2333B9647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccountResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccountResponseBody) GetDeletionType() *string {
	return s.DeletionType
}

func (s *DeleteAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccountResponseBody) SetDeletionType(v string) *DeleteAccountResponseBody {
	s.DeletionType = &v
	return s
}

func (s *DeleteAccountResponseBody) SetRequestId(v string) *DeleteAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
