// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastVideosByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoIds(v []*string) *ListBroadcastVideosByIdRequest
	GetVideoIds() []*string
}

type ListBroadcastVideosByIdRequest struct {
	VideoIds []*string `json:"videoIds,omitempty" xml:"videoIds,omitempty" type:"Repeated"`
}

func (s ListBroadcastVideosByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastVideosByIdRequest) GoString() string {
	return s.String()
}

func (s *ListBroadcastVideosByIdRequest) GetVideoIds() []*string {
	return s.VideoIds
}

func (s *ListBroadcastVideosByIdRequest) SetVideoIds(v []*string) *ListBroadcastVideosByIdRequest {
	s.VideoIds = v
	return s
}

func (s *ListBroadcastVideosByIdRequest) Validate() error {
	return dara.Validate(s)
}
