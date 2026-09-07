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
	SetBizConfig(v string) *BatchGetMediasRequest
	GetBizConfig() *string
	SetMediaIds(v string) *BatchGetMediasRequest
	GetMediaIds() *string
	SetReturnDynamicMeta(v bool) *BatchGetMediasRequest
	GetReturnDynamicMeta() *bool
}

type BatchGetMediasRequest struct {
	// The validity period of the signed file access URL. Unit: seconds.
	//
	// example:
	//
	// 3600
	AuthTimeout *int64  `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	BizConfig   *string `json:"BizConfig,omitempty" xml:"BizConfig,omitempty"`
	// The IDs of the media assets to query, separated by commas.
	//
	// example:
	//
	// ******b48fb04483915d4f2cd8******,******c48fb37407365d4f2cd8******
	MediaIds          *string `json:"MediaIds,omitempty" xml:"MediaIds,omitempty"`
	ReturnDynamicMeta *bool   `json:"ReturnDynamicMeta,omitempty" xml:"ReturnDynamicMeta,omitempty"`
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

func (s *BatchGetMediasRequest) GetBizConfig() *string {
	return s.BizConfig
}

func (s *BatchGetMediasRequest) GetMediaIds() *string {
	return s.MediaIds
}

func (s *BatchGetMediasRequest) GetReturnDynamicMeta() *bool {
	return s.ReturnDynamicMeta
}

func (s *BatchGetMediasRequest) SetAuthTimeout(v int64) *BatchGetMediasRequest {
	s.AuthTimeout = &v
	return s
}

func (s *BatchGetMediasRequest) SetBizConfig(v string) *BatchGetMediasRequest {
	s.BizConfig = &v
	return s
}

func (s *BatchGetMediasRequest) SetMediaIds(v string) *BatchGetMediasRequest {
	s.MediaIds = &v
	return s
}

func (s *BatchGetMediasRequest) SetReturnDynamicMeta(v bool) *BatchGetMediasRequest {
	s.ReturnDynamicMeta = &v
	return s
}

func (s *BatchGetMediasRequest) Validate() error {
	return dara.Validate(s)
}
