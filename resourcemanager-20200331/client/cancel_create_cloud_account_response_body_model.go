// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelCreateCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelCreateCloudAccountResponseBody
	GetRequestId() *string
}

type CancelCreateCloudAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelCreateCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelCreateCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *CancelCreateCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelCreateCloudAccountResponseBody) SetRequestId(v string) *CancelCreateCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelCreateCloudAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
