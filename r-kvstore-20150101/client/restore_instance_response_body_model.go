// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RestoreInstanceResponseBody
	GetRequestId() *string
}

type RestoreInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8D0C0AFC-E9CD-47A4-8395-5C31BF9B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RestoreInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestoreInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestoreInstanceResponseBody) SetRequestId(v string) *RestoreInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestoreInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
