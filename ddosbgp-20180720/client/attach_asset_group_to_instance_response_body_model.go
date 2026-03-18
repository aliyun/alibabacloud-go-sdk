// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachAssetGroupToInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachAssetGroupToInstanceResponseBody
	GetRequestId() *string
}

type AttachAssetGroupToInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 52B49F64-5A36-5CE0-BD00-765792C26AA9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachAssetGroupToInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachAssetGroupToInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *AttachAssetGroupToInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachAssetGroupToInstanceResponseBody) SetRequestId(v string) *AttachAssetGroupToInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachAssetGroupToInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
