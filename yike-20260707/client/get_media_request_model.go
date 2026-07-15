// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetMediaRequest
	GetAuthTimeout() *int64
	SetInputURL(v string) *GetMediaRequest
	GetInputURL() *string
	SetMediaId(v string) *GetMediaRequest
	GetMediaId() *string
}

type GetMediaRequest struct {
	// example:
	//
	// 3600
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// example:
	//
	// https://example-bucket.oss-cn-shanghai.aliyuncs.com/example.mp4
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetMediaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaRequest) GoString() string {
	return s.String()
}

func (s *GetMediaRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetMediaRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *GetMediaRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetMediaRequest) SetAuthTimeout(v int64) *GetMediaRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetMediaRequest) SetInputURL(v string) *GetMediaRequest {
	s.InputURL = &v
	return s
}

func (s *GetMediaRequest) SetMediaId(v string) *GetMediaRequest {
	s.MediaId = &v
	return s
}

func (s *GetMediaRequest) Validate() error {
	return dara.Validate(s)
}
