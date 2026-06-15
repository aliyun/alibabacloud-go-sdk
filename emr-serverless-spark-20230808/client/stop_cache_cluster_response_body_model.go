// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCacheClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopCacheClusterResponseBody
	GetRequestId() *string
}

type StopCacheClusterResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StopCacheClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopCacheClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StopCacheClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopCacheClusterResponseBody) SetRequestId(v string) *StopCacheClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopCacheClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
