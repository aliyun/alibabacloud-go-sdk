// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaIds(v string) *BatchGetMediaInfosRequest
	GetMediaIds() *string
	SetReferenceIds(v string) *BatchGetMediaInfosRequest
	GetReferenceIds() *string
}

type BatchGetMediaInfosRequest struct {
	// The media asset IDs, namely the audio or video IDs (VideoId). Separate multiple IDs with commas (,). A maximum of 20 IDs are supported. You can obtain the IDs by using one of the following methods:
	//
	// - For audio or video files uploaded through the console, log on to the ApsaraVideo VOD console and choose Media Library > Audio/Video to view the audio or video IDs.
	//
	// - When you call the operation to obtain the upload URL and credential for an audio or video file, the ID is the value of the VideoId response parameter.
	//
	// - After an audio or video file is uploaded, you can call the operation to search for media asset information. The ID is the value of the VideoId response parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****,7d2fbc380b0e08e55f****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The custom ID list. It consists of one or more custom IDs. Separate multiple IDs with commas (,). A maximum of 20 IDs are supported.
	//
	// example:
	//
	// 123-123,1234-1234
	ReferenceIds *string `json:"ReferenceIds,omitempty" xml:"ReferenceIds,omitempty"`
}

func (s BatchGetMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *BatchGetMediaInfosRequest) GetReferenceIds() *string {
	return s.ReferenceIds
}

func (s *BatchGetMediaInfosRequest) SetMediaIds(v string) *BatchGetMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetMediaInfosRequest) SetReferenceIds(v string) *BatchGetMediaInfosRequest {
	s.ReferenceIds = &v
	return s
}

func (s *BatchGetMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
