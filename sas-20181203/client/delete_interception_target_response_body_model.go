// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInterceptionTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteInterceptionTargetResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteInterceptionTargetResponseBody
	GetResult() *bool
}

type DeleteInterceptionTargetResponseBody struct {
	// The request ID. Alibaba Cloud generates a unique ID for each request. You can use the ID to troubleshoot issues.
	//
	// example:
	//
	// D81DD78E-E006-5C65-A171-C8CB09XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the network objects were deleted. Valid values:
	//
	// - **true**: The network objects were deleted.
	//
	// - **false**: The network objects failed to be deleted.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteInterceptionTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInterceptionTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInterceptionTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInterceptionTargetResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteInterceptionTargetResponseBody) SetRequestId(v string) *DeleteInterceptionTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInterceptionTargetResponseBody) SetResult(v bool) *DeleteInterceptionTargetResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteInterceptionTargetResponseBody) Validate() error {
	return dara.Validate(s)
}
