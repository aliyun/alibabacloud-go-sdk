// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveMfaDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemoveMfaDeviceResponseBody
	GetRequestId() *string
}

type RemoveMfaDeviceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FB550AAB-FB36-4A91-93F6-F4374AF65403
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemoveMfaDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveMfaDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveMfaDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveMfaDeviceResponseBody) SetRequestId(v string) *RemoveMfaDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveMfaDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
