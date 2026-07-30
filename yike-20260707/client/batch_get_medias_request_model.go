// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchGetMediasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *BatchGetMediasRequest
	GetAuthTimeout() *int64
	SetMediaIds(v string) *BatchGetMediasRequest
	GetMediaIds() *string
}

type BatchGetMediasRequest struct {
	// The validity period of the signed file access URL. Unit: seconds.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The IDs of the media assets to query, separated by commas.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
}

func (s BatchGetMediasRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchGetMediasRequest) GoString() string {
	return s.String()
}

func (s *BatchGetMediasRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *BatchGetMediasRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *BatchGetMediasRequest) SetAuthTimeout(v int64) *BatchGetMediasRequest {
	s.AuthTimeout = &v
	return s
}

func (s *BatchGetMediasRequest) SetMediaIds(v string) *BatchGetMediasRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetMediasRequest) Validate() error {
	return dara.Validate(s)
}
