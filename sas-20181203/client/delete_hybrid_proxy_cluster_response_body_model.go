// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridProxyClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHybridProxyClusterResponseBody
	GetRequestId() *string
}

type DeleteHybridProxyClusterResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 414EC213-AD2D-56C3-B140-108773B24405
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHybridProxyClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridProxyClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHybridProxyClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHybridProxyClusterResponseBody) SetRequestId(v string) *DeleteHybridProxyClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHybridProxyClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
