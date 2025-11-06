// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSipMuteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrTracks(v []*RtcSipMuteResponseBodyErrTracks) *RtcSipMuteResponseBody
	GetErrTracks() []*RtcSipMuteResponseBodyErrTracks
	SetRequestId(v string) *RtcSipMuteResponseBody
	GetRequestId() *string
}

type RtcSipMuteResponseBody struct {
	ErrTracks []*RtcSipMuteResponseBodyErrTracks `json:"ErrTracks,omitempty" xml:"ErrTracks,omitempty" type:"Repeated"`
	// example:
	//
	// E7997404-5858-5C4D-94E4-33677412ACDC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RtcSipMuteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteResponseBody) GoString() string {
	return s.String()
}

func (s *RtcSipMuteResponseBody) GetErrTracks() []*RtcSipMuteResponseBodyErrTracks {
	return s.ErrTracks
}

func (s *RtcSipMuteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RtcSipMuteResponseBody) SetErrTracks(v []*RtcSipMuteResponseBodyErrTracks) *RtcSipMuteResponseBody {
	s.ErrTracks = v
	return s
}

func (s *RtcSipMuteResponseBody) SetRequestId(v string) *RtcSipMuteResponseBody {
	s.RequestId = &v
	return s
}

func (s *RtcSipMuteResponseBody) Validate() error {
	if s.ErrTracks != nil {
		for _, item := range s.ErrTracks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RtcSipMuteResponseBodyErrTracks struct {
	// ErrMsg。
	//
	// example:
	//
	// participant does not existed.
	ErrMsg *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	// example:
	//
	// 12
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// OperationId。
	//
	// example:
	//
	// 12122121
	OperationId *string `json:"OperationId,omitempty" xml:"OperationId,omitempty"`
}

func (s RtcSipMuteResponseBodyErrTracks) String() string {
	return dara.Prettify(s)
}

func (s RtcSipMuteResponseBodyErrTracks) GoString() string {
	return s.String()
}

func (s *RtcSipMuteResponseBodyErrTracks) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *RtcSipMuteResponseBodyErrTracks) GetId() *string {
	return s.Id
}

func (s *RtcSipMuteResponseBodyErrTracks) GetOperationId() *string {
	return s.OperationId
}

func (s *RtcSipMuteResponseBodyErrTracks) SetErrMsg(v string) *RtcSipMuteResponseBodyErrTracks {
	s.ErrMsg = &v
	return s
}

func (s *RtcSipMuteResponseBodyErrTracks) SetId(v string) *RtcSipMuteResponseBodyErrTracks {
	s.Id = &v
	return s
}

func (s *RtcSipMuteResponseBodyErrTracks) SetOperationId(v string) *RtcSipMuteResponseBodyErrTracks {
	s.OperationId = &v
	return s
}

func (s *RtcSipMuteResponseBodyErrTracks) Validate() error {
	return dara.Validate(s)
}
