// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoInfosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVideoIds(v string) *GetVideoInfosRequest
	GetVideoIds() *string
}

type GetVideoInfosRequest struct {
	// The list of video IDs. Separate multiple IDs with commas (,). A maximum of 20 IDs can be specified.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7753d144efd8e649c6c45fe0579****,7753d144efd74d6c45fe0570****
	VideoIds *string `json:"VideoIds,omitempty" xml:"VideoIds,omitempty"`
}

func (s GetVideoInfosRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoInfosRequest) GoString() string {
	return s.String()
}

func (s *GetVideoInfosRequest) GetVideoIds() *string {
	return s.VideoIds
}

func (s *GetVideoInfosRequest) SetVideoIds(v string) *GetVideoInfosRequest {
	s.VideoIds = &v
	return s
}

func (s *GetVideoInfosRequest) Validate() error {
	return dara.Validate(s)
}
