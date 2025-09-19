// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEmgVulSubmitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEmgVulSubmitResponseBody
	GetRequestId() *string
}

type ModifyEmgVulSubmitResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 52A3AEE6-114A-499D-8990-4BA9B27FE0AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEmgVulSubmitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEmgVulSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEmgVulSubmitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEmgVulSubmitResponseBody) SetRequestId(v string) *ModifyEmgVulSubmitResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEmgVulSubmitResponseBody) Validate() error {
	return dara.Validate(s)
}
