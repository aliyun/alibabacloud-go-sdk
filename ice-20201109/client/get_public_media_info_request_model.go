// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPublicMediaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMediaId(v string) *GetPublicMediaInfoRequest
	GetMediaId() *string
}

type GetPublicMediaInfoRequest struct {
	// example:
	//
	// icepublic-****14e501538aeef0a3140176f6****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s GetPublicMediaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPublicMediaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetPublicMediaInfoRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetPublicMediaInfoRequest) SetMediaId(v string) *GetPublicMediaInfoRequest {
	s.MediaId = &v
	return s
}

func (s *GetPublicMediaInfoRequest) Validate() error {
	return dara.Validate(s)
}
