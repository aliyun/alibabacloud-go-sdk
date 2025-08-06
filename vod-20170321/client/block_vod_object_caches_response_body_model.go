// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBlockVodObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBlockTaskId(v string) *BlockVodObjectCachesResponseBody
	GetBlockTaskId() *string
	SetRequestId(v string) *BlockVodObjectCachesResponseBody
	GetRequestId() *string
}

type BlockVodObjectCachesResponseBody struct {
	BlockTaskId *string `json:"BlockTaskId,omitempty" xml:"BlockTaskId,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BlockVodObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BlockVodObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *BlockVodObjectCachesResponseBody) GetBlockTaskId() *string {
	return s.BlockTaskId
}

func (s *BlockVodObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BlockVodObjectCachesResponseBody) SetBlockTaskId(v string) *BlockVodObjectCachesResponseBody {
	s.BlockTaskId = &v
	return s
}

func (s *BlockVodObjectCachesResponseBody) SetRequestId(v string) *BlockVodObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *BlockVodObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
