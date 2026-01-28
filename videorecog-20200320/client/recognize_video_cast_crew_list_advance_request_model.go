// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRecognizeVideoCastCrewListAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParams(v []*RecognizeVideoCastCrewListAdvanceRequestParams) *RecognizeVideoCastCrewListAdvanceRequest
	GetParams() []*RecognizeVideoCastCrewListAdvanceRequestParams
	SetVideoUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type RecognizeVideoCastCrewListAdvanceRequest struct {
	Params []*RecognizeVideoCastCrewListAdvanceRequestParams `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// https://shanghai.oss-cn-shanghai.aliyuncs.com/download/xxxx.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) GetParams() []*RecognizeVideoCastCrewListAdvanceRequestParams {
	return s.Params
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetParams(v []*RecognizeVideoCastCrewListAdvanceRequestParams) *RecognizeVideoCastCrewListAdvanceRequest {
	s.Params = v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) SetVideoUrlObject(v io.Reader) *RecognizeVideoCastCrewListAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequest) Validate() error {
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

type RecognizeVideoCastCrewListAdvanceRequestParams struct {
	// example:
	//
	// cast
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCastCrewListAdvanceRequestParams) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCastCrewListAdvanceRequestParams) GetType() *string {
	return s.Type
}

func (s *RecognizeVideoCastCrewListAdvanceRequestParams) SetType(v string) *RecognizeVideoCastCrewListAdvanceRequestParams {
	s.Type = &v
	return s
}

func (s *RecognizeVideoCastCrewListAdvanceRequestParams) Validate() error {
	return dara.Validate(s)
}
