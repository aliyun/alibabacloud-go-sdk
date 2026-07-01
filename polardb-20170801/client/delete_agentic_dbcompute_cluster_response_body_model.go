// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDBComputeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAgenticDBComputeClusterResponseBody
	GetRequestId() *string
}

type DeleteAgenticDBComputeClusterResponseBody struct {
	// example:
	//
	// D4E5F6A7-B8C9-0123-DEFA-456789012DEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAgenticDBComputeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDBComputeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDBComputeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDBComputeClusterResponseBody) SetRequestId(v string) *DeleteAgenticDBComputeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDBComputeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
