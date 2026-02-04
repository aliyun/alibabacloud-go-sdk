// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastAudiosByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAudioIds(v []*string) *ListBroadcastAudiosByIdRequest
	GetAudioIds() []*string
}

type ListBroadcastAudiosByIdRequest struct {
	AudioIds []*string `json:"audioIds,omitempty" xml:"audioIds,omitempty" type:"Repeated"`
}

func (s ListBroadcastAudiosByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastAudiosByIdRequest) GoString() string {
	return s.String()
}

func (s *ListBroadcastAudiosByIdRequest) GetAudioIds() []*string {
	return s.AudioIds
}

func (s *ListBroadcastAudiosByIdRequest) SetAudioIds(v []*string) *ListBroadcastAudiosByIdRequest {
	s.AudioIds = v
	return s
}

func (s *ListBroadcastAudiosByIdRequest) Validate() error {
	return dara.Validate(s)
}
