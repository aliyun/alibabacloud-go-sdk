// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDettachAssetGroupToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DettachAssetGroupToInstanceResponseBody
	GetRequestId() *string
}

type DettachAssetGroupToInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E54BA258-9DE8-59BE-B7A8-DAD28E6E8DAF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DettachAssetGroupToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DettachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DettachAssetGroupToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DettachAssetGroupToInstanceResponseBody) SetRequestId(v string) *DettachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DettachAssetGroupToInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
