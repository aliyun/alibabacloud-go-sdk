// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSnapshotCallbackAuthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetSnapshotCallbackAuthResponseBody
	GetRequestId() *string
}

type SetSnapshotCallbackAuthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetSnapshotCallbackAuthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSnapshotCallbackAuthResponseBody) GoString() string {
	return s.String()
}

func (s *SetSnapshotCallbackAuthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSnapshotCallbackAuthResponseBody) SetRequestId(v string) *SetSnapshotCallbackAuthResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSnapshotCallbackAuthResponseBody) Validate() error {
	return dara.Validate(s)
}
