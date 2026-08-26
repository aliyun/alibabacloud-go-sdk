// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingJobInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *GetEditingJobInfoResponseBody
	GetCasterId() *string
	SetEditingTasksInfo(v string) *GetEditingJobInfoResponseBody
	GetEditingTasksInfo() *string
	SetRequestId(v string) *GetEditingJobInfoResponseBody
	GetRequestId() *string
}

type GetEditingJobInfoResponseBody struct {
	// The production studio ID.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The video clip task information. This includes:
	//
	// - **OutputVodId**: The ID of the output video-on-demand file.
	//
	// - **TaskStatus**: The status of the video clip task. (-1: failed. 0: task initialized. 1: clipping in progress. 2: uploading. 3: task succeeded.)
	//
	// - **StorageLocation**: The video-on-demand storage address.
	//
	// - **FileName**: The name of the clipped file.
	//
	// - **ShowId**: The show ID.
	//
	// example:
	//
	// "EditingTasksInfo": {     "OutputVodId": "3e34733b40b9a96ccf5c1ff6f69****",     "TaskStatus": 1,     "StorageInfo": {       "StorageLocation": "***bucket***",       "FileName": "EditFile****"     },     "ShowId": "42200b81-b761-4c10-842a-a0726d97****"   },
	EditingTasksInfo *string `json:"EditingTasksInfo,omitempty" xml:"EditingTasksInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingJobInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEditingJobInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingJobInfoResponseBody) GetCasterId() *string {
	return s.CasterId
}

func (s *GetEditingJobInfoResponseBody) GetEditingTasksInfo() *string {
	return s.EditingTasksInfo
}

func (s *GetEditingJobInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEditingJobInfoResponseBody) SetCasterId(v string) *GetEditingJobInfoResponseBody {
	s.CasterId = &v
	return s
}

func (s *GetEditingJobInfoResponseBody) SetEditingTasksInfo(v string) *GetEditingJobInfoResponseBody {
	s.EditingTasksInfo = &v
	return s
}

func (s *GetEditingJobInfoResponseBody) SetRequestId(v string) *GetEditingJobInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingJobInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
