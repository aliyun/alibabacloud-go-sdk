// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadInstanceKeyPairResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UploadInstanceKeyPairResponseBody
	GetRequestId() *string
}

type UploadInstanceKeyPairResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UploadInstanceKeyPairResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadInstanceKeyPairResponseBody) GoString() string {
	return s.String()
}

func (s *UploadInstanceKeyPairResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadInstanceKeyPairResponseBody) SetRequestId(v string) *UploadInstanceKeyPairResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadInstanceKeyPairResponseBody) Validate() error {
	return dara.Validate(s)
}
