// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterInspectConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateClusterInspectConfigResponseBody
	GetRequestId() *string
}

type UpdateClusterInspectConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 49511F2D-D56A-5C24-B9AE-C8491E09B***
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateClusterInspectConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterInspectConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateClusterInspectConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateClusterInspectConfigResponseBody) SetRequestId(v string) *UpdateClusterInspectConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateClusterInspectConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
