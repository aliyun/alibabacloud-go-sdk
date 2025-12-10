// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveCloudAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveCloudAccountResponseBody
	GetRequestId() *string
}

type RemoveCloudAccountResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveCloudAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveCloudAccountResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveCloudAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveCloudAccountResponseBody) SetRequestId(v string) *RemoveCloudAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveCloudAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
