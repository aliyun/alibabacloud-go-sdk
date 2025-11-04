// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDNAJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobList(v []*QueryDNAJobListResponseBodyJobList) *QueryDNAJobListResponseBody
	GetJobList() []*QueryDNAJobListResponseBodyJobList
	SetRequestId(v string) *QueryDNAJobListResponseBody
	GetRequestId() *string
}

type QueryDNAJobListResponseBody struct {
	// The queried media fingerprint analysis jobs.
	JobList []*QueryDNAJobListResponseBodyJobList `json:"JobList,omitempty" xml:"JobList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 25818875-5F78-4A13-BEF6-D7393642CA58
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryDNAJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryDNAJobListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryDNAJobListResponseBody) GetJobList() []*QueryDNAJobListResponseBodyJobList {
	return s.JobList
}

func (s *QueryDNAJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryDNAJobListResponseBody) SetJobList(v []*QueryDNAJobListResponseBodyJobList) *QueryDNAJobListResponseBody {
	s.JobList = v
	return s
}

func (s *QueryDNAJobListResponseBody) SetRequestId(v string) *QueryDNAJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryDNAJobListResponseBody) Validate() error {
	if s.JobList != nil {
		for _, item := range s.JobList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryDNAJobListResponseBodyJobList struct {
	// The response code.
	//
	// example:
	//
	// "InvalidParameter.ResourceNotFound"
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configurations of the media fingerprint analysis job.
	//
	// example:
	//
	// {"SaveType": "save","MediaType"":"video"}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// The time when the job was created.
	//
	// example:
	//
	// 2022-12-28T03:21:37Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the media fingerprint library.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	DBId *string `json:"DBId,omitempty" xml:"DBId,omitempty"`
	// The URL of the media fingerprint analysis result.
	//
	// example:
	//
	// http://test_bucket.oss-cn-shanghai.aliyuncs.com/fingerprint/video/search_result/5/5.txt
	DNAResult *string `json:"DNAResult,omitempty" xml:"DNAResult,omitempty"`
	// The time when the job was complete.
	//
	// example:
	//
	// 2022-12-28T03:21:44Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// 88c6ca184c0e47098a5b665e2a12****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The details of the input file.
	Input *QueryDNAJobListResponseBodyJobListInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// "The resource operated \\"a887d0b***d805ef6f7f6786302\\" cannot be found"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The primary key of the video. You must make sure that each primary key is unique.
	//
	// example:
	//
	// 3ca84a39a9024f19853b21be9cf9****
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The job state. Valid values:
	//
	// 	- **Queuing**: The job is waiting in the queue.
	//
	// 	- **Analysing**: The job is in progress.
	//
	// 	- **Success**: The job is successful.
	//
	// 	- **Fail**: The job failed.
	//
	// example:
	//
	// Queuing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The user-defined data.
	//
	// example:
	//
	// testdna
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s QueryDNAJobListResponseBodyJobList) String() string {
	return dara.Prettify(s)
}

func (s QueryDNAJobListResponseBodyJobList) GoString() string {
	return s.String()
}

func (s *QueryDNAJobListResponseBodyJobList) GetCode() *string {
	return s.Code
}

func (s *QueryDNAJobListResponseBodyJobList) GetConfig() *string {
	return s.Config
}

func (s *QueryDNAJobListResponseBodyJobList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *QueryDNAJobListResponseBodyJobList) GetDBId() *string {
	return s.DBId
}

func (s *QueryDNAJobListResponseBodyJobList) GetDNAResult() *string {
	return s.DNAResult
}

func (s *QueryDNAJobListResponseBodyJobList) GetFinishTime() *string {
	return s.FinishTime
}

func (s *QueryDNAJobListResponseBodyJobList) GetId() *string {
	return s.Id
}

func (s *QueryDNAJobListResponseBodyJobList) GetInput() *QueryDNAJobListResponseBodyJobListInput {
	return s.Input
}

func (s *QueryDNAJobListResponseBodyJobList) GetMessage() *string {
	return s.Message
}

func (s *QueryDNAJobListResponseBodyJobList) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *QueryDNAJobListResponseBodyJobList) GetStatus() *string {
	return s.Status
}

func (s *QueryDNAJobListResponseBodyJobList) GetUserData() *string {
	return s.UserData
}

func (s *QueryDNAJobListResponseBodyJobList) SetCode(v string) *QueryDNAJobListResponseBodyJobList {
	s.Code = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetConfig(v string) *QueryDNAJobListResponseBodyJobList {
	s.Config = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetCreationTime(v string) *QueryDNAJobListResponseBodyJobList {
	s.CreationTime = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetDBId(v string) *QueryDNAJobListResponseBodyJobList {
	s.DBId = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetDNAResult(v string) *QueryDNAJobListResponseBodyJobList {
	s.DNAResult = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetFinishTime(v string) *QueryDNAJobListResponseBodyJobList {
	s.FinishTime = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetId(v string) *QueryDNAJobListResponseBodyJobList {
	s.Id = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetInput(v *QueryDNAJobListResponseBodyJobListInput) *QueryDNAJobListResponseBodyJobList {
	s.Input = v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetMessage(v string) *QueryDNAJobListResponseBodyJobList {
	s.Message = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetPrimaryKey(v string) *QueryDNAJobListResponseBodyJobList {
	s.PrimaryKey = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetStatus(v string) *QueryDNAJobListResponseBodyJobList {
	s.Status = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) SetUserData(v string) *QueryDNAJobListResponseBodyJobList {
	s.UserData = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobList) Validate() error {
	if s.Input != nil {
		if err := s.Input.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryDNAJobListResponseBodyJobListInput struct {
	// The input file. The file can be an OSS object or a media asset. The path of an OSS object can be in one of the following formats:
	//
	// 1\\. oss://bucket/object
	//
	// 2\\. http(s)://bucket.oss-[regionId].aliyuncs.com/object
	//
	// In the preceding paths, bucket indicates an OSS bucket that resides in the same region as the current project, and object indicates the path of the object in the bucket.
	//
	// example:
	//
	// 1b1b9cd148034739af413150fded****
	Media *string `json:"Media,omitempty" xml:"Media,omitempty"`
	// The type of the input file. Valid values:
	//
	// 1.  OSS: Object Storage Service (OSS) object.
	//
	// 2.  Media: media asset.
	//
	// example:
	//
	// Media
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryDNAJobListResponseBodyJobListInput) String() string {
	return dara.Prettify(s)
}

func (s QueryDNAJobListResponseBodyJobListInput) GoString() string {
	return s.String()
}

func (s *QueryDNAJobListResponseBodyJobListInput) GetMedia() *string {
	return s.Media
}

func (s *QueryDNAJobListResponseBodyJobListInput) GetType() *string {
	return s.Type
}

func (s *QueryDNAJobListResponseBodyJobListInput) SetMedia(v string) *QueryDNAJobListResponseBodyJobListInput {
	s.Media = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobListInput) SetType(v string) *QueryDNAJobListResponseBodyJobListInput {
	s.Type = &v
	return s
}

func (s *QueryDNAJobListResponseBodyJobListInput) Validate() error {
	return dara.Validate(s)
}
