// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreloadCachesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PreloadCachesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *PreloadCachesResponseBody
	GetTaskId() *string
}

type PreloadCachesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9732E117-8A37-49FD-A36F-ABBB87556CA7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The prefetch task ID.
	//
	// example:
	//
	// 16401427840
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PreloadCachesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreloadCachesResponseBody) GoString() string {
	return s.String()
}

func (s *PreloadCachesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreloadCachesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *PreloadCachesResponseBody) SetRequestId(v string) *PreloadCachesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreloadCachesResponseBody) SetTaskId(v string) *PreloadCachesResponseBody {
	s.TaskId = &v
	return s
}

func (s *PreloadCachesResponseBody) Validate() error {
	return dara.Validate(s)
}
