// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPlayInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthTimeout(v int64) *GetPlayInfoRequest
	GetAuthTimeout() *int64
	SetInputURL(v string) *GetPlayInfoRequest
	GetInputURL() *string
	SetMediaId(v string) *GetPlayInfoRequest
	GetMediaId() *string
}

type GetPlayInfoRequest struct {
	AuthTimeout *int64 `json:"AuthTimeout,omitempty" xml:"AuthTimeout,omitempty"`
	// The input URL that you specified for the media asset when you registered the media asset. For more information, see [RegisterMediaInfo](https://help.aliyun.com/document_detail/441152.html).
	//
	// >  You must specify at least one of the MediaId and InputURL parameters.
	InputURL *string `json:"InputURL,omitempty" xml:"InputURL,omitempty"`
	// The ID of the media asset.
	//
	// >  You must specify at least one of the MediaId and InputURL parameters.
	//
	// example:
	//
	// 86434e152b7d4f20be480574439fe***
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetPlayInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPlayInfoRequest) GetAuthTimeout() *int64 {
	return s.AuthTimeout
}

func (s *GetPlayInfoRequest) GetInputURL() *string {
	return s.InputURL
}

func (s *GetPlayInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetPlayInfoRequest) SetAuthTimeout(v int64) *GetPlayInfoRequest {
	s.AuthTimeout = &v
	return s
}

func (s *GetPlayInfoRequest) SetInputURL(v string) *GetPlayInfoRequest {
	s.InputURL = &v
	return s
}

func (s *GetPlayInfoRequest) SetMediaId(v string) *GetPlayInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetPlayInfoRequest) Validate() error {
	return dara.Validate(s)
}
