// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAudioFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAudioFileResponseBody
	GetCode() *string
	SetData(v *GetAudioFileResponseBodyData) *GetAudioFileResponseBody
	GetData() *GetAudioFileResponseBodyData
	SetHttpStatusCode(v int32) *GetAudioFileResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAudioFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAudioFileResponseBody
	GetRequestId() *string
}

type GetAudioFileResponseBody struct {
	// example:
	//
	// OK
	Code *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetAudioFileResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EEE26562-D921-5CB2-AE49-E4C45A42D432
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAudioFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileResponseBody) GoString() string {
	return s.String()
}

func (s *GetAudioFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAudioFileResponseBody) GetData() *GetAudioFileResponseBodyData {
	return s.Data
}

func (s *GetAudioFileResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAudioFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAudioFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAudioFileResponseBody) SetCode(v string) *GetAudioFileResponseBody {
	s.Code = &v
	return s
}

func (s *GetAudioFileResponseBody) SetData(v *GetAudioFileResponseBodyData) *GetAudioFileResponseBody {
	s.Data = v
	return s
}

func (s *GetAudioFileResponseBody) SetHttpStatusCode(v int32) *GetAudioFileResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAudioFileResponseBody) SetMessage(v string) *GetAudioFileResponseBody {
	s.Message = &v
	return s
}

func (s *GetAudioFileResponseBody) SetRequestId(v string) *GetAudioFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAudioFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAudioFileResponseBodyData struct {
	// example:
	//
	// test-file.wav
	AudioFileName *string `json:"AudioFileName,omitempty" xml:"AudioFileName,omitempty"`
	// example:
	//
	// c1a06b46-302a-4c6e-928b-a43c0df485cf
	AudioResourceId *string `json:"AudioResourceId,omitempty" xml:"AudioResourceId,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// ccc-test/test-file.wav
	OssFileKey *string `json:"OssFileKey,omitempty" xml:"OssFileKey,omitempty"`
	// example:
	//
	// 2021-07-14 10:48:43.0
	UpdatedTime *string `json:"UpdatedTime,omitempty" xml:"UpdatedTime,omitempty"`
}

func (s GetAudioFileResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAudioFileResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAudioFileResponseBodyData) GetAudioFileName() *string {
	return s.AudioFileName
}

func (s *GetAudioFileResponseBodyData) GetAudioResourceId() *string {
	return s.AudioResourceId
}

func (s *GetAudioFileResponseBodyData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetAudioFileResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetAudioFileResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAudioFileResponseBodyData) GetOssFileKey() *string {
	return s.OssFileKey
}

func (s *GetAudioFileResponseBodyData) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *GetAudioFileResponseBodyData) SetAudioFileName(v string) *GetAudioFileResponseBodyData {
	s.AudioFileName = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetAudioResourceId(v string) *GetAudioFileResponseBodyData {
	s.AudioResourceId = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetCreatedTime(v string) *GetAudioFileResponseBodyData {
	s.CreatedTime = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetInstanceId(v string) *GetAudioFileResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetName(v string) *GetAudioFileResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetOssFileKey(v string) *GetAudioFileResponseBodyData {
	s.OssFileKey = &v
	return s
}

func (s *GetAudioFileResponseBodyData) SetUpdatedTime(v string) *GetAudioFileResponseBodyData {
	s.UpdatedTime = &v
	return s
}

func (s *GetAudioFileResponseBodyData) Validate() error {
	return dara.Validate(s)
}
