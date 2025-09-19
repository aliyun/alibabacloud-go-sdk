// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreHcCheckWarningsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *IgnoreHcCheckWarningsResponseBody
	GetRequestId() *string
}

type IgnoreHcCheckWarningsResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B113119F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s IgnoreHcCheckWarningsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IgnoreHcCheckWarningsResponseBody) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IgnoreHcCheckWarningsResponseBody) SetRequestId(v string) *IgnoreHcCheckWarningsResponseBody {
	s.RequestId = &v
	return s
}

func (s *IgnoreHcCheckWarningsResponseBody) Validate() error {
	return dara.Validate(s)
}
