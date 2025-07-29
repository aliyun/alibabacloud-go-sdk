// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunClusterCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckId(v string) *RunClusterCheckResponseBody
	GetCheckId() *string
	SetRequestId(v string) *RunClusterCheckResponseBody
	GetRequestId() *string
}

type RunClusterCheckResponseBody struct {
	// The ID of the cluster check task.
	//
	// example:
	//
	// 1697100584236600453-ce0da5a1d627e4e9e9f96cae8ad07****-clustercheck-lboto
	CheckId *string `json:"check_id,omitempty" xml:"check_id,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F04DF81D-5C12-1524-B36A-86E02526****
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty"`
}

func (s RunClusterCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunClusterCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RunClusterCheckResponseBody) GetCheckId() *string {
	return s.CheckId
}

func (s *RunClusterCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunClusterCheckResponseBody) SetCheckId(v string) *RunClusterCheckResponseBody {
	s.CheckId = &v
	return s
}

func (s *RunClusterCheckResponseBody) SetRequestId(v string) *RunClusterCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunClusterCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
