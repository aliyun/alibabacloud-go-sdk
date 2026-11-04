// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePictureSearchJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceId(v string) *CreatePictureSearchJobRequest
	GetAppInstanceId() *string
	SetBodyThreshold(v float32) *CreatePictureSearchJobRequest
	GetBodyThreshold() *float32
	SetEndTime(v int64) *CreatePictureSearchJobRequest
	GetEndTime() *int64
	SetPictureSearchType(v int32) *CreatePictureSearchJobRequest
	GetPictureSearchType() *int32
	SetSearchPicUrl(v string) *CreatePictureSearchJobRequest
	GetSearchPicUrl() *string
	SetStartTime(v int64) *CreatePictureSearchJobRequest
	GetStartTime() *int64
	SetThreshold(v float32) *CreatePictureSearchJobRequest
	GetThreshold() *float32
}

type CreatePictureSearchJobRequest struct {
	// Instance ID of the image search application instance. You can call the [QueryPictureSearchApps](https://help.aliyun.com/document_detail/206608.html) operation to query instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e429560124b44b685095df81a27****
	AppInstanceId *string `json:"AppInstanceId,omitempty" xml:"AppInstanceId,omitempty"`
	// The threshold for human body search. Valid values: 0 to 1.
	//
	// To ensure the accuracy of search results, set this parameter to a value greater than or equal to 0.8. If this parameter is not specified, the value of the Threshold parameter is used.
	//
	// example:
	//
	// 0.9
	BodyThreshold *float32 `json:"BodyThreshold,omitempty" xml:"BodyThreshold,omitempty"`
	// The end time of the search range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614327915000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The search type. Valid values:
	//
	// - 1: face
	//
	// - 0: human body
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PictureSearchType *int32 `json:"PictureSearchType,omitempty" xml:"PictureSearchType,omitempty"`
	// The URL of the image to be searched.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://xxx.xxx/1.jpg
	SearchPicUrl *string `json:"SearchPicUrl,omitempty" xml:"SearchPicUrl,omitempty"`
	// The start time of the search range. The time is a 13-digit timestamp in milliseconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1614227915000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The threshold for face search. Valid values: 0 to 1.
	//
	// To ensure the accuracy of search results, set this parameter to a value between 0.3 and 0.6.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.5
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
}

func (s CreatePictureSearchJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePictureSearchJobRequest) GoString() string {
	return s.String()
}

func (s *CreatePictureSearchJobRequest) GetAppInstanceId() *string {
	return s.AppInstanceId
}

func (s *CreatePictureSearchJobRequest) GetBodyThreshold() *float32 {
	return s.BodyThreshold
}

func (s *CreatePictureSearchJobRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *CreatePictureSearchJobRequest) GetPictureSearchType() *int32 {
	return s.PictureSearchType
}

func (s *CreatePictureSearchJobRequest) GetSearchPicUrl() *string {
	return s.SearchPicUrl
}

func (s *CreatePictureSearchJobRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *CreatePictureSearchJobRequest) GetThreshold() *float32 {
	return s.Threshold
}

func (s *CreatePictureSearchJobRequest) SetAppInstanceId(v string) *CreatePictureSearchJobRequest {
	s.AppInstanceId = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetBodyThreshold(v float32) *CreatePictureSearchJobRequest {
	s.BodyThreshold = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetEndTime(v int64) *CreatePictureSearchJobRequest {
	s.EndTime = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetPictureSearchType(v int32) *CreatePictureSearchJobRequest {
	s.PictureSearchType = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetSearchPicUrl(v string) *CreatePictureSearchJobRequest {
	s.SearchPicUrl = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetStartTime(v int64) *CreatePictureSearchJobRequest {
	s.StartTime = &v
	return s
}

func (s *CreatePictureSearchJobRequest) SetThreshold(v float32) *CreatePictureSearchJobRequest {
	s.Threshold = &v
	return s
}

func (s *CreatePictureSearchJobRequest) Validate() error {
	return dara.Validate(s)
}
