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
	// The types of additional media asset information to return. If this parameter is not specified, only basic information is returned. Valid values are:
	//
	// -FileInfo
	//
	// -DynamicMetaData
	//
	// example:
	//
	// FileInfo,DynamicMetaData
	AdditionType *string `json:"AdditionType,omitempty" xml:"AdditionType,omitempty"`
	// The authentication timeout, in seconds.
	//
	// - Minimum value: **1**.
	//
	// - Maximum value: 86400.
	//
	// - Default value: 3600.
	//
	// example:
	//
	// 30
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// A comma-separated list of media asset IDs to query.
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
