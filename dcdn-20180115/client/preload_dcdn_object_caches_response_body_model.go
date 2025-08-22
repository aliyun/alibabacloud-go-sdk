// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadDcdnObjectCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPreloadTaskId(v string) *PreloadDcdnObjectCachesResponseBody
	GetPreloadTaskId() *string
	SetRequestId(v string) *PreloadDcdnObjectCachesResponseBody
	GetRequestId() *string
}

type PreloadDcdnObjectCachesResponseBody struct {
	// The ID of the prefetch task. Multiple IDs are separated by commas (,).
	//
	// example:
	//
	// 95248880
	PreloadTaskId *string `json:"PreloadTaskId,omitempty" xml:"PreloadTaskId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E5BD4B50-7A02-493A-AE0B-97B9024B4135
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PreloadDcdnObjectCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadDcdnObjectCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadDcdnObjectCachesResponseBody) GetPreloadTaskId() *string {
	return s.PreloadTaskId
}

func (s *PreloadDcdnObjectCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadDcdnObjectCachesResponseBody) SetPreloadTaskId(v string) *PreloadDcdnObjectCachesResponseBody {
	s.PreloadTaskId = &v
	return s
}

func (s *PreloadDcdnObjectCachesResponseBody) SetRequestId(v string) *PreloadDcdnObjectCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadDcdnObjectCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
