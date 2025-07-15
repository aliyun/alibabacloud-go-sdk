// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartAndroidInstanceResponseBody
	GetRequestId() *string
}

type StartAndroidInstanceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 17C731AB-AAEE-5844-A352-D8D0352D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartAndroidInstanceResponseBody) SetRequestId(v string) *StartAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartAndroidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
