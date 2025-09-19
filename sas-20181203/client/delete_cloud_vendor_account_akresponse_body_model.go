// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudVendorAccountAKResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudVendorAccountAKResponseBody
	GetRequestId() *string
}

type DeleteCloudVendorAccountAKResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4539D402-F7A4-5915-9580-EC227BF*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudVendorAccountAKResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudVendorAccountAKResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudVendorAccountAKResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudVendorAccountAKResponseBody) SetRequestId(v string) *DeleteCloudVendorAccountAKResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudVendorAccountAKResponseBody) Validate() error {
	return dara.Validate(s)
}
