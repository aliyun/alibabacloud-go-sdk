// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadVodObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreloadTaskId(v string) *PreloadVodObjectCachesResponseBody
	GetPreloadTaskId() *string
	SetRequestId(v string) *PreloadVodObjectCachesResponseBody
	GetRequestId() *string
}

type PreloadVodObjectCachesResponseBody struct {
	// The ID of the prefetch task. Separate multiple task IDs with commas (,).
	//
	// example:
	//
	// 9524****
	PreloadTaskId *string `json:"PreloadTaskId,omitempty" xml:"PreloadTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E5BD4B50-7A02-493A-*****-97B9024B4135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadVodObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadVodObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadVodObjectCachesResponseBody) GetPreloadTaskId() *string {
	return s.PreloadTaskId
}

func (s *PreloadVodObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadVodObjectCachesResponseBody) SetPreloadTaskId(v string) *PreloadVodObjectCachesResponseBody {
	s.PreloadTaskId = &v
	return s
}

func (s *PreloadVodObjectCachesResponseBody) SetRequestId(v string) *PreloadVodObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadVodObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
