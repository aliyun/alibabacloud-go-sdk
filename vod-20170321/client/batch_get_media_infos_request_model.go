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
	// The media asset IDs, which are audio/video IDs (VideoId). Separate multiple IDs with commas (,). You can specify up to 20 IDs. You can obtain the IDs by using the following methods:
	//
	// - For audio/video files uploaded through the console, log on to the ApsaraVideo VOD console and choose Media Files > Audio/Video to view the audio/video IDs.
	//
	// - When you call the operation to obtain the upload URL and credential for audio/video files, the VideoId value is returned as a response parameter.
	//
	// - After an audio/video file is uploaded, you can call the SearchMedia operation to query the VideoId value in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 61ccbdb06fa83012be4d8083f6****,7d2fbc380b0e08e55f****
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	// The list of custom IDs. Separate multiple custom IDs with commas (,). You can specify up to 20 IDs.
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
