// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchIndexFileMetaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *BatchIndexFileMetaRequest
	GetDatasetName() *string
	SetFiles(v []*InputFile) *BatchIndexFileMetaRequest
	GetFiles() []*InputFile
	SetNotification(v *Notification) *BatchIndexFileMetaRequest
	GetNotification() *Notification
	SetProjectName(v string) *BatchIndexFileMetaRequest
	GetProjectName() *string
	SetUserData(v string) *BatchIndexFileMetaRequest
	GetUserData() *string
}

type BatchIndexFileMetaRequest struct {
	// The name of the dataset.[](~~478160~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The objects in Object Storage Service (OSS). Specify OSS objects by using a JSON array. You can specify up to 100 objects in an array.
	//
	// This parameter is required.
	Files []*InputFile `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s BatchIndexFileMetaRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchIndexFileMetaRequest) GoString() string {
	return s.String()
}

func (s *BatchIndexFileMetaRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *BatchIndexFileMetaRequest) GetFiles() []*InputFile {
	return s.Files
}

func (s *BatchIndexFileMetaRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *BatchIndexFileMetaRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *BatchIndexFileMetaRequest) GetUserData() *string {
	return s.UserData
}

func (s *BatchIndexFileMetaRequest) SetDatasetName(v string) *BatchIndexFileMetaRequest {
	s.DatasetName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetFiles(v []*InputFile) *BatchIndexFileMetaRequest {
	s.Files = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetNotification(v *Notification) *BatchIndexFileMetaRequest {
	s.Notification = v
	return s
}

func (s *BatchIndexFileMetaRequest) SetProjectName(v string) *BatchIndexFileMetaRequest {
	s.ProjectName = &v
	return s
}

func (s *BatchIndexFileMetaRequest) SetUserData(v string) *BatchIndexFileMetaRequest {
	s.UserData = &v
	return s
}

func (s *BatchIndexFileMetaRequest) Validate() error {
	return dara.Validate(s)
}
