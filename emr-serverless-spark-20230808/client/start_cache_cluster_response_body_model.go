// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCacheClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StartCacheClusterResponseBody
	GetRequestId() *string
}

type StartCacheClusterResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s StartCacheClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartCacheClusterResponseBody) GoString() string {
	return s.String()
}

func (s *StartCacheClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartCacheClusterResponseBody) SetRequestId(v string) *StartCacheClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartCacheClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
