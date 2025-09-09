// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAssetOperationTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RenewAssetOperationTokenResponseBody
	GetRequestId() *string
}

type RenewAssetOperationTokenResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RenewAssetOperationTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAssetOperationTokenResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAssetOperationTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAssetOperationTokenResponseBody) SetRequestId(v string) *RenewAssetOperationTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAssetOperationTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
