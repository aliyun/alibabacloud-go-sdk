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
	// The ID of the production studio.
	//
	// example:
	//
	// 53200b81-b761-4c10-842a-a0726d97****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The information about editing tasks. The following fields are returned for each editing task:
	//
	// 	- **OutputVodId**: the ID of the output video-on-demand (VOD) file.
	//
	// 	- **TaskStatus**: the status of the editing task. Valid values: -1, 0, 1, 2, and 3. A value of -1 indicates that the editing task fails. A value of 0 indicates that the editing task is being initialized. A value of 1 indicates that editing is in progress. A value of 2 indicates that the output VOD file is being uploaded. A value of 3 indicates that the editing task is successful.
	//
	// 	- **StorageLocation**: the storage location in ApsaraVideo VOD.
	//
	// 	- **FileName**: the name of the file that is edited.
	//
	// 	- **ShowId**: the ID of the episode.
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
