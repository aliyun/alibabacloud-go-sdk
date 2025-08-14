// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediaInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionType(v string) *BatchGetMediaInfosRequest
	GetAdditionType() *string
	SetAuthTimeout(v int64) *BatchGetMediaInfosRequest
	GetAuthTimeout() *int64
	SetMediaIds(v string) *BatchGetMediaInfosRequest
	GetMediaIds() *string
}

type BatchGetMediaInfosRequest struct {
	// The additional information that you want to query about the media assets. By default, only BasicInfo is returned. The following additional information can be queried:
	//
	// \\- FileInfo
	//
	// \\- DynamicMetaData
	//
	// example:
	//
	// FileInfo,DynamicMetaData
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	AuthTimeout  *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The IDs of the media assets that you want to query. Separate the IDs with commas (,).
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s BatchGetMediaInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediaInfosRequest) GoString() string {
	return s.String()
}

func (s *BatchGetMediaInfosRequest) GetAdditionType() *string {
	return s.AdditionType
}

func (s *BatchGetMediaInfosRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *BatchGetMediaInfosRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *BatchGetMediaInfosRequest) SetAdditionType(v string) *BatchGetMediaInfosRequest {
	s.AdditionType = &v
	return s
}

func (s *BatchGetMediaInfosRequest) SetAuthTimeout(v int64) *BatchGetMediaInfosRequest {
	s.AuthTimeout = &v
	return s
}

func (s *BatchGetMediaInfosRequest) SetMediaIds(v string) *BatchGetMediaInfosRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetMediaInfosRequest) Validate() error {
	return dara.Validate(s)
}
