// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCastCrewListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v []*RecognizeVideoCastCrewListRequestParams) *RecognizeVideoCastCrewListRequest
	GetParams() []*RecognizeVideoCastCrewListRequestParams
	SetVideoUrl(v string) *RecognizeVideoCastCrewListRequest
	GetVideoUrl() *string
}

type RecognizeVideoCastCrewListRequest struct {
	Params []*RecognizeVideoCastCrewListRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequest) GetParams() []*RecognizeVideoCastCrewListRequestParams {
	return s.Params
}

func (s *RecognizeVideoCastCrewListRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *RecognizeVideoCastCrewListRequest) SetParams(v []*RecognizeVideoCastCrewListRequestParams) *RecognizeVideoCastCrewListRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) SetVideoUrl(v string) *RecognizeVideoCastCrewListRequest {
	s.VideoUrl = &v
	return s
}

func (s *RecognizeVideoCastCrewListRequest) Validate() error {
	if s.Params != nil {
		for _, item := range s.Params {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizeVideoCastCrewListRequestParams struct {
	// example:
	//
	// cast
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListRequestParams) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListRequestParams) GetType() *string {
	return s.Type
}

func (s *RecognizeVideoCastCrewListRequestParams) SetType(v string) *RecognizeVideoCastCrewListRequestParams {
	s.Type = &v
	return s
}

func (s *RecognizeVideoCastCrewListRequestParams) Validate() error {
	return dara.Validate(s)
}
