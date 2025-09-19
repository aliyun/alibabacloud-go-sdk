// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCheckConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeCheckConfigResponseBody
	GetRequestId() *string
}

type ChangeCheckConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6EBB8614-746D-555D-AB69-C801AEC7DCE0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeCheckConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ChangeCheckConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeCheckConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeCheckConfigResponseBody) SetRequestId(v string) *ChangeCheckConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeCheckConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
