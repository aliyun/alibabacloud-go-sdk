// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoShowListTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoShowListTasks(v string) *DescribeAutoShowListTasksResponseBody
	GetAutoShowListTasks() *string
	SetRequestId(v string) *DescribeAutoShowListTasksResponseBody
	GetRequestId() *string
}

type DescribeAutoShowListTasksResponseBody struct {
	// The information about the scheduled tasks. The following fields are included:
	//
	// 	- Status: the status of the scheduled task. Valid values: 0 and 1. A value of 0 indicates that the scheduled task is paused. A value of 1 indicates that the scheduled task is started.
	//
	// 	- LiveTemplate: the transcoding templates.
	//
	// 	- TranscodeConfig: the transcoding configuration for the source URL.
	//
	// 	- CasterId: the ID of the production studio.
	//
	// example:
	//
	// {"Status":0,"TranscodeConfig":{"CasterTemplate":"lp_hd", "LiveTemplate":["lhd","lsd"]}, "CasterId":"cce04ef3-2226-4865-8704-f84b8375****"}
	AutoShowListTasks *string `json:"AutoShowListTasks,omitempty" xml:"AutoShowListTasks,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAutoShowListTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoShowListTasksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoShowListTasksResponseBody) GetAutoShowListTasks() *string {
	return s.AutoShowListTasks
}

func (s *DescribeAutoShowListTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoShowListTasksResponseBody) SetAutoShowListTasks(v string) *DescribeAutoShowListTasksResponseBody {
	s.AutoShowListTasks = &v
	return s
}

func (s *DescribeAutoShowListTasksResponseBody) SetRequestId(v string) *DescribeAutoShowListTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoShowListTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
