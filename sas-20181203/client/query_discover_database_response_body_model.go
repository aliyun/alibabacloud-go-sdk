// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDiscoverDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryDiscoverDatabaseResponseBody
	GetRequestId() *string
	SetTaskProgress(v int32) *QueryDiscoverDatabaseResponseBody
	GetTaskProgress() *int32
}

type QueryDiscoverDatabaseResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The progress of the database scan task in percentage.
	//
	// example:
	//
	// 90
	TaskProgress *int32 `json:"TaskProgress,omitempty" xml:"TaskProgress,omitempty"`
}

func (s QueryDiscoverDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDiscoverDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDiscoverDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDiscoverDatabaseResponseBody) GetTaskProgress() *int32 {
	return s.TaskProgress
}

func (s *QueryDiscoverDatabaseResponseBody) SetRequestId(v string) *QueryDiscoverDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDiscoverDatabaseResponseBody) SetTaskProgress(v int32) *QueryDiscoverDatabaseResponseBody {
	s.TaskProgress = &v
	return s
}

func (s *QueryDiscoverDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
