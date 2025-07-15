// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudPhoneNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCloudPhoneNodesResponseBody
	GetRequestId() *string
}

type DeleteCloudPhoneNodesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCloudPhoneNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudPhoneNodesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCloudPhoneNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCloudPhoneNodesResponseBody) SetRequestId(v string) *DeleteCloudPhoneNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCloudPhoneNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
