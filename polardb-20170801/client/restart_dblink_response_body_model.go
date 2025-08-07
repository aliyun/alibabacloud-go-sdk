// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDBLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *RestartDBLinkResponseBody
	GetDBClusterId() *string
	SetRequestId(v string) *RestartDBLinkResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RestartDBLinkResponseBody
	GetTaskId() *string
}

type RestartDBLinkResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// ec8c4723-eac5-4f12-becb-01ac08******
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RestartDBLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDBLinkResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDBLinkResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *RestartDBLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDBLinkResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RestartDBLinkResponseBody) SetDBClusterId(v string) *RestartDBLinkResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *RestartDBLinkResponseBody) SetRequestId(v string) *RestartDBLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDBLinkResponseBody) SetTaskId(v string) *RestartDBLinkResponseBody {
	s.TaskId = &v
	return s
}

func (s *RestartDBLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
