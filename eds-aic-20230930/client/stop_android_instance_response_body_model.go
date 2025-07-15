// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAndroidInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAndroidInstanceResponseBody
	GetRequestId() *string
}

type StopAndroidInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// E38B41A8-8E00-5AE4-A957-6636ACB8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAndroidInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAndroidInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StopAndroidInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAndroidInstanceResponseBody) SetRequestId(v string) *StopAndroidInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAndroidInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
