// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBundleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBundleId(v string) *CreateBundleResponseBody
	GetBundleId() *string
	SetRequestId(v string) *CreateBundleResponseBody
	GetRequestId() *string
}

type CreateBundleResponseBody struct {
	// The ID of the cloud computer template.
	//
	// example:
	//
	// b-cezrnfgecbich****
	BundleId *string `json:"BundleId,omitempty" xml:"BundleId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBundleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBundleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBundleResponseBody) GetBundleId() *string {
	return s.BundleId
}

func (s *CreateBundleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBundleResponseBody) SetBundleId(v string) *CreateBundleResponseBody {
	s.BundleId = &v
	return s
}

func (s *CreateBundleResponseBody) SetRequestId(v string) *CreateBundleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBundleResponseBody) Validate() error {
	return dara.Validate(s)
}
